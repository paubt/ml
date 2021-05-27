package main

import (
	"fmt"
	mms2 "ml/src/uap/gdm/mms"
	"ml/src/uap/gdm/myio"
	linreg2 "ml/src/uap/mdls/reg/linreg"
)

func main() {
	//specify path
	path := "/home/pau/GolandProjects/ml/data/oring/o-ring-erosion-or-blowby.data"
	// read in data and handle errors
	inData, err := myio.ReadFileToSlice(path)
	if err != nil {
		fmt.Println(err)
	}

	myio.Print2dSliceHead(inData, 5)
	tdata := mms2.TransposeMatrix(inData)
	myio.Print2dSliceHead(tdata, 5)

	y := tdata[1]
	var betaZero = 0
	var x [][]float64
	//add temp para
	x = append(x, tdata[2])
	//add pressure parameter
	x = append(x, tdata[3])
	fmt.Println("beta0 : ", betaZero, "\nx :", x, "\ny :", y)

	betaPara, err := linreg2.LinearRegression(x, y)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(betaPara)
	fmt.Println("-------------------LinReg ends-----------")

}
