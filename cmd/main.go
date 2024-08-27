package main

import (
	"loan-tracker/api/controllers"
	// "loan-tracker/api/middleware"
	"loan-tracker/api/route"
	"loan-tracker/config"
	"loan-tracker/internal"

	"loan-tracker/repositories"
	"loan-tracker/usecases"
	"log"


	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	client, err := config.InitMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// db := middleware.NewDatabaseConnection()
	userCollection := client.Database("Loan").Collection("Users")
	// tokenCollection := client.Database("Loan").Collection("Tokens")

	userMockCollection := repositories.NewMongoCollection(userCollection)
	// tokenMockCollection := repositories.NewMongoCollection(tokenCollection)

	userRepo := repositories.NewUserRepository(userMockCollection)
	// tokenRepo := repositories.NewTokenRepository(tokenMockCollection)

	// jwtService := middleware.NewJWTService(os.Getenv("JWT_SECRET"), "Kal", os.Getenv("JWT_REFRESH_SECRET"))

	// Usecase setup
	// userUsecase := usecases.NewUserUsecase(userRepo, tokenRepo, jwtService)
	registerUsecase := usecases.NewRegisterUsecase(userRepo)
	verifyEmailUsecase := usecases.NewVerifyEmailUsecase(userRepo)
	// resetPasswordUsecase := usecases.NewResetPasswordUsecase(userRepo)


	// Email setup
	mailConfig := internal.SMTPConfig{
		Host:     "smtp.gmail.com",
		Port:     "587",
		Username: "kalkidanamare11a@gmail.com",
		Password: "jcwf vfzi njtd rayo",
	}

	registerCtrl := controllers.NewRegisterController(registerUsecase, mailConfig)
	verifyEmailCtrl := controllers.NewVerifyEmailController(verifyEmailUsecase)
	// resetPasswordCtrl := controllers.NewResetPasswordController(resetPasswordUsecase, mailConfig)
	
	
	router := route.SetupRouter(registerCtrl, verifyEmailCtrl, nil)

	// Start server
	router.Run(":8080")
}
