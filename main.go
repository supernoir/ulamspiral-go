package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ulam Spiral in Go v0.0.1")
	printPrimes(2045)
}

func checkPrime(num int) bool {
	var mod int = 0
	var isPrime bool = false

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			mod += 1
		}
	}
	if mod == 2 {
		isPrime = true
	}
	return isPrime

}

func printPrimes(max int) {
	for n := 1; n <= max; n++ {
		if checkPrime(n) == false {
			fmt.Printf("◉ ")
		} else if checkPrime(n) == true {
			fmt.Printf("◌ ")
		}
	}
}
