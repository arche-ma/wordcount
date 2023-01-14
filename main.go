package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	line := readFlag()
	line = strings.TrimSpace(line)
	wordsCount := parseLine(&line)
	fmt.Println(wordsCount)
}

func readFlag() string {
	flag.Parse()
	if len(flag.Args()) != 0 {
		return flag.Args()[0]
	}
	return ""
}

func parseLine(line *string) int {
	if *line == "" {
		return 0
	}
	return len(strings.Split(*line, " "))
}
