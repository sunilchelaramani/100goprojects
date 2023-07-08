package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFile(url string, filePath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status: %v", response.Status)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Println("Download successful")
	return nil
}

func main() {
	url := "http://example.com/file.txt" // Replace with the URL of the file you want to download
	filePath := "/path/to/save/file.txt" // Replace with the path where you want to save the downloaded file

	err := downloadFile(url, filePath)
	if err != nil {
		fmt.Println("Download error:", err)
	}
}
