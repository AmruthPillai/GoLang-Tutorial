package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

func concat(a, b string) (string, string) {
	return a, b
}

func main() {
	// var num1 float64 = 5.6
	// var num2 float64 = 9.5
	// var num1, num2 float64 = 5.6, 9.5

	// Defaults to float64
	num1, num2 := 5.6, 9.5
	fmt.Println(add(num1, num2))

	w1, w2 := "Hey,", "there!"
	fmt.Println(concat(w1, w2))

	// Type Conversion
	var a = 62
	var b = float64(a)
	fmt.Println("Value of B:", b)

	// Type Inference, Value Copy
	x := a
	fmt.Println("Value of x:", x)
}
