package main

import (
	"log"
	"os"

	"loan-tracker/api/controllers"
	"loan-tracker/api/middleware"
	"loan-tracker/api/routers"
	"loan-tracker/config"
	"loan-tracker/internal"
	"loan-tracker/repositories"
	"loan-tracker/usecases"

	"github.com/gin-gonic/gin"
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

	// Initialize collections
	userCollection := client.Database("Loan").Collection("Users")
	tokenCollection := client.Database("Loan").Collection("Tokens")
	otpCollection := client.Database("Loan").Collection("OTPs")
	loanCollection := client.Database("Loan").Collection("Loans")

	userMockCollection := repositories.NewMongoCollection(userCollection)
	tokenMockCollection := repositories.NewMongoCollection(tokenCollection)
	otpMockCollection := repositories.NewMongoCollection(otpCollection)
	loanMockCollection := repositories.NewMongoCollection(loanCollection)

	// Initialize repositories
	userRepo := repositories.NewUserRepository(userMockCollection)
	tokenRepo := repositories.NewTokenRepository(tokenMockCollection)
	otpRepo := repositories.NewOtpRepository(otpMockCollection)
	loanRepo := repositories.NewLoanRepository(loanMockCollection)

	// Initialize services
	jwtService := middleware.NewJWTService(os.Getenv("JWT_SECRET"), "Kal", os.Getenv("JWT_REFRESH_SECRET"))

	// Initialize usecases
	userUsecase := usecases.NewUserUsecase(userRepo, tokenRepo, jwtService)
	registerUsecase := usecases.NewRegisterUsecase(userRepo)
	verifyEmailUsecase := usecases.NewVerifyEmailUsecase(userRepo)
	otpUsecase := usecases.NewOTPUsecase(otpRepo, userRepo)
	tokenUsecase := usecases.NewTokenUsecase(tokenRepo, jwtService)
	loanUsecase := usecases.NewLoanUsecase(loanRepo)

	// Initialize email configuration
	mailConfig := internal.SMTPConfig{
		Host:     "smtp.gmail.com",
		Port:     "587",
		Username: "kalkidanamare11a@gmail.com",
		Password: "jcwf vfzi njtd rayo",
	}

	// Initialize controllers
	userController := controllers.NewUserController(userUsecase)
	registerCtrl := controllers.NewRegisterController(registerUsecase, mailConfig)
	verifyEmailCtrl := controllers.NewVerifyEmailController(verifyEmailUsecase)
	tokenController := controllers.NewRefreshTokenController(userUsecase, tokenUsecase, jwtService)
	forgotPController := controllers.NewForgotPasswordController(userUsecase, otpUsecase)
	logoutController := controllers.NewLogoutController(tokenUsecase)
	loanController := controllers.NewLoanController(loanUsecase)
	// resetPasswordCtrl := controllers.NewResetPasswordController(resetPasswordUsecase, mailConfig)

	// Set up the Gin router
	router := gin.Default()

	// Set up routers
	routers.UserRouter(router, userController, jwtService)
	routers.AuthRouter(router, registerCtrl, verifyEmailCtrl)
	routers.PasswordRouter(router, forgotPController /*, resetPasswordCtrl*/)
	routers.LogoutRouter(router, logoutController, jwtService)
	routers.AdminRouter(router, loanController,userController, jwtService)
	routers.LoanRouter(router, loanController, jwtService)
	routers.RefreshTokenRouter(router, tokenController)

	router.Run(":8080")
}
