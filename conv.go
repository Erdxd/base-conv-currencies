package main

import (
	"fmt"
)

type mmap = map[string]float64

var currencyyouhave string

var quantity float64
var currencyyouwant string

func main() {

	for {
		fmt.Println("Выберите вариант")
		fmt.Println("1-конвертор валют")
		fmt.Println("4-выход")
		user := ""
		fmt.Scan(&user)

		if user == "1" {
			fact()
		} else if user == "4" {
			break

		}
	}

}
func fact() {
menu1:
	for {

		fmt.Println("Конвертор валют(доллар, евро, рубль) хотите начать? да/нет")
		user := ""
		fmt.Scan(&user)
		if user == "да" {
			fmt.Println("какая у вас валюта?(usd,euro,rub)")
			fmt.Scan(&currencyyouhave)
			if currencyyouhave == "usd" {
				usd()
			} else if currencyyouhave == "eur" {
				eur()
			} else if currencyyouhave == "rub" {
				rub()

			} else {
				fmt.Println("Ошибка, попробуйте снова")
				break menu1
			}

		} else if user == "нет" {
			break
		} else {
			fmt.Println("Ошибка, начните заново")
		}

	}

}
func usd() {

	valuteusd := map[string]float64{"eur": 0.89, "rub": 80.50}
	if currencyyouhave == "usd" { //для доллара
		fmt.Println("current rate:")
		for index, value := range valuteusd {
			fmt.Println(index, value)
		}
		fmt.Println("сколько у вас", currencyyouhave, "?")
		fmt.Scan(&quantity)
		quantity1(quantity)
		fmt.Println("В какую валюту вы хотите конвертировать?")
		fmt.Scan(&currencyyouwant)
		currencyyouwant1(currencyyouwant)
		if currencyyouwant == "eur" {
			result := quantity * valuteusd["eur"]
			fmt.Println("у вас есть", result, currencyyouhave)
		} else if currencyyouwant == "rub" {
			result := quantity * valuteusd["rub"]
			fmt.Println("у вас есть", result, currencyyouhave)
		} else if currencyyouwant == "usd" {
			fmt.Println("pointless")
		}

	} // конец доллара

}
func eur() {

	valuteeur := map[string]float64{"usd": 1.12, "rub": 90.23}
	if currencyyouhave == "eur" { //для евро

		fmt.Println("current rate:")
		for index, value := range valuteeur {
			fmt.Println(index, value)
		}
		fmt.Println("сколько у вас", currencyyouhave, "?")
		fmt.Scan(&quantity)
		quantity1(quantity)
		fmt.Println("В какую валюту вы хотите конвертировать?")
		fmt.Scan(&currencyyouwant)
		currencyyouwant1(currencyyouwant)
		if currencyyouwant == "eur" {

			fmt.Println("pointless")
		} else if currencyyouwant == "rub" {
			result := quantity * valuteeur["rub"]
			fmt.Println("у вас есть", result, currencyyouhave)
		} else if currencyyouwant == "usd" {
			result := quantity * valuteeur["usd"]
			fmt.Println("у вас есть", result, currencyyouhave)
		}

	}
}

func rub() {

	valuterub := map[string]float64{"euro": 0.011, "usd": 0.012}
	if currencyyouhave == "rub" { //для рубля
		fmt.Println("current rate:")
		for index, value := range valuterub {
			fmt.Println(index, value)
		}

		fmt.Println("сколько у вас", currencyyouhave, "?")
		fmt.Scan(&quantity)
		quantity1(quantity)
		fmt.Println("В какую валюту вы хотите конвертировать?")
		fmt.Scan(&currencyyouwant)
		currencyyouwant1(currencyyouwant)
		if currencyyouwant == "eur" {
			result := quantity * valuterub["eur"]
			fmt.Println("у вас есть", result, currencyyouhave)

		} else if currencyyouwant == "rub" {
			fmt.Println("pointless")

		} else if currencyyouwant == "usd" {
			result := quantity * valuterub["usd"]
			fmt.Println("у вас есть", result, currencyyouhave)
		}

	} // конец рубля

}

func quantity1(quantity float64) float64 {
	for {

		if quantity > 0.0 {
			return quantity

		} else {
			fmt.Println("Ошибка, попробуйте еще раз")
			break
		}
	}
	return quantity
}
func currencyyouwant1(currencyyouwant string) string {
	for {

		if currencyyouwant == "eur" || currencyyouwant == "rub" || currencyyouwant == "usd" {

			return currencyyouwant

		} else {
			fmt.Println("Ошибка, попробуйте еще раз")
			break
		}
	}
	return currencyyouwant
}
