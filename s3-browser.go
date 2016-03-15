package main

import (
	"net/http"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":9090", http.FileServer(http.Dir("/usr/share/doc"))))

	svc := s3.New(session.New())
	svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("bucketName"),
		Key: aws.String("keyName"),
	})
}

