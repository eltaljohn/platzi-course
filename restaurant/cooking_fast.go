package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now()
	foods := []string{"Pizza 🍕", "Pasta 🍝", "Cake 🧁"}
	results := make(chan bool)
	for _, f := range foods {
		go func(f string) {
			cook(f)
			results <- true
		}(f)
	}
	for i := 0; i < len(foods); i++ {
		<-results
	}
	fmt.Printf("Done in %v\n", time.Since(started))
}

func cook(food string) {
	fmt.Printf("Cooking %s...\n", food)
	time.Sleep(2 * time.Second)
	fmt.Printf("Done cooking %s\n", food)
	fmt.Println("")
}
