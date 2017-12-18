package main

import (
	"fmt"
)

type vv struct {
	X, y int
}

func (v *vv) abs() int {
	return v.X + v.y
}

func main() {
	v := &vv{1, 2}
	fmt.Println(v.abs())
}
