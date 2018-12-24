package main

import "fmt"

func main() {
	x := 15
	// var a *int = &x

	a := &x // memory address

	fmt.Println(a)
	fmt.Println(*a)

	*a = 5
	fmt.Println(x)

	*a = *a * *a
	fmt.Println(x)
}
