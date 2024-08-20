package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"test-backend-pari/items/handlers"
	"test-backend-pari/items/repositories"
	"test-backend-pari/items/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	log.Println(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	itemRepository := repositories.NewItemRepository(db)
	itemUsecase := usecases.NewItemUsecase(itemRepository)
	itemHandler := handlers.NewItemHandler(itemUsecase)

	handlers.MapItem(app, itemHandler)

	app.Listen(":3000")
}
