package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	// Create a log file
	logFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create log file:", err)
		return
	}
	defer logFile.Close()

	// Set log output to the log file
	log.SetOutput(logFile)

	// Check if the required number of arguments are provided
	if len(os.Args) < 3 {
		log.Println("Invalid number of arguments")
		showUsage()
		return
	}

	// Read command-line arguments
	videoURL := os.Args[1]
	outputDir := os.Args[2]

	// Validate video URL
	urlRegex := regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
	if !urlRegex.MatchString(videoURL) {
		log.Println("Invalid video URL:", videoURL)
		showUsage()
		return
	}

	// Validate output directory
	fileInfo, err := os.Stat(outputDir)
	if os.IsNotExist(err) {
		log.Println("Output directory does not exist:", outputDir)
		showUsage()
		return
	}
	if !fileInfo.IsDir() {
		log.Println("Output directory is not a valid directory:", outputDir)
		showUsage()
		return
	}

	// If all validations pass, continue with the rest of the code
	fmt.Println("URL:", videoURL)
	fmt.Println("Output Directory:", outputDir)
	// Perform further processing here...
}

func showUsage() {
	fmt.Println("Usage: segmenter <videoURL> <outputDir>")
	fmt.Println("Example: segmenter https://example.com/video.mp4 /path/to/output")
}
