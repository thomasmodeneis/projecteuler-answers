//Smallest multiple
//Problem 5
//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//
//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
//
//solution with math version


package main

import (
	"fmt"
	"math"
)

func main() {
	divisorMax := 20
	p := generatePrimes(divisorMax)
	result := float64(1)

	for i := 0; i < len(p); i++ {
		a := math.Floor(math.Log(float64(divisorMax)) / math.Log(float64(p[i])));
		result = result*(math.Pow(float64(p[i]), a));
	}

	fmt.Println(int64(result))

}

func generatePrimes(upperLimit int) (primes []int) {
	isPrime := false
	j := 0
	primes = append(primes, 2)
	for i := 3; i <= upperLimit; i += 2 {
		j = 0;
		isPrime = true;
		for ; primes[j]*primes[j] <= i; {
			if i%primes[j] == 0 {
				isPrime = false;
				break;
			}
			j++;
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}
