package main

import "fmt"

func main() {
	fmt.Println("This function is initializing")

	example()

	fmt.Println("This function has ended succesfully")
}

func example() {
	animals := []string{
		"cow",  // indice 0
		"dog",  // indice 1
		"hawk", // indice 2
		"bear", // indice 3
	}

	fmt.Println("This animal can fly:", animals[5])
}
