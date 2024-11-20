package main

import (
	"log"
	"os"
	"rest_api_muti/config"
	"rest_api_muti/controllers"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	// Memuat file .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found. Defaulting to environment variables.")
	}
}

func main() {
	// Koneksi ke database
	config.ConnectDatabase()

	// Inisialisasi Echo
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		e.Logger.Fatal("JWT_SECRET tidak ditemukan di environment variable")
	}

	// Rute untuk registrasi dan login tanpa autentikasi
	e.POST("/api/v1/register", controllers.RegisterHandler)
	e.POST("/api/v1/login", controllers.LoginHandler)

	// Kelompok rute yang memerlukan autentikasi
	eAuth := e.Group("")
	eAuth.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwtSecret),
	}))

	// Routes untuk Category
	eAuth.GET("/api/v1/categories", controllers.GetAllCategoriesHandler)
	eAuth.GET("/api/v1/categories/:id", controllers.GetCategoryHandler)
	eAuth.POST("/api/v1/categories", controllers.CreateCategoryHandler)
	eAuth.PUT("/api/v1/categories/:id", controllers.UpdateCategoryHandler)
	eAuth.DELETE("/api/v1/categories/:id", controllers.DeleteCategoryHandler)

	// Routes untuk Product
	eAuth.GET("/api/v1/products", controllers.GetAllProductsHandler)
	eAuth.GET("/api/v1/products/:id", controllers.GetProductHandler)
	eAuth.POST("/api/v1/products", controllers.CreateProductHandler)
	eAuth.PUT("/api/v1/products/:id", controllers.UpdateProductsHandler)
	eAuth.DELETE("/api/v1/products/:id", controllers.DeleteProductHandler)

	// Routes untuk Transaction
	eAuth.GET("/api/v1/transactions", controllers.GetAllTransactionHandler)
	eAuth.GET("/api/v1/transactions/:id", controllers.GetTransactionHandler)
	eAuth.POST("/api/v1/transactions", controllers.CreateTransactionHandler)
	eAuth.PUT("/api/v1/transactions/:id", controllers.UpdateTransactionHandler)
	eAuth.DELETE("/api/v1/transactions/:id", controllers.DeleteTransactionHandler)

	// Memulai server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Starting server on port %s...", port)
	if err := e.Start(":" + port); err != nil {
		e.Logger.Fatal("Failed to start server: ", err)
	}
}
