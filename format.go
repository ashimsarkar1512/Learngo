package main

import "fmt"

func main1 () {

	var name string = "ashim"
	var age int = 23
	rating := 4.4567

	// fmt.Printf("my name is %s age is %d and rating %.2f",name,age,rating)


	formatted :=fmt.Sprintf("my name is %s age is %d and rating %.2f",name,age,rating)

	fmt.Printf(formatted)
}