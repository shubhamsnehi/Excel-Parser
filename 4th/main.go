package main

import (
	"fmt"
	"log"

	"github.com/kcasctiv/dbf3"
)

func main() {

	file, err := dbf3.OpenFile("DIST_DTL.DBF")
	if err != nil{
		log.Fatalln(err)
	}

	// file := dbf3.New(langDriver)

	// Change language driver
	// file.SetLang(newDriver)

	// Get values
	fields := file.Fields()
	for idx := 0; idx < file.Rows(); idx++ {
		for _, field := range fields {
			value, _ := file.Get(idx, field.Name())
			fmt.Println(value)
		}
	}
}
