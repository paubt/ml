package gdm

import (
	"errors"
	"fmt"
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

func getCofactor(mat [][]int, temp [][]int, p int, q int) [][]int {
	var i int = 0
	var j int = 0
	var inner []int
	for row := 0; row < len(mat); row++ {
		inner = nil
		for col := 0; col < len(mat); col++ {
			if row != p && col != q {
				inner = append(inner, mat[row][col])
				j++
				if j == len(mat)-1 {
					j = 0
					i++
				}
			}
		}
		temp = append(temp, inner)
	}
	return temp
}

func MatrixDet(a [][]int) (int, error) {
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
		temp = getCofactor(a, temp, 0, i)
		e, err := MatrixDet(temp)
		if err != nil {
			return 0, err
		}
		d = d + sign*a[0][i]*e
		sign = -sign
	}

	return d, nil
}

func MatrixInvertible(a [][]int) [][]int {

	return a
}
