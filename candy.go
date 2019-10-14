package main

import (
	"fmt"
	"strconv"
)

func main() {
	shouts := []string{} //Input

	binaryString := ""
	for _, v := range shouts {
		if v == "even" {
			binaryString = "0" + binaryString
		}
		if v == "odd" {
			binaryString = "1" + binaryString
		}
	}
	num, _ := strconv.ParseInt(binaryString, 2, 64)
	fmt.Println(num + 1)
}
