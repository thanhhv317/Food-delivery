package main

import "fmt"

func main() {
	var moneyNow float64 = 19900

	var annualRateOfReturn float64 = 16.6

	var annualGenieMoney float64 = 0

	var yourAge int = 30

	for i := 0; i <= 40; i++ {
		yourAge++
		moneyNow += moneyNow*annualRateOfReturn/100 + annualGenieMoney
		fmt.Println(moneyNow)
		if moneyNow >= float64(1000000) {
			break
		}

	}

	fmt.Println(yourAge)
}
