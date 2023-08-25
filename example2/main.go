package main

import "fmt"

func main() {
	fmt.Println("This function is initializing")

	example()

	fmt.Println("This function has ended succesfully")
}

func example() {
	// Aqui vamos a lanzar los panics forzosamente
	// panic("This is a panic")
	// panic(5)
	panic(true)
}
