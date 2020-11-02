package main

import "fmt"

//Function with parameters to find the average of an array.
func avg(xs []float64) float64 {
	avg := 0.0
	for _, v := range xs {
		avg += v
	}
	return avg / float64(len(xs))
}

func main() {
	xs := []float64{98, 94, 93, 90, 80}
	fmt.Println(avg(xs))
}
