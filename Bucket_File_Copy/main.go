package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/kcasctiv/dbf3"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:\\Users\\Shubham Snehi\\Downloads\\awacs-dev-160bf0e57dc1.json")

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name for the new bucket.
	bucketName := "balatestawacs"

	// Creates a Bucket instance.
	bucket := client.Bucket(bucketName)

	obj := bucket.Object("SALE_DTL.DBF") // see https://developer.bestbuy.com/apis

	rdr, err := obj.NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer rdr.Close()

	rc := bufio.NewReader(rdr)

	file, err := dbf3.Open(rc)
	if err != nil {
		panic(err)
	}

	err = file.SaveFile("Copied.DBF")
	log.Println("Here:", err)

}
