package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sunshineplan/imgconv"
)

func main() {
	// Define flags
	targetFormat := flag.String("target", "png", "target format")
	noDel := flag.Bool("no-delete", false, "do not delete original images")

	// Parse flags
	flag.Parse()

	// Check if help flag is provided
	if flag.NArg() == 0 || flag.Arg(0) == "-h" || flag.Arg(0) == "--help" {
		printUsage()
		return
	}

	// Get source file path
	sourceFile := flag.Arg(0)

	if flag.NArg() != 1 {
		printUsage()
		return
	}

	// Extract file extension using filepath.Ext
	fileExt := filepath.Ext(sourceFile)

	// Validate target format
	formattedTargetFormat := strings.ToLower(*targetFormat)
	allowedFormats := []string{"jpg", "png", "gif", "webp", "pdf"}
	if !contains(allowedFormats, *targetFormat) {
		fmt.Printf("Unsupported target format: %s\n", *targetFormat)
		return
	}

	// Open source image
	src, err := imgconv.Open(sourceFile)
	if err != nil {
		fmt.Println("Error opening source image:", err)
		return
	}

	// Generate target file path with the correct extension
	sourceName := strings.TrimSuffix(sourceFile, fileExt)
	target := fmt.Sprintf("%s.%s", sourceName, formattedTargetFormat)

	// Convert and save image
	outputFile, err := os.Create(target)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Define map for converting target format string to imgconv.Format
	lookUp := map[string]imgconv.Format{
		"jpg": imgconv.JPEG,
		"png": imgconv.PNG,
		"gif": imgconv.GIF,
		"pdf": imgconv.PDF,
	}

	// Convert and write image
	if err := imgconv.Write(outputFile, src, &imgconv.FormatOption{Format: lookUp[formattedTargetFormat]}); err != nil {
		fmt.Println("Error converting and writing image:", err)
		return
	}

	if !*noDel {
		fmt.Println("Original image not deleted.")
	}
}

func printUsage() {
	fmt.Println("Usage: go-convert [options] <source>")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
