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
		n2, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return l1, l2, fmt.Errorf("read input: %w", err)
		}
		l2 = append(l2, int32(n2))
	}

	return l1, l2, nil
}

func add_difference[T SignedNumber, U Number](n1 T, n2 T, diff *U) {
	d := n1 - n2
	if d < 0 {
		d = -d
	}
	*diff += U(d)
}

func compare_lists(filePath string) (uint64, error) {
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

	slices.Sort(l1)
	slices.Sort(l2)

	var diff uint64 = 0
	for i := len(l1) - 1; i >= 0; i-- {
		add_difference(l1[i], l2[i], &diff)
	}

	return diff, nil
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := filepath.Join(dir, "assets", "input_list.txt")

	diff, err := compare_lists(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Difference: %d\n", diff)
}
