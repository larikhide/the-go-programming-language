package ch1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// dup1 prints the text of each line that appears more than once
// in the standard input, along with the count of its occurrences.
func dup1(input string) string {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		counts[scanner.Text()]++ //line := input.Text(); conuts[line] = counts[line] + 1
	}
	// ingore input.Err()

	var strBuilder strings.Builder

	for k, v := range counts {
		if v > 1 {
			strBuilder.WriteString(k + "\t" + strconv.Itoa(v))
		}
	}
	return strBuilder.String()
}

// Dup2 prints the text of each line that appears more than once
// in the input. It reads from stdin or from a list of named files.

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines2(os.Stdin, counts)
	} else {
		for _, v := range files {
			f, err := os.Open(v)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines2(f, counts)
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines2(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

	// ignore input.Err()
}

// Dup 3 reads only named files, not standard input, because the ReadFile function requires an argument that is a file name.
// And now the line counting needs to be done only in one place.
// The file is read in its entirety, not line by line as in the previous examples.
func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++

		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// dup4 prints the names of all files that contain duplicate lines.
func dup4() {
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, filename := range files {
		f, err := os.Open(filename)
		defer f.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup4: %v", err)
			continue
		}

		input := bufio.NewScanner(f)
		if input.Scan() {
			counts[input.Text()]++
			fmt.Printf("File %s contains duplicate lines.", filename)
		}
	}
}
