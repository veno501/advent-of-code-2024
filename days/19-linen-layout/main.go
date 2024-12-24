package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func possiblePatternArrangementsForDesign(patterns []string, design string, cache map[string]int) (count int) {
	if design == "" {
		return 1
	}

	if cachedValue, ok := cache[design]; ok {
		return cachedValue
	}

	for _, pattern := range patterns {
		if len(pattern) <= len(design) && design[:len(pattern)] == pattern {
			count += possiblePatternArrangementsForDesign(patterns, design[len(pattern):], cache)
		}
	}
	cache[design] = count
	return count
}

func main() {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "input_designs.txt")
	inputBytes, _ := os.ReadFile(filePath)
	inputSplit := bytes.Split(inputBytes, []byte("\r\n\r\n"))

	patterns := strings.Split(string(inputSplit[0]), ", ")
	designs := strings.Split(string(inputSplit[1]), "\r\n")

	var possibleDesigns, possibleArrangementsSum int
	for _, design := range designs {
		n := possiblePatternArrangementsForDesign(patterns, design, make(map[string]int))
		if n > 0 {
			possibleDesigns += 1
		}
		possibleArrangementsSum += n
	}
	fmt.Printf("Possible designs: %d\n", possibleDesigns)
	fmt.Printf("Possible arrangements for all designs: %d\n", possibleArrangementsSum)
}
