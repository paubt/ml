package main

import (
	"fmt"
	"ml/src/gdm"
	"ml/src/oring/myio"
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
	tdata := gdm.TransposeMatrix(inData)
	myio.Print2dSliceHead(tdata, 5)

	y := tdata[1]
	var betaZero = 0
	var x [][]float64
	//add temp para
	x = append(x, tdata[2])
	//add pressure parameter
	x = append(x, tdata[3])
	fmt.Println("beta0 : ", betaZero, "\nx :", x, "\ny :", y)
	xt := gdm.TransposeMatrix(x)
	fmt.Println(xt)
	xtx := gdm.MatrixMultiplication(x, xt)
	myio.Print2dSlice(xtx)

	fmt.Println("---------Invertibel----------")
	xtxi, err := gdm.MatrixInvertible(xtx)
	if err != nil {
		fmt.Println(err)
	}
	myio.Print2dSlice(xtxi)

	var y2d [][]float64
	y2d = append(y2d, y)
	myio.Print2dSlice(y2d)
	yt := gdm.TransposeMatrix(y2d)
	xty := gdm.MatrixMultiplication(x, yt)
	myio.Print2dSlice(xty)

	betaHat := gdm.MatrixMultiplication(xtxi, xty)
	fmt.Println("betaHat matrix")
	myio.Print2dSlice(betaHat)

	/* test matrix multiplication
	fmt.Printf("\n----------------test----------------\n")
	a := [][]int{{1, 2, 1}, {3, 4, 4}}
	b := [][]int{{2, 3}, {4, 1}, {3, 3}}
	fmt.Println("a", a, "\nb", b)
	c := gdm.MatrixMultiplication(a, b)
	fmt.Println(c)
	*/
	/* test matrix determinant
	fmt.Println("----------test1--------------")
	test := [][]int{{1, 0, 2, -1},
		{3, 1, 0, 5},
		{2, 1, 4, -3},
		{1, 0, 5, 0}}
	dtest, err := gdm.MatrixDet(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dtest)

	fmt.Println("----------test2--------------")
	test2 := [][]int{{1, 0, 2, -1, 6},
		{3, 1, 0, 5, 9},
		{2, 1, 4, -3, 8},
		{4, 13, 0, 2, 1},
		{1, 0, 5, 0, 10}}
	dtest2, err := gdm.MatrixDet(test2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dtest2)
	*/
	/* matrix inverible test
	fmt.Println("--------------test--------------")
	test3 := [][]float64{{5, -2, 2, 7},
		{1, 0, 0, 3},
		{-3, 1, 5, 0},
		{3, -1, -9, 4}}
	tempAJ := gdm.MatrixAdjoint(test3)
	myio.Print2dSlice(tempAJ)
	inver , err := gdm.MatrixInvertible(test3)
	if err != nil {
		fmt.Println(err)
	}
	myio.Print2dSlice(inver)

	*/
}
