package util

import (
    "io"
    "os"
    "log"
)

// ReadFile Helper method to read a file and return it as a string
func ReadFile(infile string) (string, error) {
	// Open the current file and generate reader
	f, err := os.Open(infile)
    if err != nil {
        return "", err
    }
	defer f.Close()

	// Read the current file
	content, err := io.ReadAll(f)
	// Crash if error
    if err != nil {
        return "", err
    }

	return string(content), nil
}

// WriteFile Helper method to write a file
func WriteFile(outfile, data string) {
	// Create the file
	f, err := os.Create(outfile)
	// Crash if error
    if err != nil {
        log.Fatal(err)
    }

	// Close the file with defer
	defer f.Close()

	// Write the data
	f.WriteString(data)
}

// MakeDir Helper method to create a directory
func MakeDir(dirName string) {
    err := os.Mkdir(dirName, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}

