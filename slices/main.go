package main

import "fmt"

func main() {
	//slice no datos, apuntan a un array
	set := [7]string{"🐈", "🐔", "🐄", "🦋", "🐕", "🦜", "🌵"}
	animals := set[0:5]
	fly := set[3:7]
	fly[0] = "🦅"
	fmt.Println("Array:", set)
	fmt.Println("Animals:", animals)
	fmt.Println("Flys:", fly)
	fmt.Println("all", set[:])
}
