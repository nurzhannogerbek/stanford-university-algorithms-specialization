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
	url := flag.String("url", "https://d3c33hcgiwev3.cloudfront.net/_fe8d0202cd20a808db6a4d5d06be62f4_clustering_big.txt?Expires=1733184000&Signature=bF9fl-vxZuVopJdZfYVaKXV84nLQj4usxkcI67aIdfQRRl5L0j4JCqNW39~R0qTiInE4BI75mNIJ2VZedbsqJc3vCgjsaf5uxTKLixzjrYWWSWrzTJ~GOV3x0roVplwxbT0TybFcGXEP0z9k2bmNA9Sn0hPuvtHkHGyXFTC33Oo_&Key-Pair-Id=APKAJLTNE6QMUY6HBC5A", "URL to download the file from.")
	outputFile := flag.String("output", "clustering_big.txt", "File name to save the downloaded content.")
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
