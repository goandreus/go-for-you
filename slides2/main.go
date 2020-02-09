package main

import "fmt"

func main() {
	food := [5]string{"🐈", "🐔", "🐄", "🦋", "🐕"}
	fruits := food[1:3]

	fmt.Println("food", food)
	fmt.Println("fruits", fruits)
	fmt.Println("capacidad", len(fruits))
	fmt.Println("capacidad", cap(fruits))
}
