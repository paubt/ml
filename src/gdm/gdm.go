package gdm

import (
	"errors"
)

// TransposeMatrix row 2d slice to column 2d slice
func TransposeMatrix(matrix [][]float64) [][]float64 {
	columns := len(matrix)
	rows := len(matrix[0])
	var tempTuple []float64
	var matrixT [][]float64
	for c := 0; c < rows; c++ {
		tempTuple = nil
		for r := 0; r < columns; r++ {
			tempTuple = append(tempTuple, matrix[r][c])
		}
		matrixT = append(matrixT, tempTuple)
	}
	return matrixT
}

func MatrixMultiplication(a [][]float64, b [][]float64) [][]float64 {
	var outta [][]float64
	var inner []float64
	var sum float64
	//fmt.Println("i :", len(a), "\nj :", len(b[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			sum = 0
			//fmt.Println(i, " ", j)
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

func getCofactor(mat [][]float64, igrow int, igcol int) [][]float64 {
	var i int = 0
	var j int = 0
	var inner []float64
	var temp [][]float64
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

func MatrixDet(a [][]float64) (float64, error) {
	// throw errors in specific cases
	if len(a) == 0 {
		err := errors.New("length is zero ---> error")
		return 0, err
	}
	if len(a) != len(a[0]) {
		err := errors.New("matrix not square")
		return 0, err
	}
	//Base Cases with n==1 and n==2
	if len(a) == 1 {
		return a[0][0], nil
	}
	if len(a) == 2 {
		sum := a[0][0]*a[1][1] - a[1][0]*a[0][1]
		//print for debug
		//fmt.Println(sum)
		return sum, nil
	}
	//init array to store cofactors
	var temp [][]float64
	// sign multiplier
	var sign = 1
	// d is temp storage for determinant
	var d float64
	//iterate through each element in first row
	for i := 0; i < len(a); i++ {
		temp = getCofactor(a, i, 0)
		e, err := MatrixDet(temp)
		if err != nil {
			return 0, err
		}
		//print for debug
		//fmt.Printf("%d + %d * %d * %d = ", d, sign, a[i][0], e)

		d = d + float64(sign)*a[i][0]*e
		//print for debug
		//fmt.Printf("%d\n", d)
		sign = -sign
	}
	return d, nil
}

func MatrixAdjoint(a [][]float64) [][]float64 {
	// init variables
	var temp [][]float64
	var inner []float64
	var res [][]float64
	sign := 1
	// iterate though the input matrix
	for i := 0; i < len(a); i++ {
		inner = nil
		for j := 0; j < len(a[0]); j++ {
			// get the cofactor matrix
			temp = getCofactor(a, i, j)
			// calculate the determinant of the cofactor matrix
			tempD, _ := MatrixDet(temp)
			// ad sign factor
			tempDwithSign := float64(sign) * tempD
			// append det to inner slice
			inner = append(inner, tempDwithSign)
			// flip sign for next column
			sign = -sign
		}
		// flip sign for next row
		sign = -sign
		// append the inner slice to the res 2d slice
		res = append(res, inner)
	}
	//transpose the res
	res = TransposeMatrix(res)
	return res
}

func MatrixInvertible(a [][]float64) ([][]float64, error) {
	var res [][]float64
	var inner []float64
	aDet, err := MatrixDet(a)
	if err != nil {
		return nil, err
	}

	adjoint := MatrixAdjoint(a)

	for i := 0; i < len(a); i++ {
		inner = nil
		for j := 0; j < len(a); j++ {
			inner = append(inner, adjoint[i][j]/aDet)
		}
		res = append(res, inner)
	}
	return res, nil
}
