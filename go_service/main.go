package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	LoadConfig() // Load configuration

	// Setup database connection
	SetupDatabase()

	// Initialize MinIO client
	InitializeMinioClient()

	// Set up Gin server
	r := gin.Default()
	maxAge, err := time.ParseDuration(Conf.Server.Cors.MaxAge)
	if err != nil {
		log.Fatalf("Error parsing duration (MaxAge): %v", err)
	}
	corsConfig := cors.Config{
		AllowOrigins:     Conf.Server.Cors.AllowOrigins,
		AllowMethods:     Conf.Server.Cors.AllowMethods,
		AllowHeaders:     Conf.Server.Cors.AllowHeaders,
		ExposeHeaders:    Conf.Server.Cors.ExposeHeaders,
		AllowCredentials: Conf.Server.Cors.AllowCredentials,
		MaxAge:           maxAge,
	}
	r.Use(cors.New(corsConfig))

	// API routes
	r.POST("/upload", UploadFile)
	r.GET("/files", ListFiles)
	r.GET("/download/:id", DownloadFile)

	// Start server
	r.Run(":" + Conf.Server.Port) // Use configured port
}
