package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var id_counter int
var reader = bufio.NewReader(os.Stdin)

type Product struct {
	Name     string
	Quantity int
	Price    float64
}

func addProduct(products map[int]Product) int {
	var (
		name     string
		quantity int
		price    float64
	)

	fmt.Print("Enter product name: ")
	inputName, _ := reader.ReadString('\n')
	name = strings.TrimSpace(inputName)

	fmt.Print("Enter product quantity: ")
	fmt.Scanln(&quantity)

	fmt.Print("Enter product price: ")
	fmt.Scanln(&price)

	products[id_counter] = Product{
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}

	id_counter++

	return id_counter
}

func showProducts(products map[int]Product) {
	fmt.Println("List of products")
	for id, product := range products {
		fmt.Println()
		fmt.Println("Id", id)
		fmt.Println("Name", product.Name)
		fmt.Println("Quantity", product.Quantity)
		fmt.Println("Price", product.Price)
	}
}

func updateQuantity(products map[int]Product) {
	var id int
	fmt.Print("Enter product id: ")
	fmt.Scanln(&id)

	product, ok := products[id]
	if !ok {
		fmt.Println("Product not found")
		return
	}

	var newQuantity int
	fmt.Print("Enter new quantity: ")
	fmt.Scanln(&newQuantity)

	product.Quantity = newQuantity
	products[id] = product
}

func main() {
	products := make(map[int]Product)
	exit := false
	var choise int

	for !exit {
		fmt.Println()
		fmt.Println("1 - Add product")
		fmt.Println("2 - Show products")
		fmt.Println("3 - Update product quantity")
		fmt.Println("4 - Exit")
		fmt.Print(">> ")
		fmt.Scanln(&choise)
		fmt.Println()

		switch choise {
		case 1:
			addProduct(products)
		case 2:
			showProducts(products)
		case 3:
			updateQuantity(products)
		case 4:
			exit = true
		default:
			fmt.Println("Incorrect input")
		}
	}
}
