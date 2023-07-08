package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.vidillion.com", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Set custom User-Agent header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_4_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		fmt.Println(k, v)
	}
}
