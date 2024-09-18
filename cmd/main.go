package main

import (
	driver "database/sql"
	"log"
	"sistemaLogistica/internal/adapter/handlers"
	"sistemaLogistica/internal/adapter/repository"
	"sistemaLogistica/internal/controllers"
	"sistemaLogistica/internal/usecase/product"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
)

func main() {

	product, err := injectDependency()
	if err != nil {
		log.Fatalf("Error when injecting dependencies: %v", err)
	}

	productHandler := handlers.NewProductHandler(product)

	app := fiber.New()
	app.Post("/", productHandler.HandlerPost)
	app.Get("/", productHandler.HandlerGet)
	app.Put("/", productHandler.HandlerPut)
	app.Delete("/", productHandler.HandlerDelete)

	app.Listen(":8080")

}

func injectDependency() (controllers.ProductUsecase, error) {
	db, err := driver.Open("postgres", "host=localhost port=5432 user= test password=test dbname=test sslmode=disable")
	if err != nil {
		return nil, err
	}

	repo := repository.NewRepoSql(db)
	productRepo := repository.NewProductRepo(repo)
	totalProductRepo := repository.NewTotalProductRepo(repo)
	return product.NewProductUseCase(productRepo, totalProductRepo), nil
}
