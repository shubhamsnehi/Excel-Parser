package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kcasctiv/dbf3"
)

func main() {

	f, err := os.Open("DIST_DTL.DBF")
	if err != nil {
		panic(err)
	}

	rc := bufio.NewReader(f)

	file, err := dbf3.Open(rc)
	if err != nil {
		panic(err)
	}

	// Get values
	fields := file.Fields()
	for idx := 0; idx < file.Rows(); idx++ {
		for _, field := range fields {
			value, _ := file.Get(idx, field.Name())
			fmt.Println(value)
		}
	}
}
