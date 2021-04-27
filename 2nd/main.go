package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	csvFile, err := os.Create("./OrderDetails.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)

	dataArr, erro := ioutil.ReadFile("Sale_Dtl.txt") //Reading File
	if erro != nil {
		log.Println(erro)
	}
	line := strings.Split(string(dataArr),"\n")

	for _,val := range line{
		var row2 []string
		row := strings.Split(string(val), ";")
		for _,val2 := range row{
			// fmt.Println(i,":",strings.TrimSpace(val2))
			row2 = append(row2, strings.TrimSpace(val2))
		}
		writer.Write(row2)
	}
	writer.Flush()
}
