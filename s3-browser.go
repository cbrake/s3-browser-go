package main

import (
	//"net/http"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getS3Objects(bucket string) ([]string, error) {
	svc := s3.New(session.New(&aws.Config{Region: aws.String("us-west-2")}))

	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	}

	resp, err := svc.ListObjects(params)

	if err != nil {
		return []string{}, err
	} else {
		if (*resp.IsTruncated) {
			fmt.Println("Warning, response is truncated -- need to improve your program")
		} else {
			fmt.Println("Full response")
		}

		var ret []string

		for _, key := range resp.Contents {
			ret = append(ret, *key.Key)
		}

		return ret, nil
	}
}

func main() {
	// Simple static webserver:
	//log.Fatal(http.ListenAndServe(":9090", http.FileServer(http.Dir("/usr/share/doc"))))

	objs, err := getS3Objects("blockeng-internal-releases")

	if err != nil {
		fmt.Println("Error retrieving bucket objects: ", err)
	} else {
		for _, obj := range objs {
			fmt.Println(obj)
		}
	}
}

