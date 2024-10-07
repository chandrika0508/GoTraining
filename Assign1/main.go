package main

import "fmt"

func prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		if prime(i) {
			sum += i
		}
	}
	fmt.Println("Sum of Prime numbers is", sum)
}
