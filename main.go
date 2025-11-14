package main

import downdetector "hello_world/ratings/downDetector"

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// r := ratings.NewRating(1)

	// for {
	// 	fmt.Print("\nEnter user ID (or type 'exit' to quit): ")
	// 	input, _ := reader.ReadString('\n')
	// 	input = strings.TrimSpace(input)
	// 	if strings.ToLower(input) == "exit" {
	// 		break
	// 	}

	// 	uid, err := strconv.Atoi(input)
	// 	if err != nil {
	// 		fmt.Println("❌ Invalid user ID. Please enter a number.")
	// 		continue
	// 	}

	// 	fmt.Print("Enter rating (1 to 5): ")
	// 	ratingStr, _ := reader.ReadString('\n')
	// 	ratingStr = strings.TrimSpace(ratingStr)
	// 	rating, err := strconv.ParseFloat(ratingStr, 64)
	// 	if err != nil {
	// 		fmt.Println("❌ Invalid rating. Please enter a number.")
	// 		continue
	// 	}

	// 	fmt.Print("Enter comment: ")
	// 	comment, _ := reader.ReadString('\n')
	// 	comment = strings.TrimSpace(comment)

	// 	e := r.AddRating(uid, rating, comment)
	// 	if e != nil {
	// 		fmt.Println("⚠️", e)
	// 		continue
	// 	}

	// 	fmt.Println("\n✅ Rating added successfully!")
	// 	fmt.Println(r)
	// }

	// fmt.Println("\n👋 Exiting program.")
	// fmt.Println(r)

	urls := []string{
		"https://gdg.community.dev/gdg-cochin/",
		"https://golang.org",
		"https://httpstat.us/500",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.twitter.com/",
		"https://www.instagram.com/",
		"https://site-not-present.io",
		"https://www.youtube.com/",
		"https://www.linkedin.com/",
		"https://www.github.com/",
		"https://www.stackoverflow.com/",
		"https://www.reddit.com/",
	}

	for _, url := range urls {
		downdetector.CheckStatus(url)
	}
}
