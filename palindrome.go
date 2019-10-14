package main

import (
	"fmt"
	"strconv"
)

func main() {
	//n and k are inputs
	n := 0
	k := 2

	length := 0
	currentOver := 0
	for currentOver < n {
		length++
		currentOver += palindromeCount(length, k)
	}
	currentOver -= palindromeCount(length, k)

	fmt.Println(palindromeAt(length, k, n-currentOver-1))
}

func palindromeCount(length, k int) int {
	total := 1
	for i := 0; i < (length+1)/2; i++ {
		total *= k
	}
	return total
}

func palindromeAt(length, k, n int) string {
	result := strconv.FormatInt(int64(n), k)
	for len(result) < (length+1)/2 {
		result = "0" + result
	}

	//Mirror
	if length%2 == 0 {
		for i := 0; i < length/2; i++ {
			result += string(result[length/2-1-i])
		}
	}
	if length%2 == 1 {
		for i := 0; i < length/2; i++ {
			result += string(result[length/2-1-i])
		}
	}
	return result
}
