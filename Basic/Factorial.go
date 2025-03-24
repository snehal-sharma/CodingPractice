
package main

import "fmt"

func main() {
	n := 5
	r := 2
	factorial := fact(5)
	nPrVal := ComputeNPR(n, r)
	nCrVal := ComputeNCR(n, r)
	fmt.Printf("\nFactorial of n=%v is %v", n, factorial)
	fmt.Printf("\nCompute nPr of n=%v and r=%v is %v", n, r, nPrVal)
	fmt.Printf("\nCompute nPr of n=%v and r=%v is %v", n, r, nCrVal)
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
