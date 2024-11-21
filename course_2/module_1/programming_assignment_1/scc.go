package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Define parameters for the program: URL and output file name.
	// The URL is the address from which the file will be downloaded.
	url := flag.String("url", "https://d3c33hcgiwev3.cloudfront.net/_410e934e6553ac56409b2cb7096a44aa_SCC.txt?Expires=1732320000&Signature=KJWsQopdph4scWJbt6-3ShgYA9sYtMFT2JYiGfjILBOVBAonDwSymetW8Oghe1swytfWEfKARNKQzXM8jRBmozVKyl~7ASj~7sURYXLPj9AKsktY-s169h-xtjoKYfpzRFb9ewcY~h98IMH8z44savQevRxj35lEslWsbvPSoNE_&Key-Pair-Id=APKAJLTNE6QMUY6HBC5A", "URL to download the file from.")
	outputFile := flag.String("output", "SCC.txt", "File name to save the downloaded content.")
	flag.Parse()

	// Check if the URL parameter is provided.
	if *url == "" {
		fmt.Println("Error: URL must be provided.")
		return
	}

	// Get the current working directory to ensure the file is created in the same location as the program.
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error retrieving current directory:", err)
		return
	}
	filePath := filepath.Join(currentDir, *outputFile)

	// Create the file where the downloaded content will be saved.
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file '%s': %v\n", filePath, err)
		return
	}
	// Properly handle the error from file.Close().
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Error closing file '%s': %v\n", filePath, closeErr)
		}
	}()

	// Send an HTTP GET request to download the file.
	resp, err := http.Get(*url)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}
	// Properly handle the error from resp.Body.Close().
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Printf("Error closing response body: %v\n", closeErr)
		}
	}()

	// Check if the HTTP response status code indicates success.
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Printf("Error: HTTP status %d received.\n", resp.StatusCode)
		return
	}

	// Save the content of the response to the output file.
	fmt.Printf("Downloading file to '%s'...\n", filePath)
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	// Print success message after file is saved.
	fmt.Println("File downloaded and saved successfully at:", filePath)
}
