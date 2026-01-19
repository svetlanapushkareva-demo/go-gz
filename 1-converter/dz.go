package main

import "fmt"

func main() {
	const USD2EUR = 1.16
	const USD2RUB = 77.75
	const EUR2RUB = USD2EUR * USD2RUB / 100
	var EURMoney float64 = 100

	fmt.Println(EURMoney, "изначальные евро")

	USDMoney := EURMoney * USD2EUR

	fmt.Println(USDMoney, "доллары")

	RUBMoney := USDMoney * USD2RUB

	fmt.Println(RUBMoney, "рубли")
	fmt.Println(EUR2RUB, "евро в рублях")
}

func userScan(text string) string {
	value := ""
	fmt.Println(text)
	fmt.Scan(&value)

	return value
}
func convertCurrency(amount float64, fromCurrency string, toCurrency string) {

}
