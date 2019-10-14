package main

import (
	"fmt"
	"sort"
)

func main() {
	prices := []int{} //Input

	sort.Ints(prices)
	totalPrice := 0
	for _, price := range prices {
		totalPrice += price
	}
	for i := 0; i < len(prices)/4; i++ {
		discount := prices[len(prices)-3-i*3]
		totalPrice -= discount
	}
	fmt.Println(totalPrice)
}
