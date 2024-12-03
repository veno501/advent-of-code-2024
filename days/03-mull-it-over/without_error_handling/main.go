package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func multiply_uncorrupted_instructions(instructions []byte, regex string) (result int) {
	re := regexp.MustCompile(regex)
	matches := re.FindAllSubmatch(instructions, -1)

	for _, match := range matches {
		result += Must(strconv.Atoi(string(match[1]))) * Must(strconv.Atoi(string(match[2])))
	}
	return result
}

func main() {
	filePath := filepath.Join(Must(os.Getwd()), "..", "input_instructions.txt")
	// regex won't ignore newlines
	instructions := bytes.ReplaceAll(Must(os.ReadFile(filePath)), []byte("\n"), []byte(" "))

	result := multiply_uncorrupted_instructions(instructions, `mul\((\d{1,3}),(\d{1,3})\)`)
	fmt.Printf("Multiplied uncorrupted instructions: %d\n", result)

	result = multiply_uncorrupted_instructions(instructions, `(?m)(?:don't\(\).*?do\(\).*?)*?mul\((\d{1,3}),(\d{1,3})\)`)
	fmt.Printf("Multiplied and enabled uncorrupted instructions: %d\n", result)
}
