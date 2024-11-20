package controllers

import (
	"fmt"
	"net/http"
	"os"
	"rest_api_muti/config"
	"rest_api_muti/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaims struct {
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
	jwt.RegisteredClaims
}

// HashPassword membuat hash dari password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash membandingkan hash password dengan input
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT membuat JWT untuk pengguna
func GenerateJWT(userID int, name string) (string, error) {
	claims := &jwtCustomClaims{
		Name:   name,
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), // Berlaku 72 jam
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET tidak ditemukan di environment")
	}

	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

// LoginHandler menangani proses login
func LoginHandler(c echo.Context) error {
	userInput := models.User{}
	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "Format data tidak valid", Data: nil})
	}

	// Cari pengguna berdasarkan email
	user := models.User{}
	result := config.DB.First(&user, "email = ?", userInput.Email)
	if result.Error != nil || user.ID == 0 {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{Status: false, Message: "Email tidak ditemukan", Data: nil})
	}

	// Verifikasi password
	match := CheckPasswordHash(userInput.Password, user.Password)
	if !match {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{Status: false, Message: "Password salah", Data: nil})
	}

	// Buat token JWT
	token, err := GenerateJWT(int(user.ID), user.Nama)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Gagal generate token", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Login berhasil", Data: map[string]string{"token": token}})
}

// RegisterHandler menangani proses registrasi
func RegisterHandler(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "Format data tidak valid", Data: nil})
	}

	// Hash password sebelum menyimpan ke database
	hash, err := HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Gagal mengolah password", Data: nil})
	}
	user.Password = hash

	// Simpan user ke database
	result := config.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Gagal register", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Registrasi berhasil", Data: user})
}
