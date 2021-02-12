package main

import "fmt"

func main() {
	var a, v0, s0 float64
	fmt.Println("Please enter acceleration, initial velocity, and initial displacement values(ex 10, 2, 1)")
	_, err := fmt.Scanf("%f, %f, %f", &a, &v0, &s0)

	fmt.Println("Please enter time value.")
	var time float64
	_, err2 := fmt.Scan(&time)

	if err != nil || err2 != nil {
		print("Problem scanning the values.")
	} else {
		//fmt.Printf(" a: %.2f\nv0: %.2f\ns0: %.2f\n", a, v0, s0)
		displacementFun := GenDisplaceFn(a, v0, s0)
		fmt.Printf("Displacement after %.2f seconds: %.2f\n", time, displacementFun(time))
	}
}

/*GenDisplaceFn Generator function that return another function.*/
func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	/*Returns a function that calculates the displacement where t is used for time.*/
	return func(t float64) float64 {
		return (((1.0 / 2) * (a * (t * t))) + (v0 * t) + s0)
	}
}
