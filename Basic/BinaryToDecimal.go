/*
Input : b = 111
Output : 7
Explanation : The decimal equivalent of the binary number 111 is 22 + 21 + 20 = 7.
*/
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	string := "100001"
	intVal, err := strconv.Atoi(string)
	if err != nil {
		fmt.Println("Error While Converting")
		return
	}
	fmt.Printf("Int value after conversion is %v", intVal)
	i := 0
	sum := 0
	for intVal != 0 {
		m := intVal % 10
		intVal = intVal / 10
		sum += m * int(math.Pow(float64(2), float64(i)))
		i++
	}
	fmt.Printf("\nConverting Binary to Decimal is %v", sum)
}
