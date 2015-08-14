/**

 */
package main

import (
	"fmt"
)

func main(){
	sum := 0;
	fib := []int{2,0}
	i := 0

	for fib[i] < 4000000 {
		sum += fib[i];
		i = (i + 1) % 2;
		fib[i] = 4 * fib[(i + 1) % 2] + fib[i];
	}

	fmt.Println(sum)
}

func fib(v1,v2 int) (int,int) {
	return v1+v2,v1
}

