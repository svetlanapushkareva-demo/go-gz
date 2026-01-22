package main

import (
	"fmt"
)

func main() {
	fmt.Println("Конвертер валют")
	for {
		fromCurrency := getCurrency("Ведите исходную валюту (EUR, USD, RUB):")
		amount := getUserAmount("Ведите сумму для конвертации:")
		toCurrency := getCurrency("Ведите целевую валюту(EUR, USD, RUB):")
		convertCurrency(amount, fromCurrency, toCurrency)
		if !askYesNo("Выполнить еще одну конвертацию? (да/нет): ") {
			fmt.Println("До свидания!")
			break
		}
	}
}
func getCurrency(prompt string) string {
	fmt.Print(prompt)
	var currency string
	fmt.Scan(&currency)
	for {
		switch currency {
		case "EUR", "USD", "RUB":
			return currency
		default:
			fmt.Println("Ошибка! Доступные валюты: USD, EUR, RUB")
		}
	}
}
func getUserAmount(promt string) float64 {
	fmt.Print(promt)
	var amount float64
	for {
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println("Ошибка! Введите число")
			continue
		}
		if amount <= 0 {
			fmt.Println("Ошибка! Сумма должна быть больше нуля! ")
			continue
		}
		return amount
	}
}
func convertCurrency(amount float64, fromCurrency string, toCurrency string) {
	if fromCurrency == toCurrency {
		fmt.Println("Ошибка! Валюты одинаковые!")
		return
	}
	var result float64
	var rate float64
	switch fromCurrency + "_" + toCurrency {
	case "USD_EUR":
		rate = 0.86
	case "USD_RUB":
		rate = 77.75
	case "EUR_USD":
		rate = 1.16
	case "EUR_RUB":
		rate = 89.50
	case "RUB_USD":
		rate = 0.013
	case "RUB_EUR":
		rate = 0.011
	default:
		fmt.Println("Ошибка: неподдерживаемая конвертация")
		return
	}

	result = amount * rate

	fmt.Printf("Конвертация: %.2f %s → %s\n", amount, fromCurrency, toCurrency)
	fmt.Printf("Курс: 1 %s = %.2f %s\n", fromCurrency, rate, toCurrency)
	fmt.Printf("Результат: %.2f %s\n", result, toCurrency)

}
func askYesNo(prompt string) bool {
	fmt.Print(prompt)
	var answer string
	for {
		_, err := fmt.Scan(&answer)
		if err != nil {
			return false
		}
		switch answer {
		case "да":
			return true
		case "нет":
			return false
		default:
			fmt.Println("Пожалуйста, ответьте 'да' или 'нет'")
		}
	}
}
