package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/qhmd/gitforgits/auth-service/handler"
	"github.com/qhmd/gitforgits/auth-service/pkg/database"
	"github.com/qhmd/gitforgits/auth-service/repo"
	"github.com/qhmd/gitforgits/auth-service/usecase"
	_ "github.com/qhmd/gitforgits/cmd/server/docs"
	"github.com/qhmd/gitforgits/shared/utils"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	utils.InitValidator()
	fmt.Print("listen in port 8000...")
	db := database.InitMySQL()
	database.RunMigration(db)
	repoAuth := repo.NewMySQLAuthRepository(db)

	ucAuth := usecase.NewAuthUsecase(repoAuth)

	handler.NewAuthHandler(app, ucAuth)
	app.Listen(":8080")
}
