package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/qhmd/gitforgits/internal/delivery/http/book"
	bookRepo "github.com/qhmd/gitforgits/internal/repository/book"
	bookUseCase "github.com/qhmd/gitforgits/internal/usecase/book"
	"github.com/qhmd/gitforgits/pkg/database"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app := fiber.New()
	fmt.Print("listen in port 8080...")
	db := database.InitMySQL()
	database.RunMigaration(db)

	repo := bookRepo.NewMySQLBookRepository(db)
	uc := bookUseCase.NewBookUsecase(repo)
	book.NewBookHandler(app, uc)
	app.Get("/", Welcome)

	app.Listen(":8080")
}

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcodd Bro")
}
