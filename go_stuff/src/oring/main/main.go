package main

import (
	"fmt"
	"ml/src/uap/gdm/mms"
	"ml/src/uap/gdm/myio"
	"ml/src/uap/mdls/reg/linreg"
)

func main() {
	//specify path
	path := "/home/pau/GolandProjects/ml/data/oring/o-ring-erosion-or-blowby.data"
	// read in data and handle errors
	inData, err := myio.ReadFileToSliceWhiteSpaceSep(path)
	if err != nil {
		fmt.Println(err)
	}

	myio.Print2dSliceHead(inData, 5)
	tdata := mms.MatrixTranspose(inData)
	myio.Print2dSliceHead(tdata, 5)

	y := tdata[1]
	var x [][]float64
	//add temp para
	x = append(x, tdata[2])
	//add pressure parameter
	x = append(x, tdata[3])
	fmt.Println("---------LinReg starts-----------\nx :", x, "\ny :", y)

	betaN, betaZero, err := linreg.LinearRegression(x, y)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(betaN)
	fmt.Println(betaZero)

	fmt.Println("-------------------LinReg ends-----------")

}
