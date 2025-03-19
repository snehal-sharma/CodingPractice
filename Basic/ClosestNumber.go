/*
Given two integers n and m. The problem is to find the number closest to n and divisible by m. If there is more than one such number, then output the one having the maximum absolute value.
Input: n = 13 , m = 4
Output: 12
Explanation: 12 is the Closest Number to 13 which is divisible by 4.
Input: n = -15 , m = 6
Output: -18
Explanation: -12 and -18 are both similarly close to -15 and divisible by 6. but -18 has the maximum absolute value. So, Output is -18
*/


package main

import (
	"fmt"
	"math"
)

func main() {
	n := -14
	m := 6
	result := 0
	upperLimit := int(math.Abs(float64(n)) + math.Abs(float64(m)))
	lowerLimit := int(math.Abs(float64(n)) - math.Abs(float64(m)))
	difference := upperLimit - lowerLimit
	for i := lowerLimit; i < upperLimit; i++ {
		currentDifference := int(math.Abs((math.Abs(float64(n)) - math.Abs(float64(i)))))
		if i%m == 0 && currentDifference <= difference {
			result = i
			difference = int(math.Abs((math.Abs(float64(n)) - math.Abs(float64(i)))))
		}
	}
	fmt.Println(result)
}
