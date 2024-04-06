package bootstrap

import (
	"api-golang/config"
	"api-golang/controller"
	"api-golang/database"
	"api-golang/repositories"
	"api-golang/services"
	"log"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

func BuildContainer(cfg config.Config) *dig.Container {
	container := dig.New()

	// Database Connection Registration
	container.Provide(func() (*gorm.DB, error) {
		return database.ConnectDatabase(cfg)
	})

	// Repositories
	if err := container.Provide(repositories.NewProductoRepository); err != nil {
		log.Fatalf("Failed to provide product repository: %v", err)
	}

	// Services
	if err := container.Provide(services.NewProductoService); err != nil {
		log.Fatalf("Failed to provide product service: %v", err)
	}

	// Controller
	if err := container.Provide(controller.NewProductController); err != nil {
		log.Fatalf("Failed to provide product controller: %v", err)
	}

	return container
}
