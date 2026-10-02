package main

import "fmt"

func makingCoffee(kind string, isSuger bool) {

	fmt.Printf("making %s coffee", kind)
	fmt.Printf("sugar added %t", isSuger)

}

func main() {

	makingCoffee("black", true)

}
