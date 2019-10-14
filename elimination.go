package main

import "fmt"

func main() {
	//Input is not stored here as it is too big.
	//Please store in elimination_input.go.

	numberOfZeroes := 0
	arrayOfGaps := make([]int, 0)
	for i := 0; i < len(input); i++ {
		if input[i] == '0' {
			numberOfZeroes++
		} else {
			arrayOfGaps = append(arrayOfGaps, numberOfZeroes)
			numberOfZeroes = 0
		}
	}
	//First element is useless
	arrayOfGaps = arrayOfGaps[1:]
	max := 0
	for i := 0; i < len(arrayOfGaps); i++ {
		j := i - 1
		totalEncountered := 0
		for totalEncountered <= k {
			j++
			if j >= len(arrayOfGaps) {
				break
			}
			totalEncountered += arrayOfGaps[j]
		}
		if j-i+1 > max {
			max = j - i + 1
		}
	}
	fmt.Println(max)
}
