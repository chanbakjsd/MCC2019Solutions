package main

import "fmt"

func main() {
	//Expects variable "prices" to hold it. Didn't store here cause test cases are like 100KB.
	//See ditcoin_input.go for information on where to store it.

	currentHighest := 0
	total := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > currentHighest {
			currentHighest = prices[i]
		}
		total += currentHighest
	}
	fmt.Println(total)
}
