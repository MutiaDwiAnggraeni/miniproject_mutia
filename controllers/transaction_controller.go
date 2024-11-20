package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAllTransactionHandler mengembalikan daftar semua transaksi
func GetAllTransactionHandler(c echo.Context) error {
	// Logika untuk mengambil semua transaksi
	return c.JSON(http.StatusOK, "Get all transactions")
}

// GetTransactionHandler mengembalikan transaksi berdasarkan ID
func GetTransactionHandler(c echo.Context) error {
	id := c.Param("id")
	// Logika untuk mengambil transaksi berdasarkan ID
	return c.JSON(http.StatusOK, "Get transaction with ID: "+id)
}

// CreateTransactionHandler membuat transaksi baru
func CreateTransactionHandler(c echo.Context) error {
	// Logika untuk membuat transaksi
	return c.JSON(http.StatusOK, "Transaction created")
}

// UpdateTransactionHandler mengupdate transaksi berdasarkan ID
func UpdateTransactionHandler(c echo.Context) error {
	id := c.Param("id")
	// Logika untuk memperbarui transaksi berdasarkan ID
	return c.JSON(http.StatusOK, "Transaction updated with ID: "+id)
}

// DeleteTransactionHandler menghapus transaksi berdasarkan ID
func DeleteTransactionHandler(c echo.Context) error {
	id := c.Param("id")
	// Logika untuk menghapus transaksi berdasarkan ID
	return c.JSON(http.StatusOK, "Transaction deleted with ID: "+id)
}
