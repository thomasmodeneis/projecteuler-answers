/**
Largest palindrome product
Problem 4
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
 */

package main

import (
	"fmt"
)

func main(){
	h := 0
	for i := 999; i>99; i-- {
		for j := 100; j < 1000; j++ {
			p := i * j
			r := reverse(p)
			if isPalindrome(r, p) {
				if h < p {
					h = p
				}
			}
		}
	}

	fmt.Println(h)
}

func isPalindrome(r,v int) bool{
	return r == v
}

func reverse(n int) (r int) {
	for {
		if n > 0 {
			r = r * 10 + n % 10
			n = n / 10
		}else {
			break
		}
	}
	return r
}
