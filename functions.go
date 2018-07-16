package main

import "fmt"

func fibfind (x int) int {
	switch x {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibfind(x-1) + fibfind(x-2)
	}
	
}

func main () {
	fmt.Print(fibfind(8))
}