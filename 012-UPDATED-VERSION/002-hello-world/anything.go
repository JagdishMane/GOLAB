package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is the best class..................................................................................................................")
	foo()
	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("I'm in bar and then exited control flow.")
}

// control flows:
// 1. Sequence
//2. loop; iterative
//3.  conditional
