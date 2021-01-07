package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		var x float32
		fmt.Println("Enter a floating point number.")
		fmt.Scan(&x)

		var newX = int32(x)
		fmt.Printf("%d\n", newX)
	}

}
