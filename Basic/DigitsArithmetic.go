package main

import (
	"fmt"
)

func main() {
	n := 153
	a := 3
	b := 3
	k := 1
	ArmstrongNumber(n)
	SumOfDigits(n)
	ReverseDigits(n)
	KthDigit(a, b, k)

}

/*
An Armstrong number of three digits is a number such that the sum of the cubes of its digits is equal to the number itself. 371 is an Armstrong number since 33 + 73 + 13 = 371.
*/

func ArmstrongNumber(n int) {
	sum := 0
	original := n
	for n != 0 {
		m := n % 10
		n = n / 10
		sum += m * m * m
	}
	if original == sum {
		fmt.Printf("\n%v is an Armstrong Number", original)
	} else {
		fmt.Printf("\n%v is not an Armstrong Number", original)
	}
}

// Given a number n. Find the sum of all the digits of n.
func SumOfDigits(n int) {
	sum := 0

	for n != 0 {
		sum += n % 10
		n = n / 10
	}
	fmt.Printf("\nSum of digits is %v", sum)
}

/*
Input: n = 122
Output: 221
Explanation: By reversing the digits of number, number will change into 221.
*/
func ReverseDigits(n int) {
	reverse := 0
	for n != 0 {
		m := n % 10
		n = n / 10
		reverse = reverse*10 + m
	}
	fmt.Printf("\nReversed Digit is : %v", reverse)

}

/*
Input: a = 3, b = 3, k = 1
Output: 7
Explanation: 33 = 27 and 1st digit from right is 7
*/
func KthDigit(a, b, k int) {
	power := 1
	kthDigit := -1
	for i := 0; i < b; i++ {
		power *= a
	}
	//power := int(math.Pow(float64(a), float64(b)))
	fmt.Printf("\nValue of a = %v to the power of b = %v is %v", a, b, power)
	i := 1
	n := power
	for n != 0 {
		if i == k {
			kthDigit = n % 10
			break
		}
		i++
		n = n / 10
	}
	if kthDigit != -1 {
		fmt.Printf("\nk=%v kthDigit from right is %v", k, kthDigit)
	} else {
		fmt.Printf("\nInvalid value of k=%v", k)
	}
}
