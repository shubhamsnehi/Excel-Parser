package main

import (
	"bufio"
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"github.com/kcasctiv/dbf3"
	"google.golang.org/api/iterator"
)

func main() {

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
	query := &storage.Query{}
	it := bucket.Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println(attrs.Name)
	}

	obj := bucket.Object("SALE_DTL.DBF").ReadCompressed(true) // see https://developer.bestbuy.com/apis

	rdr, err := obj.NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer rdr.Close()

	// fileName := "SALE_DTL2.DBF"
	// f, err := os.Open(fileName)
	// if err != nil {
	// 	panic(err)
	// }

	rc := bufio.NewReader(rdr)

	file, err := dbf3.Open(rc)
	if err != nil {
		panic(err)
	}

	// Get values
	var header []string
	// fields := file.Fields()
	//Headers
	for _, field := range file.Fields() {
		header = append(header, field.Name())
	}
	fmt.Println(":", header)

	// for idx := 0; idx < file.Rows(); idx++ {
	// 	for _, field := range fields {
	// 		value, _ := file.Get(idx, field.Name())
	// 		fmt.Println(value)
	// 	}
	// }
}
