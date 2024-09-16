package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"demo/db"
)

// Constant variables used across all modules
const (
	// Database file name
	dbFileName = "yolo.db"

	// Server port
	serverPort = ":8080"
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

// A handler for adding a new user
func addUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("name")
	email := r.FormValue("email")

	database := db.GetDB(dbFileName)
	// Add user to the database
	db.AddUser(database, username, email)

	// Return success message
	getUsers(w, r)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	database := db.GetDB(dbFileName)
	users, err := db.GetAllUsers(database)

	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	// Return the users in HTML format
	for _, user := range users {
		fmt.Fprintf(w, `<li>%s</li>`, user.Name+" "+user.Email)
	}
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	productName := r.FormValue("name")
	price := r.FormValue("price")

	// Convert price to float64
	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		http.Error(w, "Invalid price value", http.StatusBadRequest)
		return
	}

	database := db.GetDB(dbFileName)
	// Add product to the database
	db.AddProduct(database, productName, priceFloat)

	// Return success message
	getProducts(w, r)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	database := db.GetDB(dbFileName)
	products, err := db.GetAllProducts(database)

	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	// Return the products in HTML format
	for _, product := range products {
		fmt.Fprintf(w, `<li>%s</li>`, product.Name+" "+fmt.Sprintf("%f", product.Price))
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	// Serve static files like index.html
	fs := http.FileServer(http.Dir("./web"))

	db.CreateProductTable(db.GetDB(dbFileName))

	// Define routes
	http.Handle("/", fs)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/load-more-posts", loadMorePosts)
	http.HandleFunc("/submit-comment", submitComment)
	http.HandleFunc("/add-user", addUserHandler)
	http.HandleFunc("/get-users", getUsers)
	http.HandleFunc("/add-product", addProduct)
	http.HandleFunc("/get-products", getProducts)

	// Start the server
	fmt.Println("Starting server at port " + serverPort + "...")
	if err := http.ListenAndServe(serverPort, nil); err != nil {
		log.Fatal(err)
	}
}
