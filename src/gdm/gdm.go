package gdm

import (
	"errors"
	"fmt"
	"ml/src/oring/myio"
)

// TransposeMatrix row 2d slice to column 2d slice
func TransposeMatrix(matrix [][]int) [][]int {
	columns := len(matrix)
	rows := len(matrix[0])
	var tempTuple []int
	var matrixT [][]int
	for c := 0; c < rows; c++ {
		tempTuple = nil
		for r := 0; r < columns; r++ {
			tempTuple = append(tempTuple, matrix[r][c])
		}
		matrixT = append(matrixT, tempTuple)
	}
	return matrixT
}

func MatrixMultiplication(a [][]int, b [][]int) [][]int {
	var outta [][]int
	var inner []int
	fmt.Println("i :", len(a), "\nj :", len(b[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			sum := 0
			fmt.Println(i, " ", j)
			for k := 0; k < len(b); k++ {
				sum = sum + a[i][k]*b[k][j]
			}
			inner = append(inner, sum)
		}
		outta = append(outta, inner)
		inner = nil
	}
	return outta
}

func getCofactor(mat [][]int, igrow int, igcol int) [][]int {
	var i int = 0
	var j int = 0
	var inner []int
	var temp [][]int
	for row := 0; row < len(mat); row++ {
		inner = nil
		for col := 0; col < len(mat); col++ {
			if row != igrow && col != igcol {
				inner = append(inner, mat[row][col])
				j++
				if j == len(mat)-1 {
					j = 0
					i++
				}
			}
		}
		if inner != nil {
			temp = append(temp, inner)
		}
	}
	return temp
}

func MatrixDet(a [][]int) (int, error) {
	myio.Print2dSlice(a)
	fmt.Printf("\n")
	if len(a) == 0 {
		err := errors.New("length is zero ---> error")
		return 0, err
	}
	//Base Cases with n==1 and n==2
	if len(a) == 1 {
		return a[0][0], nil
	}
	if len(a) == 2 {
		sum := a[0][0]*a[1][1] - a[1][0]*a[0][1]
		fmt.Println(sum)
		return sum, nil
	}
	//init array to store cofactors
	var temp [][]int
	// sign multiplier
	var sign = 1
	// d is temp storage for determinant
	var d int
	//iterate through each element in first row
	for i := 0; i < len(a); i++ {
		temp = getCofactor(a, i, 0)
		e, err := MatrixDet(temp)
		if err != nil {
			return 0, err
		}
		fmt.Printf("%d + %d * %d * %d = ", d, sign, a[0][i], e)
		d = d + sign*a[0][i]*e
		fmt.Printf("%d\n", d)
		sign = -sign
	}
	return d, nil
}

func MatrixInvertible(a [][]int) [][]int {

	return a
}
