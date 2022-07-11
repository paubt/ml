package linreg

import (
	mms2 "ml/src/uap/gdm/mms"
)

func LinearRegression(x [][]float64, y []float64) ([]float64, float64, error) {
	//transpose x
	xt := mms2.MatrixTranspose(x)
	//matrix multiplication of xt with x
	xtx := mms2.MatrixMultiplication(x, xt)
	//calculation of the invertible of xtx
	//if error then pass it on an terminate
	xtxi, err := mms2.MatrixInvertible(xtx)
	if err != nil {
		return nil, 0, err
	}
	//transform the list ([n]) of y into a array of 2d with only one row ([1][n])
	var y2d [][]float64
	y2d = append(y2d, y)
	//transpose the "2d" y array
	yt := mms2.MatrixTranspose(y2d)
	//matrix multiplication of x and yt
	xty := mms2.MatrixMultiplication(x, yt)
	//final matrix multiplication of xtxi and xty to get the lin reg parameters
	betaHat := mms2.MatrixMultiplication(xtxi, xty)
	//transform 2d slice betaHat to 1d slice for output
	var betaHatn []float64
	for i := 0; i < len(betaHat); i++ {
		for j := 0; j < len(betaHat[0]); j++ {
			betaHatn = append(betaHatn, betaHat[i][j])
		}
	}
	// calculate the mean of Y
	var meanY float64
	for _, v := range y {
		meanY += v
	}
	meanY = meanY / float64(len(y))
	// calc the the mean for each x parameter and subtract it form the mean of y
	// for all x (columns) in X
	var betaHatZero, sumCol, meanCol float64
	for i, col := range x {
		sumCol = 0
		for _, v := range col {
			sumCol += v
		}
		meanCol = sumCol / float64(len(col))
		betaHatZero = meanY - betaHatn[i]*meanCol
	}
	return betaHatn, betaHatZero, nil
}
