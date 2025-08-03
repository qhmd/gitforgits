package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/qhmd/gitforgits/shared/utils"
	"github.com/qhmd/gitforgits/users-service/handler"
	"github.com/qhmd/gitforgits/users-service/pkg/database"
	"github.com/qhmd/gitforgits/users-service/repo"
	"github.com/qhmd/gitforgits/users-service/usecase"
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
	fmt.Print("listen in port 8080...")
	db := database.InitMySQL()
	database.RunMigration(db)
	repoUsers := repo.NewUserMySqlRepo(db)

	ucUsers := usecase.NewUsersUseCase(repoUsers)

	handler.NewHandlerUser(app, ucUsers)
	app.Listen(":8080")
}
