package main

import (
	"fmt"
	"math"
)

func main() {
	var t float64
	a, v0, s0 := getParams()
	fn := GenDisplaceFn(a, v0, s0)

	fmt.Println("Enter value for time (t):")
	fmt.Scan(&t)
	fmt.Printf("The displacent is %f \n", fn(t))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}

func getParams() (float64, float64, float64) {
	var a, v0, s0 float64
	fmt.Println("Enter value for acceleration (a):")
	fmt.Scan(&a)
	fmt.Println("Enter value for initial velocity (v0):")
	fmt.Scan(&v0)
	fmt.Println("Enter value for initial displacement (s0):")
	fmt.Scan(&s0)
	return a, v0, s0
}
