package output

import (
	"fmt"
	"os"
)

// Outputter is the interface for output types
type Outputter interface {
	Write(string, string, bool) error
}

// FileOutput is a struct for writing output to a file
type PathOutput struct{}

// Write writes the output to a file
func (o *PathOutput) Write(fileName string, permutatedURL string, silent bool) error {
	if fileName != "" {
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		if !silent {
			fmt.Println(permutatedURL)
		}

		_, err = file.WriteString(permutatedURL + "\n")
		return err
	} else {
		_, err := fmt.Println(permutatedURL)
		if err != nil {
			return err
		}
	}
	return nil
}
