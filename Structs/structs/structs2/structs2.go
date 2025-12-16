package structs2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Product struct {
	Id          string
	Title       string
	Description string
	Price       float64
}

func getProduct() *Product {
	fmt.Println("Product Details")
	fmt.Println("---------------")

	reader := bufio.NewReader(os.Stdin)
	idInput := ReaduserInput(reader, "Enter Product Id: ")
	titleInput := ReaduserInput(reader, "Enter Product Title: ")
	descInput := ReaduserInput(reader, "Enter Product Description: ")
	priceInput := ReaduserInputFloat64(reader, "Enter Product Price: ")

	product := NewProducts(idInput, titleInput, descInput, priceInput)
	return product

}

func ReaduserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	return userInput
}

func ReaduserInputFloat64(reader *bufio.Reader, promptText string) float64 {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	userval, _ := strconv.ParseFloat(userInput, 64)
	return userval
}

func NewProducts(i string, t string, d string, p float64) *Product {
	return &Product{i, t, d, p}
}

func (pro *Product) PrintData() {
	fmt.Printf("Id: %s\n", pro.Id)
	fmt.Printf("Title: %s\n", pro.Title)
	fmt.Printf("Description: %s\n", pro.Description)
	fmt.Printf("Price: %0.2f\n", pro.Price)
}

func (pro *Product) store() {
	reg := regexp.MustCompile(`[<>:"/\\|?*]`)
	safeId := reg.ReplaceAllString(pro.Id, "_")
	file, err := os.Create(safeId)
	if err != nil {
		fmt.Println(err.Error())
	}
	content := fmt.Sprintf("Id: %s\nTitle: %s\nDescription: %s\nPrice: %0.2f\n", pro.Id, pro.Title, pro.Description, pro.Price)

	file.WriteString(content)
	file.Close()
}

func main() {

	createProduct := getProduct()
	createProduct.PrintData()
	createProduct.store()
}
