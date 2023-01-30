package infrastructure

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func InitStorage() error {
	endpoint := os.Getenv("minio_endpoint")
	user := os.Getenv("minio_user")
	password := os.Getenv("minio_password")
	useSsl := false

	_minioClient, err := minio.New(endpoint, &minio.Options{
		Secure: useSsl,
		Creds:  credentials.NewStaticV4(user, password, ""),
	})
	if err != nil {
		return err
	}
	minioClient = _minioClient
	return nil
}

func GetClient() *minio.Client {
	return minioClient
}

func CreateBucket(bucketName string) error {
	err := minioClient.MakeBucket(context.Background(), "payment_proof", minio.MakeBucketOptions{
		Region: "us-east-1",
	})

	if err != nil {
		exist, existErr := minioClient.BucketExists(context.Background(), "payment_proof")
		if exist && existErr == nil {
			log.Default().Printf("bucket %v already exist", "payment_proof")
			return nil
		} else {
			return err
		}
	}
	return nil
}
