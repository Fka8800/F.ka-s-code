package main

import "fmt"

func main() {
	a := 0.0
	b := 0.0
	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)
	fmt.Println(a, "+", b, "=", a+b)
}
