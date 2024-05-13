package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract the file extension
	fileExtension := filepath.Ext(file.Filename)

	// Remove the extension from the original filename
	originalFilenameWithoutExt := file.Filename[0 : len(file.Filename)-len(fileExtension)]

	// Create a new UUID for the stored filename without extension
	uuidPart := uuid.New().String()
	storedFileNameWithoutExt := uuidPart
	storedFileName := uuidPart + fileExtension

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	// Upload the file to MinIO
	_, err = MinioClient.PutObject(c, Conf.Minio.BucketName, storedFileName, src, file.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Insert file metadata into the database
	_, err = DB.Exec("INSERT INTO asset (original_filename, stored_filename, file_extension, filesize) VALUES ($1, $2, $3, $4)",
		originalFilenameWithoutExt, storedFileNameWithoutExt, fileExtension, file.Size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the ID of the new file (using stored filename as ID here)
	c.JSON(http.StatusOK, gin.H{"id": storedFileName})
}

func DownloadFile(c *gin.Context) {
	id := c.Param("id")
	var originalFilename, storedFilename, file_extension string

	// Retrive the original filename from the database
	err := DB.QueryRow("SELECT original_filename, stored_filename, file_extension FROM asset WHERE stored_filename = $1", id).Scan(&originalFilename, &storedFilename, &file_extension)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the object using minioClient
	object, err := MinioClient.GetObject(c, Conf.Minio.BucketName, storedFilename+file_extension, minio.GetObjectOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer object.Close()

	// Stream the object to the client
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", originalFilename+file_extension))
	_, err = io.Copy(c.Writer, object)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func ListFiles(c *gin.Context) {
	// Get page and limit from the query string parameters
	page, err := strconv.Atoi(c.DefaultQuery("page", "1")) // Default to page 1 if not specified
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", Conf.Database.DefaultQueryLimit)) // Default to Config value per page if not specified
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Prepare SQL query using limit and offset
	rows, err := DB.Query("SELECT id, original_filename, stored_filename, file_extension, filesize FROM asset ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	files := []gin.H{}
	for rows.Next() {
		var original_filename, stored_filename, file_extension string
		var id, filesize int64
		if err := rows.Scan(&id, &original_filename, &stored_filename, &file_extension, &filesize); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		files = append(files, gin.H{
			"id":                id,
			"original_filename": original_filename,
			"stored_filename":   stored_filename,
			"file_extension":    file_extension,
			"size":              filesize,
		})
	}
	c.JSON(http.StatusOK, files)
}
