package main

import (
	"errors"
	"os"
	"strings"
)

func Loadbanner(file string) (map[rune][]string, error) {
	graph := map[rune][]string{}
	filename, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.New("error reading file")
	}
	if len(filename) == 0 {
		return nil, errors.New("empty file")
	}
	result := strings.Split(string(filename), "\n")
	if len(result) != 856 {
		return nil, errors.New("incomplete file")
	}
	for i := ' '; i <= '~'; i++ {
		start := (i - ' ') * 9
		graph[i] = result[start+1 : start+9]

	}
	return graph, nil
}
