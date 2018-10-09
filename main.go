package main

import "fmt"

func fib() int {
	x, y := 0, 1
	var z int
	for i := 0; i < 4000000; i = x {
		x, y = y, x+y
		if x%2 == 0 {
			//fmt.Println(x)
			z += x
		}
	}
	return z
}

func main() {
	fibsum := fib()
	fmt.Println("The sum of all even Fibonacci numbers whose values do not exceed four million are: ", fibsum)
}

/*By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.*/
