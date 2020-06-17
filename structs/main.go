package main

import "fmt"

func main() {
	type course struct {
		Name      string
		Professor string
		Country   string
	}

	db := course{
		Name:      "Bd",
		Professor: "andres",
		Country:   "peru",
	}

	/* 	git := course{"git", "carlos", "mexico"}
	   	css := course{Professor: "cesar"} */

	p := &db
	p.Professor = "Cesar"

	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", p)
}
