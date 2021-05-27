package linreg

import (
	mms2 "ml/src/uap/gdm/mms"
)

func LinearRegression(x [][]float64, y []float64) ([]float64, error) {
	//transpose x
	xt := mms2.TransposeMatrix(x)
	//matrix multiplication of xt with x
	xtx := mms2.MatrixMultiplication(x, xt)
	//calculation of the invertible of xtx
	//if error then pass it on an terminate
	xtxi, err := mms2.MatrixInvertible(xtx)
	if err != nil {
		return nil, err
	}
	//transform the list ([n]) of y into a array of 2d with only one row ([1][n])
	var y2d [][]float64
	y2d = append(y2d, y)
	//transpose the "2d" y array
	yt := mms2.TransposeMatrix(y2d)
	//matrix multiplication of x and yt
	xty := mms2.MatrixMultiplication(x, yt)
	//final matrix multiplication of xtxi and xty to get the lin reg parameters
	betaHat := mms2.MatrixMultiplication(xtxi, xty)
	//transform 2d slice betaHat to 1d slice for output
	var res []float64
	for i := 0; i < len(betaHat); i++ {
		for j := 0; j < len(betaHat[0]); j++ {
			res = append(res, betaHat[i][j])
		}
	}
	return res, nil
}
