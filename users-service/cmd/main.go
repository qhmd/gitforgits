package main

import (
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	usersproto "github.com/qhmd/gitforgits/shared/proto/auth-proto"
	"github.com/qhmd/gitforgits/shared/utils"
	"github.com/qhmd/gitforgits/users-service/client"
	"github.com/qhmd/gitforgits/users-service/handler"
	"github.com/qhmd/gitforgits/users-service/pkg/database"
	"github.com/qhmd/gitforgits/users-service/repo"
	"github.com/qhmd/gitforgits/users-service/usecase"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// init and db
	utils.InitValidator()
	db := database.InitMySQL()
	database.RunMigration(db)

	// for uc
	repoUsers := repo.NewUserMySqlRepo(db)
	clientAuth := client.NewAuthServiceClient("authservice:50052")

	// app from fiber
	app := fiber.New()
	app.Use(cors.New())

	ucUsers := usecase.NewUsersUseCase(repoUsers, clientAuth)

	// handler
	handler.NewHandlerUser(app, ucUsers)
	userGrcpHandler := handler.NewAuthGrcpHandler(*ucUsers)

	go func() {
		listener, err := net.Listen("tcp", "0.0.0.0:50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Println("gRPC server running on port 9000")

		grpcServer := grpc.NewServer()
		usersproto.RegisterAuthServiceServer(grpcServer, userGrcpHandler)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		log.Println("gRPC serve finished")
	}()
	log.Println("HTTP server running on port 8080")
	app.Listen(":8080")

}
