package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now()
	foods := []string{"Pizza 🍕", "Pasta 🍝", "Cake 🧁"}
	for _, f := range foods {
		cook(f)
	}
	fmt.Printf("Done in %v\n", time.Since(started))
}

func cook(food string) {
	fmt.Printf("Cooking %s...\n", food)
	time.Sleep(2 * time.Second)
	fmt.Printf("Done cooking %s\n", food)
	fmt.Println("")
}
