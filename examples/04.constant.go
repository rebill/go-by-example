package main

import (
	"fmt"
	"math"
)

// const declares a constant value
const S string = "constant"

func main() {
	fmt.Println(S)

	// A const statement can appear anywhere a var statement can
	const N = 500000000

	// Constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / N
	fmt.Println(d)

	// A numeric constant has no type until it’s given one, such as by an explicit cast
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one,
	// such as a variable assignment or function call.
	// For example, here math.Sin expects a float64
	fmt.Println(math.Sin(N))
}
