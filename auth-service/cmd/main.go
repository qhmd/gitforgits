package main

import (
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	usersproto "github.com/qhmd/gitforgits/shared/proto/users-proto"

	"github.com/qhmd/gitforgits/auth-service/client"
	"github.com/qhmd/gitforgits/auth-service/handler"
	"github.com/qhmd/gitforgits/auth-service/pkg/database"
	"github.com/qhmd/gitforgits/auth-service/repo"
	"github.com/qhmd/gitforgits/auth-service/usecase"
	_ "github.com/qhmd/gitforgits/cmd/server/docs"
	"github.com/qhmd/gitforgits/shared/utils"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// setup ...
	authClient := client.NewAuthServiceClient("userservice:50051")
	app := fiber.New()
	app.Use(cors.New())
	utils.InitValidator()

	db := database.InitMySQL()
	database.RunMigration(db)

	repoAuth := repo.NewMySQLAuthRepository(db)

	ucAuth := usecase.NewAuthUsecase(repoAuth, authClient)

	// handler
	handler.NewAuthHandler(app, ucAuth)
	GrcpHandler := handler.NewUsersGrcpHandler(ucAuth)

	go func() {
		listener, err := net.Listen("tcp", ":50052")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		usersproto.RegisterUsersServiceServer(grpcServer, GrcpHandler)
		log.Println("gRPC server running on port 50052")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Println("HTTP server running on port 8080")
	app.Listen(":8080")
}
