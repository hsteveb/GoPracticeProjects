package main

import "fmt"
import "time"

func main() {
	var i = 0
	fmt.Println("Hello World!")

	go incrementValue("1", &i)
	go incrementValue("2", &i)

	time.Sleep(1000000)

}

func incrementValue(name string, value *int) {
	*value++
	fmt.Printf("thread %s increment value to %d\n", name, *value)
}
