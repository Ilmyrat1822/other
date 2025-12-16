package main

import (
	"bmicalc/info"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetUermetrics() (weight float64, height float64) {

	weight = readUserMetrics(info.Weighttitle)
	height = readUserMetrics(info.Heighttitle)

	return
}
func readUserMetrics(usermetricstxt string) (value float64) {
	fmt.Print(usermetricstxt)
	enteredvalue, _ := reader.ReadString('\n')
	enteredvalue = strings.TrimSpace(enteredvalue)
	value, _ = strconv.ParseFloat(enteredvalue, 64)
	return
}
