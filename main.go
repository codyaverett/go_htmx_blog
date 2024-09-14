package main

import (
	"fmt"
	"log"
	"net/http"
)

func loadMorePosts(w http.ResponseWriter, r *http.Request) {
	// Generate HTML for more blog posts
	fmt.Fprint(w, `
      <article>
          <h3>New Post Title</h3>
          <p>This is another dynamic blog post loaded with HTMX.</p>
      </article>
  `)
}

func submitComment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.FormValue("comment")

	// Return the comment in HTML format
	fmt.Fprintf(w, `
      <li>%s</li>
  `, comment)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func handleDefault(w http.ResponseWriter, r *http.Request) {
	// Serve static files like index.html
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
}

func main() {
	// Serve static files like index.html
	fs := http.FileServer(http.Dir("./web"))

	// Define routes
	http.Handle("/", fs)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/load-more-posts", loadMorePosts)
	http.HandleFunc("/submit-comment", submitComment)

	// Start the server
	fmt.Println("Starting server at port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
