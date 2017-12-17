package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
	test()
}

func test() {

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
