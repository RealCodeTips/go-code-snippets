package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) string {
	var result string
	for i := 1; i <= n; i++ {
		multipleOfThree := i%3 == 0
		multipleOfFive := i%5 == 0

		if multipleOfThree && multipleOfFive {
			result += "fizzbuzz"
		} else if multipleOfThree {
			result += "fizz"
		} else if multipleOfFive {
			result += "buzz"
		} else {
			result += strconv.Itoa(i)
		}
	}
	return result
}

func main() {
	fmt.Println(fizzBuzz(4))  // 12fizz4
	fmt.Println(fizzBuzz(15)) // 12fizz4buzzfizz78fizzbuzz11fizz1314fizzbuzz
	fmt.Println(fizzBuzz(1))  // 1
}
