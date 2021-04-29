package main

import (
	"fmt"
	"log"

	dbf "github.com/SebastiaanKlippert/go-foxpro-dbf"
)

func main() {
	// csvFile, err := os.Create("./DistDetails.csv")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer csvFile.Close()
	// writer := csv.NewWriter(csvFile)

	testdbf, err := dbf.OpenFile("DIST_DTL.DBF", new(dbf.Win1250Decoder))
	if err != nil {
		log.Println(err)
	}
	defer testdbf.Close()

	var header []string
	// noOfFields := len(testdbf.Fields())
	for _, field := range testdbf.Fields() {
		// fmt.Println(field.FieldName())
		header = append(header, field.FieldName())
	}
	// writer.Write(header)
	var details []string
	record, err := testdbf.Record()
	rec, _ := record.Field(4)
	a := rec.(string)
	fmt.Println(a)
	// for i := 0; i < noOfFields; i++ {
	// 	rec, _ := record.Field(i)
	// 	// fmt.Printf("%T\n",rec)
	// 	fmt.Printf("%T",rec)
	// 	// details = append(details, rec)
	// }
	fmt.Println(details)
	// writer.Flush()

	// dataArr, erro := ioutil.ReadFile("SALE_DTL.DBF") //Reading File
	// if erro != nil {
	// 	log.Println(erro)
	// }
	// // line := strings.Split(string(dataArr), "\n")
	// line := dataArr
	// fmt.Println(string(line))
	// for _, val := range line {
	// 	var row2 []string
	// 	row := strings.Split(string(val), " ")
	// 	for _, val2 := range row {
	// 		// fmt.Println(i, ":", strings.TrimSpace(val2))
	// 		row2 = append(row2, strings.TrimSpace(val2))
	// 	}
	// 	writer.Write(row2)
	// }
	// writer.Flush()
}
