package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func multiply_uncorrupted_instructions(instructions []byte, regex string) (result int, err error) {
	re := regexp.MustCompile(regex)
	matches := re.FindAllSubmatch(instructions, -1)

	for _, match := range matches {
		a, err := strconv.Atoi(string(match[1]))
		if err != nil {
			return 0, err
		}
		b, err := strconv.Atoi(string(match[2]))
		if err != nil {
			return 0, err
		}
		result += a * b
	}
	return result, err
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := filepath.Join(dir, "input_instructions.txt")

	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// regex won't ignore newlines
	instructions := bytes.ReplaceAll(b, []byte("\n"), []byte(" "))

	result, err := multiply_uncorrupted_instructions(instructions, `mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Multiplied uncorrupted instructions: %d\n", result)

	result, err = multiply_uncorrupted_instructions(instructions, `(?m)(?:don't\(\).*?do\(\).*?)*?mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Multiplied enabled and uncorrupted instructions: %d\n", result)
}
