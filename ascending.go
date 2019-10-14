package main

import "fmt"

func main() {
	numbers := []int{} //Input
	lowest := 99999999
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i+1]-numbers[i] < lowest {
			lowest = numbers[i+1] - numbers[i]
		}
	}
	fmt.Println(lowest)
}
