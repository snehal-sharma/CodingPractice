/*
The factorial of a number is the product of all positive integers less than or equal to that number.
For example, 7! means 1 × 2 × 3 × 4 × 5 × 6 × 7. 

nPr  represents n permutation r and value of nPr  is (n!) / (n-r)!.

The binomial coefficient nCr is calculated as : C(n,r) = n! / r! * (n-r)!.

*/
package main

import "fmt"

func main() {
	n := 5
	r := 2
	factorial := fact(5)
	if n >= r {
		nPrVal := ComputeNPR(n, r)
		nCrVal := ComputeNCR(n, r)
		fmt.Printf("\nFactorial of n=%v is %v", n, factorial)
		fmt.Printf("\nCompute nPr of n=%v and r=%v is %v", n, r, nPrVal)
		fmt.Printf("\nCompute nPr of n=%v and r=%v is %v", n, r, nCrVal)
	} else {
		fmt.Printf("\nCompute nPr of n=%v and r=%v is %v", n, r, 0)
		fmt.Printf("\nCompute nPr of n=%v and r=%v is %v", n, r, 0)

	}
}

func fact(n int) int {
	fact := 1
	for i := n; i >= 2; i-- {
		fact *= i
	}
	return fact
}

func ComputeNPR(n, r int) int {
	return fact(n) / fact(n-r)
}

func ComputeNCR(n, r int) int {
	return fact(n) / (fact(r) * fact(n-r))
}
