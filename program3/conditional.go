package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// 1. IF-ELSE: Basic conditional
	age := rand.Intn(30) + 10 // Random age between 10-39
	if age < 18 {
		fmt.Printf("Age %d: You're a minor.\n", age)
	} else if age < 65 {
		fmt.Printf("Age %d: You're an adult.\n", age)
	} else {
		fmt.Printf("Age %d: You're a senior.\n", age)
	}

	// 2. FOR LOOP: Three styles
	fmt.Println("\nCountdown:")
	// Style 1: Traditional for loop
	for i := 5; i > 0; i-- {
		fmt.Printf("T-%d...\n", i)
	}

	// Style 2: While-like loop
	n := 0
	for n < 3 {
		fmt.Printf("Processing batch %d\n", n+1)
		n++
	}

	// Style 3: Infinite loop (with break)
	for {
		fmt.Println("This runs once (break after)")
		break
	}

	// 3. SWITCH: Multi-case branching
	day := time.Now().Weekday()
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("\nIt's the weekend! 🎉")
	default:
		fmt.Println("\nWeekday mode activated.")
	}

	// 4. DEFER: Execute later (LIFO order)
	defer fmt.Println("\nCleanup: Closing resources...")
	defer fmt.Println("Cleanup: Logging completion time")

	// 5. RANGE: Iterate over collections
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Println("\nFruits:")
	for idx, fruit := range fruits {
		fmt.Printf("%d. %s\n", idx+1, fruit)
	}

	// 6. CONTINUE/BREAK: Loop control
	fmt.Println("\nOdd numbers under 10:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skip evens
		}
		if i > 7 {
			break // Exit early
		}
		fmt.Println(i)
	}
}
