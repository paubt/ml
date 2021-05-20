package myio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadFileToSlice read file and return it as a int slice
func ReadFileToSlice(filePath string) ([][]int, error) {
	//open file
	file, err := os.Open(filePath)
	//close file when exiting main
	defer file.Close()
	//if an error occurs throw exception
	if err != nil {
		return nil, err
	}
	//create a new scanner to go through the file
	scanner := bufio.NewScanner(file)
	//empty output 2d slice
	var inputData [][]int
	//go through the file row by row
	for scanner.Scan() {
		//take a row and split it by the whitespaces into a slice
		res1 := strings.Fields(scanner.Text())
		//temporal slice to store converted values
		var tempRowSlice []int
		//go through the sliced row, convert each value from string to int
		//and append it to the temporal slice
		for _, val := range res1 {
			i, _ := strconv.Atoi(val)
			tempRowSlice = append(tempRowSlice, i)
		}
		//append the temporal slice of a row to the 2d output slice
		inputData = append(inputData, tempRowSlice)
	}
	return inputData, nil
}

func Print2dSlice(data [][]int) {
	for i, row := range data {
		fmt.Printf("%d ", i)
		fmt.Println(row)
	}
}

func Print2dSliceHead(data [][]int, rows int) {
	for i, row := range data {
		if i < rows {
			fmt.Printf("%d ", i)
			fmt.Println(row)
		}
	}
}
