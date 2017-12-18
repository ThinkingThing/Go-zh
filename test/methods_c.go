package main

import (
	"fmt"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	//f1 := MyFloat(-math.Sqrt2)
	f1 := MyFloat(12.12)
	fmt.Println(f1.Abs())
}
