package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kcasctiv/dbf3"
)

func main() {

	fileName := "SALE_DTL2.DBF"
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	rc := bufio.NewReader(f)

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
	fmt.Println(fileName, ":", header)

	// for idx := 0; idx < file.Rows(); idx++ {
	// 	for _, field := range fields {
	// 		value, _ := file.Get(idx, field.Name())
	// 		fmt.Println(value)
	// 	}
	// }
}
