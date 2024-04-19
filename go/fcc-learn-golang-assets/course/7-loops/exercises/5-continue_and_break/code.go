package main

import (
	"fmt"
	"math"
)

// printPrimes prints all prime numbers up to max
func printPrimes(max int) {
	for n := 2; n <= max; n++ {
		// Handle the base case for 2
		if n == 2 {
			fmt.Println(n)
			continue
		}

		// Skip even numbers
		if n%2 == 0 {
			continue
		}

		// Check for prime numbers
		isPrime := true
		for i := 3; i <= int(math.Sqrt(float64(n)))+1; i += 2 {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

		// Print if prime
		if isPrime {
			fmt.Println(n)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
