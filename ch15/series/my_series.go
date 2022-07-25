package series

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

var LessThanTwoError = errors.New("n should not less than 2")
var LargerThanHundredError = errors.New("n should not larger than 100")

func GetFibonacciSerie(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHundredError
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func Square(n int) int {
	return n * n
}
