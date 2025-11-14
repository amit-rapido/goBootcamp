package downdetector

import (
	"fmt"
	"net/http"
)

func CheckStatus(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("URL : ", url, " Status Down")
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("URL : ", url, " Status Up")
		return
	}
}
