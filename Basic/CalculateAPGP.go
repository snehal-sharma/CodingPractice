
package main

import "fmt"

func main() {
	first := 2
	second := 3
	ratio := 3
	n := 5

	ArithmeticProgression(first, second, n)
	GeometricProgression(first, ratio, n)

}

func ArithmeticProgression(first, second, n int) {
	fmt.Println("Printing all values for ArithmeticProgression")
	fmt.Println(first)
	fmt.Println(second)
	for i := 0; i < n; i++ {
		first, second = second, first+second
		fmt.Println(second)
	}
}

func GeometricProgression(first, ratio, n int) {
	fmt.Println("Printing all values for GeometricProgression")
	fmt.Println(first)
	result := first
	for i := 0; i < n; i++ {
		result *= ratio
		fmt.Println(result)
	}
}
