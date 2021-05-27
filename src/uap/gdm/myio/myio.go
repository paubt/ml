package myio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadFileToSliceWhiteSpaceSep read file and return it as a int slice
func ReadFileToSliceWhiteSpaceSep(filePath string) ([][]float64, error) {
	//open file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	//close file when exiting main
	defer file.Close()

	//create a new scanner to go through the file
	scanner := bufio.NewScanner(file)
	//empty output 2d slice
	var inputData [][]float64
	//go through the file row by row
	for scanner.Scan() {
		//take a row and split it by the whitespaces into a slice
		res1 := strings.Fields(scanner.Text())
		//temporal slice to store converted values
		var tempRowSlice []float64
		//go through the sliced row, convert each value from string to int
		//and append it to the temporal slice
		for _, val := range res1 {
			//transform string value to float value
			// int solution: i, _ := strconv.Atoi(val)
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}
			//append parsed float to temporal row
			tempRowSlice = append(tempRowSlice, f)
		}
		//append the temporal slice of a row to the 2d output slice
		inputData = append(inputData, tempRowSlice)
	}
	return inputData, nil
}

func Print2dSlice(data [][]float64) {
	for i, row := range data {
		fmt.Printf("%d ", i)
		fmt.Println(row)
	}
}

func Print2dSliceHead(data [][]float64, rows int) {
	for i, row := range data {
		if i < rows {
			fmt.Printf("%d ", i)
			fmt.Println(row)
		}
	}
}
