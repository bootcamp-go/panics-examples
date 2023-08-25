package main

import "fmt"

type Cat struct {
	name string
}

func main() {
	fmt.Println("This function is initializing")

	example()

	fmt.Println("This function has ended succesfully")
}

func example() {
	s := &Cat{"Yummy"}
	s = nil

	s.Speak()
}

func (s *Cat) Speak() {
	fmt.Println(s.name, "The cat goes 'meow'")
}
