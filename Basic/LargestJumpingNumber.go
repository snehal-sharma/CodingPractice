/*
Given a positive number, n. Find the largest Jumping Number with equal number of digits that is smaller than or equal to x. 
Jumping Number: A number is called a Jumping Number if all adjacent digits in it differ by only 1. All single-digit numbers are considered as Jumping Numbers.
For example, 7, 898, 7 and 4343456 are Jumping numbers but 796, 677, and 89098 are not.
Input: n = 50
Output: 45
Explanation: 45 is the largest Jumping Number possible for X = 50.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	n := 100
	result := -1
	for i := n; i > 0; i-- {
		result = JumpingNumber(i)
		if result != -1 {
			fmt.Printf("\n%v is the largest Jumping Number", i)
			break
		}
	}

}

func JumpingNumber(n int) int {
	prev := -1
	original := n
	for n != 0 {
		m := n % 10
		if prev != -1 && int(math.Abs(float64(m)-float64(prev))) != 1 {
			fmt.Printf("\n%v is not a Jumping Number", original)
			return -1
		}
		prev = m
		n = n / 10
	}
	fmt.Printf("\n%v is a Jumping Number", original)
	return original
}
