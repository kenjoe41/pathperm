package input

import (
	"bufio"
	"os"
	"strings"
)

// Inputter is the interface for input types
type Inputter interface {
	ReadInput(string) ([]string, error)
}

// FileInput is a struct for reading input from a file
type Input struct{}

// ReadInput reads the input from a file and Stdin, and returns the contents as a slice of strings
func (i *Input) ReadInput(fileName string) ([]string, error) {
	var lines []string
	if fileName != "" {
		file, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSuffix(strings.TrimSpace(scanner.Text()), "/")

			lines = append(lines, line)
		}
	}

	// Check for stdin input
	stat, err := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil, err
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	// check there were no errors reading stdin (unlikely)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func (i *Input) ReadWordlist(fileName string) (map[string]bool, error) {
	var lines map[string]bool
	if fileName != "" {
		file, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		lines = make(map[string]bool)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines[strings.TrimSpace(scanner.Text())] = true
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}
	return lines, nil
}
