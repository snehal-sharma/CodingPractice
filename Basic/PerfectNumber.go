/*
Given a number n, check if a number is perfect or not. A number is said to be perfect if sum of all its factors excluding the number itself is equal to the number. 
Input: n = 6
Output: true 
Explanation: Factors of 6 are 1, 2, 3 and 6. Excluding 6 their sum is 6 which is equal to N itself. So, it's a Perfect Number.
*/

package main

import "fmt"

func main() {
	n := 10
	sum := 0
	arr := factors(n)
	fmt.Println(arr)
	if len(arr) == 2 {
		fmt.Println("Prime numbers are not perfect numbers")
	}

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	if sum == n {
		fmt.Printf("%v is a Perfect Number", n)
	} else {
		fmt.Printf("%v is not a Perfect Number", n)
	}

}

func factors(n int) []int {
	var factorsArray []int
	for i := n - 1; i >= 1; i-- {
		if n%i == 0 {
			factorsArray = append(factorsArray, i)
		}
	}
	if len(factorsArray) == 1 {
		factorsArray = append(factorsArray, n)
	}
	return factorsArray
}
