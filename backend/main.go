package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"backend/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

// Middleware CORS khusus mengizinkan origin Vite dev server Vue (misal http://localhost:5173 atau http://localhost:5174, dll)
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Periksa apakah origin berasal dari dev server Vue (misal localhost:5173 atau 127.0.0.1:5173)
		isAllowed := false
		if origin != "" {
			if strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "http://127.0.0.1:") {
				isAllowed = true
			}
		}

		if isAllowed {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	// Buat folder uploads jika belum ada
	if err := os.MkdirAll("uploads", 0755); err != nil {
		log.Printf("Peringatan pembuatan direktori uploads: %v", err)
	}

	// Baca konfigurasi dari .env
	if err := godotenv.Load(); err != nil {
		log.Println("Peringatan: File .env tidak ditemukan atau gagal dimuat, menggunakan default/environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data/app.db"
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	r := gin.Default()
	r.Use(corsMiddleware())

	// Layani file statis agar file unggahan di folder uploads dapat diakses publik melalui browser
	r.Static("/uploads", "./uploads")

	serviceHandler := handlers.NewServiceHandler(db)

	// Endpoint persis sesuai blueprint.md
	r.GET("/api/services", serviceHandler.GetServices)
	r.GET("/api/services/:id", serviceHandler.GetServiceByID)
	r.POST("/api/services", serviceHandler.CreateService)
	r.PUT("/api/services/:id", serviceHandler.UpdateService)
	r.DELETE("/api/services/:id", serviceHandler.DeleteService)

	fmt.Printf("Server berjalan di port :%s (Database: %s)\n", port, dbPath)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
