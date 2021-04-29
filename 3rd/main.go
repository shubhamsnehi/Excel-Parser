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
	csvFile, err := os.Create("./SaleDetails.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)

	testdbf, err := dbf.OpenFile("SALE_DTL.DBF", new(dbf.Win1250Decoder))
	if err != nil {
		log.Println(err)
	}
	defer testdbf.Close()

	var header []string
	noOfFields := len(testdbf.Fields())
	for _, field := range testdbf.Fields() {
		header = append(header, field.FieldName())
	}
	writer.Write(header)
	var details []string
	record, err := testdbf.Record()

	for i := 0; i < noOfFields; i++ {
		rec, _ := record.Field(i)
		details = append(details, strings.TrimSpace(fmt.Sprint(rec)))
	}
	writer.Write(details)
	writer.Flush()
}
