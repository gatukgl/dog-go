package main

import (
	"fmt"

	"github.com/gatukgl/dog-go/fizzbuzz"
)

type person struct {
	name  string
	money float64
}

func (p person) say() string {
	return "Hey, " + p.name
}

func main() {
	fmt.Println("vim-go")
	var x string = "Helloworld"
	fmt.Println(x)
	y := 10
	fmt.Println(y)

	var number int = 1
	result := fizzbuzz.Fizzbuzz(number)
	fmt.Println(result)

	var gatuk person
	gatuk.name = "gatuk"
	gatuk.money = 2.2
	fmt.Println(gatuk)

	mac := person{
		name:  "mac",
		money: 3.0,
	}
	fmt.Println(mac)

	fmt.Println(gatuk.say())
	fmt.Println(mac.say())
}
