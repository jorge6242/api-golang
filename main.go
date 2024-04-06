package main

import (
	"api-golang/bootstrap"
	"api-golang/config"
	"api-golang/controller"
	"log"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.LoadConfig()

	container := bootstrap.BuildContainer(cfg)

	// Invoke the container to inject the controller and configure the routes
	err := container.Invoke(func(productController *controller.ProductController) {
		http.HandleFunc("/productos", productController.GetProducts)
		http.HandleFunc("/producto", productController.CreateProduct)
	})

	if err != nil {
		log.Fatalf("Failed to invoke container: %v", err)
	}

	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
