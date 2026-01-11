package main

import "fmt"

func main() {
	n := 50
	arr := make([]int, n+1)
	var primes []int
	arr[0] = 1
	arr[1] = 1

	for i := 2; i <= n; i++ {
		if arr[i] == 0 {
			primes = append(primes, i)
			for j := i + i; j <= n; j = j + i {
				arr[j] = 1
			}
		}
	}

	fmt.Printf("List of Prime numbers under %d : %v", n, primes)
}
