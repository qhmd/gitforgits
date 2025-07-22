package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	handler "github.com/qhmd/gitforgits/internal/delivery/http"
	repo "github.com/qhmd/gitforgits/internal/repository"
	useCase "github.com/qhmd/gitforgits/internal/usecase"
	"github.com/qhmd/gitforgits/pkg/database"
	"github.com/qhmd/gitforgits/utils"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app := fiber.New()
	utils.InitValidator()
	fmt.Print("listen in port 8080...")
	db := database.InitMySQL()
	database.RunMigration(db)

	repoBook := repo.NewMySQLBookRepository(db)
	repoAuth := repo.NewMySQLAuthRepository(db)

	ucBook := useCase.NewBookUsecase(repoBook)
	ucAuth := useCase.NewAuthUsecase(repoAuth)

	handler.NewBookHandler(app, ucBook)
	handler.NewAuthHandler(app, ucAuth)
	app.Get("/", Welcome)

	app.Listen(":8080")
}

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcodd Bro")
}
