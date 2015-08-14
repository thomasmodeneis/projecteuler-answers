/**
Multiples of 3 and 5
Problem 1
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
 */
package main

import (
	"fmt"
)

func main(){
	total := getMultiples(0,1000)
	fmt.Println(total)
}

func getMultiples(i, m int64) (t int64) {
	for ; i<m;i++ {
		if i % 3 == 0 || i % 5 == 0 {
			t = t + i
		}
	}
	return t
}

