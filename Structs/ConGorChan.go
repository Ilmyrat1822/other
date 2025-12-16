package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var rnd = rand.New(source)

func main() {

	var a = make(chan int)
	var b = make(chan int)
	var limit = make(chan int, 3)
	var x, y int
	go generateNumber(a, limit)
	go generateNumber(b, limit)

	select {
	case x = <-a:
		fmt.Printf("Value from a channel= %v", x)
	case y = <-b:
		fmt.Printf("Value from b channel= %v", y)
	}

	/*		sum := 0
			go generateNumber(c, limit)
			go generateNumber(c, limit)
				i := 0
				for num := range c {
					i++
					sum += num
					if i == 4 {
						close(c)
					}
				}
				fmt.Println(sum)*/
}

func generateNumber(c chan int, limit chan int) {
	limit <- 1
	fmt.Println("Generated")
	sleeptime := rnd.Intn(3)
	time.Sleep(time.Duration(sleeptime) * time.Second)
	result := rnd.Intn(10)

	c <- result
	<-limit
}
