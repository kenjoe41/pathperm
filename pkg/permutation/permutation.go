package permutation

import (
	"bufio"
	"net/url"
	"os"
	"strings"

	"github.com/kenjoe41/pathperm/pkg/config"
)

// Permuter is the interface for permutation types
type Permuter interface {
	Permute(string, *config.InputOptions, *[]string, *chan string) error
}

// PathPermuter is a struct for permuting path segments in a string
type PathPermuter struct{}

// Permute returns all possible permutations of the path segments in the input string
func GeneratePermutations(inputURL string, conf *config.InputOptions, wordlistMap *map[string]bool) error {
	u, err := url.Parse(inputURL)
	if err != nil {
		return err
	}

	if conf.Domain == "" && u.Host != "" {
		// This should be set once and not to an empty host. Edge case.
		conf.Domain = u.Host
	}

	if conf.Scheme == "" && u.Scheme != "" {
		conf.Scheme = u.Scheme
	}

	path := u.Path
	segments := strings.Split(path, "/")

	// if wordlistMap == nil {
	// 	tempMap := make(map[string]bool)
	// 	wordlistMap = &tempMap
	// }
	for _, word := range segments {
		if word == "" {
			continue
		}
		if !(*wordlistMap)[word] {
			(*wordlistMap)[word] = true
		}
	}

	return nil
}

// ReadInput reads the input urls and wordlist and returns a slice of words
func ReadInput(urls []string, wordlist string) ([]string, error) {
	var words []string

	// read words from wordlist
	if wordlist != "" {
		file, err := os.Open(wordlist)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			words = append(words, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	// add path segments from input urls to words
	for _, urla := range urls {
		u, err := url.Parse(urla)
		if err != nil {
			words = append(words, strings.Split(u.Path, "/")...)
		}
	}

	return words, nil
}
