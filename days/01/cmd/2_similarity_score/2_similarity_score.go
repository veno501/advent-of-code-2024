package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"

	"golang.org/x/exp/constraints"
)

type SignedNumber interface {
	constraints.Signed | constraints.Float
}

type Number interface {
	constraints.Integer | constraints.Float
}

func readInput(reader io.Reader) ([]int32, []int32, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	var l1, l2 []int32
	for {
		if !scanner.Scan() {
			break
		}
		n1, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return l1, l2, fmt.Errorf("read input: %w", err)
		}
		l1 = append(l1, int32(n1))

		if !scanner.Scan() {
			break
		}
		n2, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return l1, l2, fmt.Errorf("read input: %w", err)
		}
		l2 = append(l2, int32(n2))
	}

	return l1, l2, nil
}

func compare_lists(filePath string) (int, error) {
	fileReader, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("open file %s: %w", filePath, err)
	}

	l1, l2, err := readInput(fileReader)
	if err != nil {
		return 0, fmt.Errorf("read input from file %s: %w", filePath, err)
	}

	if len(l1) == 0 {
		return 0, errors.New("input was parsed as empty")
	}
	if len(l1) != len(l2) {
		return 0, fmt.Errorf(
			"input list sizes %d, %d do not match", len(l1), len(l2),
		)
	}

	similarity := 0
	occurences := map[int32]int{}
	for i := range l2 {
		if slices.Contains(l1, l2[i]) {
			occurences[l2[i]] += 1
		}
	}

	for key, occuredTimes := range occurences {
		similarity += int(key) * occuredTimes
	}

	return similarity, nil
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := filepath.Join(dir, "assets", "input_list.txt")

	similarity, err := compare_lists(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Similarity: %d\n", similarity)
}
