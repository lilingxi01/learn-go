package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("=== Go File I/O Tutorial ===\n")

	// ===================================
	// 1. Write Entire File
	// ===================================
	fmt.Println("1. Writing entire file:")

	content := []byte("Hello, Go!\nThis is a test file.\n")
	err := os.WriteFile("example.txt", content, 0644)
	if err != nil {
		fmt.Printf("   Error writing file: %v\n", err)
		return
	}
	fmt.Println("   ✓ File written successfully\n")

	// ===================================
	// 2. Read Entire File
	// ===================================
	fmt.Println("2. Reading entire file:")

	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Printf("   Error reading file: %v\n", err)
		return
	}
	fmt.Printf("   Content:\n%s\n", string(data))

	// ===================================
	// 3. Read Line by Line
	// ===================================
	fmt.Println("3. Reading line by line:")

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("   Line %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("   Scanner error: %v\n", err)
	}
	fmt.Println()

	// ===================================
	// 4. Buffered Writing
	// ===================================
	fmt.Println("4. Buffered writing:")

	outFile, err := os.Create("buffered.txt")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	for i := 1; i <= 5; i++ {
		fmt.Fprintf(writer, "Line %d\n", i)
	}
	writer.Flush() // Important! Flush buffer to file
	fmt.Println("   ✓ Buffered write complete\n")

	// ===================================
	// 5. Append to File
	// ===================================
	fmt.Println("5. Appending to file:")

	appendFile, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer appendFile.Close()

	if _, err := appendFile.WriteString("Appended line\n"); err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	fmt.Println("   ✓ Line appended\n")

	// ===================================
	// 6. Check if File Exists
	// ===================================
	fmt.Println("6. Checking if file exists:")

	checkFiles := []string{"example.txt", "nonexistent.txt"}
	for _, filename := range checkFiles {
		if fileExists(filename) {
			fmt.Printf("   ✓ %s exists\n", filename)
		} else {
			fmt.Printf("   ✗ %s does not exist\n", filename)
		}
	}
	fmt.Println()

	// ===================================
	// 7. File Information
	// ===================================
	fmt.Println("7. Getting file information:")

	info, err := os.Stat("example.txt")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Name: %s\n", info.Name())
		fmt.Printf("   Size: %d bytes\n", info.Size())
		fmt.Printf("   Mode: %s\n", info.Mode())
		fmt.Printf("   Modified: %s\n", info.ModTime())
		fmt.Printf("   Is directory: %t\n", info.IsDir())
	}
	fmt.Println()

	// ===================================
	// 8. Copy File
	// ===================================
	fmt.Println("8. Copying file:")

	if err := copyFile("example.txt", "example_copy.txt"); err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Println("   ✓ File copied successfully\n")
	}

	// ===================================
	// 9. Working with Directories
	// ===================================
	fmt.Println("9. Working with directories:")

	// Create directory
	err = os.Mkdir("testdir", 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("   Error creating dir: %v\n", err)
	} else {
		fmt.Println("   ✓ Directory created")
	}

	// Create nested directories
	err = os.MkdirAll("testdir/subdir/nested", 0755)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Println("   ✓ Nested directories created")
	}
	fmt.Println()

	// ===================================
	// 10. List Directory Contents
	// ===================================
	fmt.Println("10. Listing directory contents:")

	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		fmt.Println("    Current directory contents:")
		for _, entry := range entries {
			if entry.IsDir() {
				fmt.Printf("    [DIR]  %s\n", entry.Name())
			} else {
				fmt.Printf("    [FILE] %s\n", entry.Name())
			}
		}
	}
	fmt.Println()

	// ===================================
	// 11. Walk Directory Tree
	// ===================================
	fmt.Println("11. Walking directory tree:")

	err = filepath.Walk("testdir", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("    %s\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	}
	fmt.Println()

	// ===================================
	// 12. Temporary Files
	// ===================================
	fmt.Println("12. Creating temporary file:")

	tmpFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		fmt.Printf("    ✓ Temp file: %s\n", tmpFile.Name())
		tmpFile.WriteString("Temporary data\n")
		tmpFile.Close()
		os.Remove(tmpFile.Name()) // Cleanup
		fmt.Println("    ✓ Temp file cleaned up")
	}
	fmt.Println()

	// ===================================
	// Cleanup
	// ===================================
	fmt.Println("Cleaning up test files...")
	os.Remove("example.txt")
	os.Remove("example_copy.txt")
	os.Remove("buffered.txt")
	os.RemoveAll("testdir")
	fmt.Println("✓ Cleanup complete")

	fmt.Println("\n=== File I/O Tutorial Complete! ===")
}

// fileExists checks if a file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
