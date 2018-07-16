package main

import "fmt"

func main() {
	vals := []int{4, 2, 6}
	for i, v := range vals {
		vals[i] = v - 2
	}
	fmt.Println(vals)
}