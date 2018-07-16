package main

import "fmt"

var feet float64= 100

func main () {
	fmt.Print("Enter the length in feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)

	var meters float64= (feet * 0.3048)
	fmt.Println("The length in meters is: ", meters)
}