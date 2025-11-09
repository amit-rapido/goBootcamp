package main

import "fmt"

func main() {

	type Rating struct {
		productName string
		rating      int
		comment     []string
		user        string
	}

	// map1 := map[string]int{}

	// map1["apple"] = 5
	// map1["banana"] = 4
	// map1["cherry"] = 3
	// map1["date"] = 2
	// map1["elderberry"] = 1

	// fmt.Println(map1)

	// switch rating {
	// case 1:
	// 	fmt.Println("*")
	// case 2:
	// 	fmt.Println("**")
	// case 3:
	// 	fmt.Println("***")
	// case 4:
	// 	fmt.Println("****")
	// case 5:
	// 	fmt.Println("*****")
	// }

	var r Rating

	fmt.Print("Enter product name: ")
	fmt.Scanln(&r.productName)

	fmt.Print("Enter your name: ")
	fmt.Scanln(&r.user)

	fmt.Print("Enter rating (1–5): ")
	fmt.Scanln(&r.rating)

	for r.rating < 1 || r.rating > 5 {
		fmt.Print("Invalid rating. Enter again (1–5): ")
		fmt.Scanln(&r.rating)
	}

	// Keep reading comments until user types "done"
	fmt.Println("Enter comments (type 'done' to finish):")
	for {
		var input string
		fmt.Scanln(&input)
		if input == "done" {
			break
		}
		r.comment = append(r.comment, input)
	}

	// Print results
	fmt.Println("\n--- Rating Summary ---")
	fmt.Println("Product:", r.productName)
	fmt.Println("User:", r.user)
	fmt.Println("Rating:", r.rating)
	fmt.Println("Comments:")
	for _, c := range r.comment {
		fmt.Println("-", c)
	}
}
