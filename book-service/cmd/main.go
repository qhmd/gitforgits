package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/qhmd/gitforgits/book-service/handler"
	"github.com/qhmd/gitforgits/book-service/repo"
	"github.com/qhmd/gitforgits/book-service/usecase"
	"github.com/qhmd/gitforgits/pkg/database"
	"github.com/qhmd/gitforgits/utils"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @title           GitForGits API
// @version         1.0
// @description     API documentation for project GitForGits
// @description     Login as admin:
// @description     email: admingitforgits12@gmail.com
// @description     password: @GitForGitsAdmin21
// @termsOfService  http://swagger.io/terms/

// @securityDefinitions.apikey  BearerAuth
// @in                         header
// @name                       Authorization
// @description                Type "Bearer" followed by a space and JWT token.

func main() {
	app := fiber.New()
	app.Use(cors.New())

	utils.InitValidator()
	fmt.Print("listen in port 8080...")
	db := database.InitMySQL()
	database.RunMigration(db)
	repoBook := repo.NewMySQLBookRepository(db)

	ucBook := usecase.NewBookUsecase(repoBook)

	handler.NewBookHandler(app, ucBook)
	app.Listen(":8080")
}
