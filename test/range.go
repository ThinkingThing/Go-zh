package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32}

func main() {

	slicetest := make([]int, 10)
	for i := range slicetest {
		slicetest[i] = 1 << uint(i)
	}
	for _, value := range slicetest {
		fmt.Println(value)
	}
}
func test() {

	for i, v := range pow {
		fmt.Println("i=%d,v=%d", i, v)
	}
}
