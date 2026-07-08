package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(fileName string) (map[rune][]string, error) {
	graphic := make(map[rune][]string)
	fontFile, err := os.ReadFile(fileName)
	fontFile = []byte(strings.ReplaceAll(string(fontFile), "\r\n", "\n"))
	if err != nil {
		return nil, err
	}
	if len(fontFile) == 0 {
		return nil, errors.New("empty file")
	}
	fontFileLines := strings.Split(string(fontFile), "\n")
	if len(fontFileLines) != 856 {
		return nil, errors.New("invalid file content")
	}
	for i := ' '; i <= '~'; i++ {
		start := (i - ' ') * 9
		graphic[i] = fontFileLines[start+1 : start+9]
	}
	return graphic, nil
}
