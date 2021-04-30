package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/kcasctiv/dbf3"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:\\Users\\Shubham Snehi\\Downloads\\awacs-dev-160bf0e57dc1.json")

	csvFile, err := os.Create("./OrderDetails.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)

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

	obj := bucket.Object("Excel1.xlsx") // see https://developer.bestbuy.com/apis

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


	line := strings.Split(string(dataArr), "\n")

	for _, val := range line {
		var row2 []string
		row := strings.Split(string(val), ";")
		for _, val2 := range row {
			// fmt.Println(i,":",strings.TrimSpace(val2))
			row2 = append(row2, strings.TrimSpace(val2))
		}
		writer.Write(row2)
	}
	writer.Flush()
}
