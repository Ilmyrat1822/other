package random

import (
	"fmt"
	"math/rand"
)

func Calculate() {
	fmt.Println("Random adding numbers")
	a, b := Random()
	sum := Add(a, b)
	println(a, " ", b, " ", sum)

}
func Random() (int, int) {
	rnd1 := rand.Intn(10)
	rnd2 := rand.Intn(10)
	return rnd1, rnd2
}

func Add(a int, b int) int {
	return a + b
}
