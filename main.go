package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/option"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("command line args is not exist")
	}
	filename := os.Args[1]

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(context.Background(), option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_KEY"))))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	defer file.Close()

	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("failed to create image: %v", err)
	}

	text, err := client.DetectDocumentText(context.Background(), image, nil)
	if err != nil {
		log.Fatalf("failed to detect document text: %v", err)
	}
	fmt.Println(text.Text)
}
