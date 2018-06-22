package main

import (
	"math"
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number ", float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return f, ErrNegativeSqrt(f)
	}
	return math.Sqrt(f), nil
}


func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
