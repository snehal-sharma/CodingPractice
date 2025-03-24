package main

import "fmt"

func main() {
	a := 14
	b := 8
	gcdVal := FindGCD(a, b)
	lcmVal := (a * b) / gcdVal
	fmt.Printf("\nGCD of a=%v and b=%v is %v", a, b, gcdVal)
	fmt.Printf("\nLCM of a=%v and b=%v is %v", a, b, lcmVal)
}

func FindGCD(a, b int) int {
	min := 0
	if a < b {
		min = a
	} else {
		min = b
	}
	for i := min; i > 1; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}
