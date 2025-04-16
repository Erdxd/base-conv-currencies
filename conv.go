package main

import "fmt"

const a float64 = 83.64
const b float64 = 94.55

func main() {

	var d int = 0
	var c string = ""
	fmt.Println("Please write your money in rub")

	fmt.Scan(&d)
	fmt.Println("Please write the currency to convert")
	fmt.Scan(&c)
	if c == "usd" || c == "USD" || c == "Usd" {
		z := float64(d) / a
		fmt.Println(z)

	} else if c == "eur" || c == "EUR" || c == "Eur" {
		v := float64(b) / a
		fmt.Println(v)

	}

}
