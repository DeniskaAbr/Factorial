/*
read https://github.com/PeterLuschny/Fast-Factorial-Functions/ and
https://oeis.org/A056040
for more faster algorithms
*/

package main

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

var one = big.NewInt(1)

func main() {
	userInputValue, err := GetUserInput()
	s := ""

	if err != nil {
		if err.Error() == "big value for uint64 capacity" {
			s = BigLoopFactorial(userInputValue)
		}

	} else {
		factorial, err := ComputeFactorialForValue(userInputValue)
		if err != nil {

		} else {
			fmt.Printf("Computed factorial for %v is: %v \n", userInputValue, factorial)
		}
		s = strconv.FormatUint(factorial, 10)
	}

	countStr, err := TrailingZeroForStringifiedFactorialResultat(s)
	if err == nil {

		fmt.Println("Trailing zeros count from stringified factorial value? is: ", countStr)
	}

	count, err := TailingZerosCountsOptimizedWithNoComputeFactorial(userInputValue)

	if err == nil {
		fmt.Println("Trailing zeros count is: " + strconv.FormatUint(count, 10))
	}

}

func GetUserInput() (uint64, error) {
	var userInputValue uint64
	fmt.Println("Please enter positive integer: ")
	fmt.Scanln(&userInputValue)

	if userInputValue == 0 {
		fmt.Println("Default value is zero. Factorial from zero is: 1")
		return 0, errors.New("zero value")
	}
	if userInputValue < 0 {
		fmt.Println("Factorial for negative values does not exist")
		return 0, errors.New("negative value")
	}
	if userInputValue >= 66 {
		fmt.Println("Factorial for inputted value is so bigger int64 capacity, " +
			"use other factorial compute methods and value presentation.")
		return userInputValue, errors.New("big value for uint64 capacity")
	}
	return userInputValue, nil
}

func ComputeFactorialForValue(value uint64) (uint64, error) {
	return LoopFactorial(value), nil
}

// not optimized, small values limit factorial compute (loop)
func LoopFactorial(n uint64) uint64 {
	value := uint64(1)
	for i := uint64(1); i <= n; i++ { // 65 limit
		// golang not panic when value type is overflow
		// use big or checking overflow with https://github.com/JohnCGriffin/overflow
		// read in https://roylee0704.medium.com/overflowing-constants-93b0cb567a0d
		// https://stackoverflow.com/questions/33641717/detect-signed-int-overflow-in-go
		// https://stackoverflow.com/questions/30833314/golang-overflows-int64
		value *= uint64(i)
	}

	return value
}

// for big values not optimized, small values limit factorial compute (loop)
func BigLoopFactorial(n uint64) string {

	start := big.NewInt(1)

	end := new(big.Int).SetUint64(n)

	value := big.NewInt(1)
	fmt.Println()
	fmt.Println("Please wait long time ")
	// INFORMATION FROM https://stackoverflow.com/questions/29930501/big-int-ranges-in-go
	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, one) {
		value.Mul(value, i)
		// fmt.Println(i)
	}

	fmt.Println("Big factorial is: " + value.String())
	return value.String()
}

/*
optimized trailing zeros counter
https://www.geeksforgeeks.org/golang-program-to-count-trailing-zeros-in-factorial-of-a-number/
https://medium.com/@ezekielphlat/count-trailing-zeros-in-a-factorial-using-golang-933262a5f5fa
To calculate Trailing zeros in n!
Tailing n! = n/5 + n/5*5 + n/5*5*5...............r<1where "r" is the result of each iteration of n/divisor
*/

func TailingZerosCountsOptimizedWithNoComputeFactorial(value uint64) (uint64, error) {

	divisor := big.NewInt(5)
	count := big.NewInt(0)
	bigvalue := new(big.Int).SetUint64(value)

	fmt.Println("Compute trailing zero count...")

	for true {
		temp := new(big.Int).Div(bigvalue, divisor)
		if temp.Cmp(one) == -1 {
			break
		} else {
			count = new(big.Int).Add(count, temp)
			divisor = divisor.Mul(divisor, new(big.Int).SetInt64(5))
		}
	}
	// TODO: not equal value, maybe big Type error?
	return count.Uint64(), nil

}

func TrailingZeroForStringifiedFactorialResultat(s string) (int64, error) {

	var count int64 = 0

	fmt.Println(len(s) - 1)
	fmt.Println("Compute trailing zero count...")
	for i := len(s) - 1; i >= 0; i = i - 1 {
		if string(s[i]) == "0" {
			count += 1
		} else {
			break
		}
	}

	return count, nil
}
