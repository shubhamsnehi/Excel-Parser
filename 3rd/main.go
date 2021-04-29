package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	dbf "github.com/SebastiaanKlippert/go-foxpro-dbf"
)

func main() {
	//Creating CSV File
	csvFile, err := os.Create("./SaleDetails.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile) //Writer for writing CSV
	writer.Comma = '^'

	//DBF Reader
	dbf.SetValidFileVersionFunc(func(version byte) error {
		return nil
	})

	//Opening .DBF file
	testdbf, err := dbf.OpenFile("SALE_DTL.DBF", new(dbf.Win1250Decoder))
	if err != nil {
		log.Println(err)
	}
	defer testdbf.Close()

	//Writing Header
	var header []string
	for _, field := range testdbf.Fields() {
		header = append(header, field.FieldName())
	}
	writer.Write(header) //Writing Header

	//Writing rest Records
	noOfFields := len(testdbf.Fields())
	for !testdbf.EOF() {
		// for b := 0; b < 10; b++ {
		var details []string
		record, _ := testdbf.Record()
		for i := 0; i < noOfFields; i++ {
			rec, _ := record.Field(i)
			details = append(details, strings.TrimSpace(fmt.Sprint(rec)))
		}
		testdbf.Skip(1)
		// fmt.Println(details)
		writer.Write(details)
	}
	writer.Flush()
}
