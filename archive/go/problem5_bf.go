//Smallest multiple
//Problem 5
//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//
//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
//
//solution with brute force version



package main



import (
	"fmt"
)

func main(){

	counter := 0
	n := 0

	for j :=2520; counter < 20; j++ {

		for i := 1; i < 21; i++ {

			if j%i == 0 {
				n = j
				counter = counter + 1
			}else {
				counter = 0
			}
		}
	}

	fmt.Println(counter)
	fmt.Println(n)
}
