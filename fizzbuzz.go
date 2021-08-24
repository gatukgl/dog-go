package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "FizzBuzz"
	} else if number%3 == 0 {
		return "Fizz"
	} else if number%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(number)
	}
}

func main() {
	fmt.Println("vim-go")
	var x string = "Helloworld"
	fmt.Println(x)
	y := 10
	fmt.Println(y)

	var number int = 1
	result := fizzbuzz(number)
	fmt.Println(result)
}
