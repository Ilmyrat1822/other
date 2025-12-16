package main

import (
	"bmicalc/info"
	"fmt"
)

func main() {

	info.Start()

	we, he := GetUermetrics()

	fmt.Printf("Your bmi %.2f", calc(we, he))
}

func calc(we float64, he float64) (bmi float64) {
	bmi = we / (he * he)
	return
}
