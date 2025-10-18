package utils

import "math/big"

func GetFactorial(number int64) *big.Int {
	if number == 1 {
		return big.NewInt(1)
	}
	result := new(big.Int)
	result.Mul(big.NewInt(number), GetFactorial(number-1))
	return result
}
