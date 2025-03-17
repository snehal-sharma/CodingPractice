/*
n=3
3	3	3	2	2	2	1	1	1	
3	3	2	2	1	1	
3	2	1	
*/
package main

import "fmt"

func PrintPattern(n int) {
	for i := n; i > 0; i-- {
		for j := n; j > 0; j-- {
			for k := i; k > 0; k-- {
				fmt.Printf("%v\t", j)
			}
		}
		fmt.Printf("\n")
	}
}
func main() {
	PrintPattern(3)
}
