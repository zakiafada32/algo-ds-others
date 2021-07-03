package main

import "fmt"

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	adidasShoe := adidasFactory.makeShoe()
	adidasShort := adidasFactory.makeShort()

	nikeFactory, _ := getSportsFactory("nike")
	nikeShoe := nikeFactory.makeShoe()
	nikeShort := nikeFactory.makeShort()

	printShoeDetails(adidasShoe)
	printShortDetails(adidasShort)

	printShoeDetails(nikeShoe)
	printShortDetails(nikeShort)

}

func printShoeDetails(s iShoe) {
	fmt.Printf("logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShortDetails(s iShort) {
	fmt.Printf("logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
