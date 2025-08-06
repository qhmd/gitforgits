package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/qhmd/gitforgits/cart-service/handler"
	"github.com/qhmd/gitforgits/cart-service/pkg"
	"github.com/qhmd/gitforgits/cart-service/repo"
	"github.com/qhmd/gitforgits/cart-service/usecase"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	db := pkg.InitMySQL()
	pkg.RunMigration(db)

	repoBook := repo.NewCartRepository(db)

	ucBook := usecase.NewCartUsecase(repoBook)

	handler.NewCartHandler(app, ucBook)

	app.Listen(":8080") // Jangan lupa tambahkan ini agar servernya berjalan
}
