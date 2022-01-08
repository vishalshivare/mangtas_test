package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	filename := "words_and_spaces.txt"
	fmt.Println("Reading file: ", filename)

	fh, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open file '%v': %v", filename, err)
		os.Exit(1)
	}
	reader := bufio.NewReader(fh)
	counter := make(map[string]int)
	for {
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(line)
		for _, word := range fields {
			word = strings.ToLower(word)
			counter[word]++
		}
		if line == "" {
			break
		}
	}
	// Creating new slice for keys
	words := make([]string, 0, len(counter))
	for word := range counter {
		words = append(words, word)
	}

	// sort the key slice on basis of there count
	sort.Slice(words, func(i, j int) bool {
		return counter[words[i]] > counter[words[j]]
	})

	fmt.Println("Printing Top 10 words with counter:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v %v\n", words[i], counter[words[i]])
	}
}

