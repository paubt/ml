package main

import (
	"fmt"
	"ml/src/uap/gdm/mms"
	"ml/src/uap/gdm/myio"
	"ml/src/uap/mdls/reg/linreg"
)

func main() {
	//specify path
	path := "/home/pau/GolandProjects/ml/data/brainbodyratio/brainbodyratio.txt"
	//read in the data and handle the errors
	inData, err := myio.ReadFileToSliceWhiteSpaceSep(path)
	if err != nil {
		fmt.Println(err)
	}
	//transpose data to split for easy split
	inDatat := mms.MatrixTranspose(inData)
	//take the second column as list (parameter) and append to x to get a 2d array
	var x [][]float64
	x = append(x, inDatat[1])
	//take the third column as y (target)
	var y []float64 = inDatat[2]
	//print input data
	fmt.Println("-----the input data-----\n", "x :\n", x, "\ny :\n", y)
	//call linear regression function
	fmt.Println("-----lin reg begins-----")
	betaN, betaZero, err := linreg.LinearRegression(x, y)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(betaN)
	fmt.Println(betaZero)
	fmt.Println("-----lin reg ends-----")

}
