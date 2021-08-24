package main

import (
	"fmt"
	"io"
	"net/http"

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

	const url string = "https://dog.ceo/api/breeds/image/random"
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		fmt.Println("Get error")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error")
	}
	fmt.Printf("%s", body)
}
