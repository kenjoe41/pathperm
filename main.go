package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/kenjoe41/pathperm/pkg/config"
	"github.com/kenjoe41/pathperm/pkg/input"
	"github.com/kenjoe41/pathperm/pkg/output"
	"github.com/kenjoe41/pathperm/pkg/permutation"
)

func main() {
	var options config.Config
	options.ParseFlags()

	permuter := &permutation.PathPermuter{}
	inputter := &input.Input{}
	outputter := &output.PathOutput{}

	// Load words from wordlist
	wordsInput, _ := inputter.ReadWordlist(options.InputOpts.Wordlist)

	// Create channels
	urlChan := make(chan string)
	permutatedURLChan := make(chan string)

	var wg sync.WaitGroup
	// Start go routines for permutations and output
	wg.Add(2)

	if wordsInput == nil {
		wordsInput = make(map[string]bool)
	}

	go permuteURLs(permuter, &options.InputOpts, &wordsInput, &urlChan, &permutatedURLChan, &wg)
	go writePermutations(outputter, &options.InputOpts, &permutatedURLChan, &wg)

	// Read input
	urls, err := inputter.ReadInput(options.InputOpts.InFile)
	if err != nil {
		panic(err)
	}

	// Send URLs to channel
	go func() {
		for _, url := range urls {
			urlChan <- url
		}
		close(urlChan)
	}()

	// Wait for all go routines to finish
	wg.Wait()
}

func permuteURLs(permuter *permutation.PathPermuter, conf *config.InputOptions, wordlistMap *map[string]bool, urlChan *chan string, permutatedURLChan *chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range *urlChan {
		err := permutation.GeneratePermutations(url, conf, wordlistMap)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

	}

	if conf.Scheme == "" {
		conf.Scheme = "https"
	}

	var paths []string
	for word := range *wordlistMap {
		path := strings.TrimLeft(word, "/")
		paths = append(paths, path)
		permutatedURL := fmt.Sprintf("%v://%v/%v", conf.Scheme, conf.Domain, path)

		*permutatedURLChan <- permutatedURL
	}
	for i := 0; i < (conf.PermuteLevel - 1); i++ {
		numPaths := len(paths)
		for word := range *wordlistMap {
			for j := 0; j < numPaths; j++ {
				path := paths[j]

				newPath := fmt.Sprintf("%v/%v", path, word)
				newPath = strings.TrimLeft(newPath, "/")
				if !stringInSlice(newPath, paths) {
					paths = append(paths, newPath)
					permutatedURL := fmt.Sprintf("%v://%v/%v", conf.Scheme, conf.Domain, newPath)

					*permutatedURLChan <- permutatedURL
				}
			}

		}
	}
	close(*permutatedURLChan)
}

func writePermutations(outputter output.Outputter, conf *config.InputOptions, permutatedURLChan *chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for permutatedURL := range *permutatedURLChan {
		err := outputter.Write(conf.OfFile, permutatedURL, conf.Silent)
		if err != nil {
			println(err)
			return
		}
	}
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
