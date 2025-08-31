// A goroutine is a lightweight thread of execution.

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ": ", i)
	}
}

func main() {

	f("synchronous")

	// to invoke a function in a go routine
	// use go f()

	go f("goroutine")

	// we can also use an anonymous function
	go func(mesage string) {
		fmt.Println(mesage)
	}("going")

	// two function calls are running asynchronously in separate goroutines
	// wait for them to finish

	time.Sleep(time.Second)
	fmt.Println("done")
}
