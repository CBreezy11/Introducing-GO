package main

import "fmt"



func minint () {
	x := []int {
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	var min int= 100
	for _, value := range x {
		if value < min {
			min = value
		} else {
			min = min
		}
	}
	fmt.Println(min) 

}
func main () {
	minint()
}