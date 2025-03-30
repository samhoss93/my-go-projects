package main

import "fmt"

func foobar() {
	// Immediately Invoked Function Expression (IIFE)
	func(msg string) {
		fmt.Println(msg) // Output: Hello
	}("Hello")

	// Assign to a variable
	greet := func(name string) string {
		return "Hi, " + name
	}
	fmt.Println(greet("Alice")) // Output: Hi, Alice
}

func main() {
	foobar()
}
