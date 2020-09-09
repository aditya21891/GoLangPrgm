package main

import "fmt"

func main() {
	var temp float64
	var cel float64

	fmt.Println("Please Enter Temp in F")
	fmt.Scanf("%f", &temp)
	cel = (temp - 32) * 5 / 9
	fmt.Println("Temperature in celsius: ", cel)

}
