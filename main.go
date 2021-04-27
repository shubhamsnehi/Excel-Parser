package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	// //Creating
	// f := excelize.NewFile()
	// // Create a new sheet.
	// index := f.NewSheet("Sheet2")
	// // Set value of a cell.
	// f.SetCellValue("Sheet2", "A2", "Hello world.")
	// f.SetCellValue("Sheet1", "B2", 100)
	// // Set active sheet of the workbook.
	// f.SetActiveSheet(index)
	// // Save spreadsheet by the given path.
	// if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }

	f, err := excelize.OpenFile("Excel1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// // Get value from cell by given worksheet name and axis.
	// cell, err := f.GetCellValue("Sheet1", "B2")
	// if err != nil {
	//     fmt.Println(err)
	//     return
	// }
	// fmt.Println(cell)

	// Get all the rows in the Sheet1.

	csvFile, err := os.Create("./OrderDetails.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)

	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		var row2 []string
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
			row2 = append(row2, colCell)
		}
		fmt.Println()
		writer.Write(row2)
	}
	// remember to flush!
	writer.Flush()
}
