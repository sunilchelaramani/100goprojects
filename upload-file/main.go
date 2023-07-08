package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func uploadFile(url string, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	writer.Close()

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("upload failed with status: %v", response.Status)
	}

	fmt.Println("Upload successful")
	return nil
}

func main() {
	url := "http://example.com/upload" // Replace with your upload endpoint URL
	filePath := "/path/to/file.txt"    // Replace with the path to the file you want to upload

	err := uploadFile(url, filePath)
	if err != nil {
		fmt.Println("Upload error:", err)
	}
}
