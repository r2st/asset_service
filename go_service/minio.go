package main

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

// InitializeMinioClient sets up the MinIO client and is meant to be called once at the start of the application.
func InitializeMinioClient() {
	var err error

	MinioClient, err = minio.New(Conf.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(Conf.Minio.AccessKey, Conf.Minio.SecretKey, ""),
		Secure: Conf.Minio.Secure,
	})
	if err != nil {
		log.Fatalf("Failed to initialize MinIO client: %v", err)
	}
	log.Println("MinIO client initialized successfully")

	ctx := context.Background() // Create a context for MinIO operations

	// Ensure that the bucket exists or create it
	exists, err := MinioClient.BucketExists(ctx, Conf.Minio.BucketName)
	if err == nil && !exists {
		err = MinioClient.MakeBucket(ctx, Conf.Minio.BucketName, minio.MakeBucketOptions{
			Region:        Conf.Minio.BucketRegion,
			ObjectLocking: Conf.Minio.BucketObjectLocking,
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// GetMinioClient returns the initialized MinIO client
func GetMinioClient() *minio.Client {
	if MinioClient == nil {
		log.Fatal("MinIO client is not initialized")
	}
	return MinioClient
}
