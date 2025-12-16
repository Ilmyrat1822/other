package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	newnumbers := transformNumbers(&numbers, func(num int) int {
		return num * 2
	})

	fmt.Println(newnumbers)

}
func transformNumbers(numbers *[]int, transform transformFn) []int {
	resnumbers := []int{}

	for _, value := range *numbers {
		resnumbers = append(resnumbers, transform(value))
	}

	return resnumbers
}
