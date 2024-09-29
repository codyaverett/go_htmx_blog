package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

// Server struct to hold uploaded images and a mutex for concurrency safety
type Server struct {
	images []string
	mu     sync.Mutex
}

// Handler for serving the index.html file
func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

// Handler for processing image uploads
func (s *Server) uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		http.Error(w, "Could not parse multipart form", http.StatusBadRequest)
		return
	}

	// Retrieve the file from form data
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Could not get uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Ensure the uploads directory exists
	os.MkdirAll("uploads", os.ModePerm)

	// Save the uploaded file to the server
	dst, err := os.Create(fmt.Sprintf("uploads/%s", handler.Filename))
	if err != nil {
		http.Error(w, "Could not save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Could not save file", http.StatusInternalServerError)
		return
	}

	// Add the filename to the images slice
	s.mu.Lock()
	s.images = append(s.images, handler.Filename)
	s.mu.Unlock()

	// Render the updated images table
	s.renderImages(w)
}

// Function to render the images table as HTML
func (s *Server) renderImages(w http.ResponseWriter) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tmpl := `
    <div id="images">
        <table border="1">
            <tr>
                <th>Image Name</th>
                <th>Image</th>
            </tr>
            {{range .}}
            <tr>
                <td>{{.}}</td>
                <td><img src="/uploads/{{.}}" width="100"></td>
            </tr>
            {{end}}
        </table>
    </div>
    `
	t := template.Must(template.New("images").Parse(tmpl))
	t.Execute(w, s.images)
}

// Main function to start the server and route handlers
func main() {
	srv := &Server{}
	http.HandleFunc("/", srv.indexHandler)
	http.HandleFunc("/upload", srv.uploadHandler)
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
