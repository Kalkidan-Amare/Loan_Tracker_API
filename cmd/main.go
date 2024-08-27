package main

import (
	"loan-tracker/api/controllers"
	"loan-tracker/api/middleware"
	"loan-tracker/api/route"
	"loan-tracker/config"
	"loan-tracker/internal"

	"loan-tracker/repositories"
	"loan-tracker/usecases"
	"log"
	"os"


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
	tokenCollection := client.Database("Loan").Collection("Tokens")
	otpCollection := client.Database("Loan").Collection("OTPs")

	userMockCollection := repositories.NewMongoCollection(userCollection)
	tokenMockCollection := repositories.NewMongoCollection(tokenCollection)
	otpMockCollection := repositories.NewMongoCollection(otpCollection)

	userRepo := repositories.NewUserRepository(userMockCollection)
	tokenRepo := repositories.NewTokenRepository(tokenMockCollection)
	otpRepo := repositories.NewOtpRepository(otpMockCollection)

	jwtService := middleware.NewJWTService(os.Getenv("JWT_SECRET"), "Kal", os.Getenv("JWT_REFRESH_SECRET"))

	// Usecase setup
	userUsecase := usecases.NewUserUsecase(userRepo, tokenRepo, jwtService)
	registerUsecase := usecases.NewRegisterUsecase(userRepo)
	verifyEmailUsecase := usecases.NewVerifyEmailUsecase(userRepo)
	// resetPasswordUsecase := usecases.NewResetPasswordUsecase(userRepo)
	otpUsecase := usecases.NewOTPUsecase(otpRepo,userRepo)
	tokenUsecase := usecases.NewTokenUsecase(tokenRepo, jwtService)


	// Email setup
	mailConfig := internal.SMTPConfig{
		Host:     "smtp.gmail.com",
		Port:     "587",
		Username: "kalkidanamare11a@gmail.com",
		Password: "jcwf vfzi njtd rayo",
	}

	userController := controllers.NewUserController(userUsecase)
	registerCtrl := controllers.NewRegisterController(registerUsecase, mailConfig)
	verifyEmailCtrl := controllers.NewVerifyEmailController(verifyEmailUsecase)
	// resetPasswordCtrl := controllers.NewResetPasswordController(resetPasswordUsecase, mailConfig)
	// otpController := controllers.NewOTPController(otpUsecase)
	tokenController := controllers.NewRefreshTokenController(userUsecase,tokenUsecase,jwtService)
	forgotPController := controllers.NewForgotPasswordController(userUsecase, otpUsecase)
	logoutController := controllers.NewLogoutController(tokenUsecase)
	
	
	router := route.SetupRouter(userController,registerCtrl, verifyEmailCtrl,tokenController,forgotPController,logoutController, nil,jwtService)

	// Start server
	router.Run(":8080")
}
