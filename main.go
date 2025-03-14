// filepath: /Users/joseflores/git/go/main.go
package main

import (
	"fmt"
	_ "my-go-service/docs" // Import the generated docs
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// helloHandler godoc
// @Summary Say hello to user
// @Description get string by name
// @Tags example
// @Accept  json
// @Produce  json
// @Param name query string true "Name"
// @Success 200 {string} string "Hello {name}"
// @Failure 400 {string} string "Name is required"
// @Router /hello [get]
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s", name)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
