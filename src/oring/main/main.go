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
	var betaZero int = 0
	var x [][]int
	//add temp para
	x = append(x, tdata[1])
	//add pressure parameter
	x = append(x, tdata[2])
	fmt.Println("beta0 : ", betaZero, "\nx :", x, "\ny :", y)
	xt := gdm.TransposeMatrix(x)
	fmt.Println(xt)
	xtx := gdm.MatrixMultiplication(x, xt)
	fmt.Println(xtx)

	fmt.Printf("\n----------------test----------------\n")
	a := [][]int{{1, 2, 1}, {3, 4, 4}}
	b := [][]int{{2, 3}, {4, 1}, {3, 3}}
	fmt.Println("a", a, "\nb", b)
	c := gdm.MatrixMultiplication(a, b)
	fmt.Println(c)
	d, err2 := gdm.MatrixDet(c)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(d)
	test := [][]int{{1, 0, 2, -1},
		{3, 0, 0, 5},
		{2, 1, 4, -3},
		{1, 0, 5, 0}}
	dtest, err := gdm.MatrixDet(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dtest)

}
