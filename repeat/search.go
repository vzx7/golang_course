package repeat

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// читаем с stdin
func SearchV1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(counts, 111)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// читаем с файла
func SearchV2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func SearchV3() {
	type str struct {
		fName string
		value string
	}
	counts := make(map[str]int)
	for _, filename := range os.Args[1:] {
		//https://stackoverflow.com/questions/75206234/for-go-ioutil-readall-ioutil-readfile-ioutil-readdir-deprecated
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Search3: %v", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[str{
				fName: filename,
				value: line,
			}]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("count: %d,\t fileName: %s,\t value: %s\n", n, line.fName, line.value)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
