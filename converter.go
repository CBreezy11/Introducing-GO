package main

import "fmt"


func main () {
	fmt.Print("Enter the temp in Farenheit: ")
	var fdegree float64
	fmt.Scanf("%f", &fdegree)

	celcius := ((fdegree -32) *5/9)
	fmt.Println("The temperature in Celcius is: ", celcius)
}