package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func readInput(reader io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var reports [][]int
	for scanner.Scan() {
		lineScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
		lineScanner.Split(bufio.ScanWords)

		var levels []int
		for lineScanner.Scan() {
			level, err := strconv.Atoi(lineScanner.Text())
			if err != nil {
				return reports, fmt.Errorf("read input: %w", err)
			}
			levels = append(levels, int(level))
		}
		reports = append(reports, levels)
	}
	return reports, nil
}

func are_levels_safe(are_levels_increasing bool, levels []int, can_tolerate_unsafe_level bool) bool {
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if !are_levels_increasing {
			diff = -diff
		}

		if !(diff >= 1 && diff <= 3) {
			if !can_tolerate_unsafe_level {
				return false
			}
			levels = append(levels[:i-1], levels[i:]...)
			i--
			can_tolerate_unsafe_level = false
		}
	}
	return true
}

func count_safe_levels(reports [][]int, can_tolerate_unsafe_level bool) (int, error) {
	safe_count := 0
	for _, levels := range reports {
		if are_levels_safe(levels[1] > levels[0], levels, can_tolerate_unsafe_level) {
			safe_count++
		}
	}
	return safe_count, nil
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := filepath.Join(dir, "input_levels.txt")

	fileReader, err := os.Open(filePath)
	if err != nil {
		log.Fatal(fmt.Errorf("open file %s: %w", filePath, err))
	}

	reports, err := readInput(fileReader)
	if err != nil {
		log.Fatal(err)
	}

	safe_count, err := count_safe_levels(reports, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Safe levels: %d\n", safe_count)

	safe_count_with_tolerance, err := count_safe_levels(reports, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Safe levels with tolerance: %d\n", safe_count_with_tolerance)
}
