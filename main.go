package main

import (
	"fmt"
	"math"
	_ "math"
	"strings"
	_ "strings"
)

func encryption(s string) string {
	var restOfColumns []string
	var lastColumn []string

	rows, columns := returnNumb(s)
	matrix := returnMatrix(s, rows, columns)

	for d, _ := range matrix {
		for i := 0; i < rows; i++ {
			for j := 0; j < columns; j++ {
				if d == j {
					restOfColumns = append(restOfColumns, matrix[i][j])
				}
			}
		}
		restOfColumns = append(restOfColumns, " ")
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if j == rows {
				lastColumn = append(lastColumn, matrix[i][j])
			}
		}
	}

	restOfString :=  strings.Join(restOfColumns, "")
	lastString := strings.Join(lastColumn, "")
	newString := restOfString+lastString
	return newString
}

func returnNumb(str string) (int, int){
	var rows = int(math.Sqrt(float64(len(str))))
	fmt.Println(rows)
	var columns = rows + 1
	if rows * columns < len(str){
		rows++
	}
	if rows * rows == len(str){
		return rows, rows
	}
	return rows, columns
}

func returnMatrix(str string, rows int, columns int) [][]string {
	matrix := make([][]string, rows)
	for i := range matrix {
		matrix[i] = make([]string, columns)
	}
	incr := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if incr < len(str){
				matrix[i][j] = string(str[incr])
				incr++
			}
		}
	}
	return matrix
}


func main(){
	str := encryption("iffactsdontfittotheorychangethefacts")
	fmt.Println(str)
}