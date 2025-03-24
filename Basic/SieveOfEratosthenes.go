
package main

import "fmt"

func main() {
	n := 35
	arr := make([]int, n+1)
	var primes []int
	arr[0] = 1
	arr[1] = 1
	for i := 2; i <= n; i++ {
		if arr[i] == 0 {
			for j := i + i; j <= n; j = j + i {
				arr[j] = 1
			}
		}
	}

	for i := 2; i <= n; i++ {
		if arr[i] == 0 {
			primes = append(primes, i)
		}
	}

	fmt.Println(primes)
}
