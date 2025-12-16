package main

import "fmt"

type Product struct {
	Id    int
	Title string
	Price float64
}

func main() {
	//1
	var hobby = [3]string{"Football", "Basketball", "Voleyball"}
	fmt.Println(hobby)
	//2
	fmt.Println(hobby[1])
	fmt.Println(hobby[1:3])

	//3
	fmt.Println(hobby[0:2])
	hobby2 := hobby[0:2]
	fmt.Println(hobby2)

	//4
	hobby2 = hobby[1:3]
	fmt.Println(hobby2)

	//5
	goals := []string{"Professional footballer", "Go developer"}
	fmt.Println(goals)

	//6
	goals[1] = "Senior GO developer"
	goals = append(goals, "Became a billioner")
	fmt.Println(goals)

	//7
	products := []Product{
		{
			1,
			"Nike Predator",
			1000.0,
		},
		{
			2,
			"MacBook",
			1000.0,
		},
	}

	fmt.Println(products)

	NewProduct := Product{3, "Iphone16", 1000.0}

	products = append(products, NewProduct)

	fmt.Println(products)

	prices := []float64{23.23, 87.87}

	disprices := []float64{33.33, 45.45}

	prices = append(prices, disprices...)

	fmt.Println(prices)
}
