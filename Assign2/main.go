package main

import (
	"Assign2/calc"
	"fmt"
)

func main() {
	num1 := 39
	num2 := 23
	fmt.Println("Addition :", calc.Add(num1, num2))
	fmt.Println("Subtraction :", calc.Sub(num1, num2))
	fmt.Println("Multiplication :", calc.Multiplication(num1, num2))
	fmt.Println("Division :", calc.Division(num1, num2))
}
