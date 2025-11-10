package main

import (
	"fmt"

	"hello_world/ratings"
)

func main() {
	r := ratings.NewRating(1)

	e := r.AddRating(1, 5, "Excellent product!")
	if e != nil {
		fmt.Println(e)
	}
	e = r.AddRating(2, 6, "Good value for money")
	if e != nil {
		fmt.Println(e)
	}
	e = r.AddRating(3, 3.5, "Okay but could be better")
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(r)
}
