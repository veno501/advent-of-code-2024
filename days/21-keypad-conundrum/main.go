package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

var numericKeypad = map[byte]Point{
	'7': {0, 0},
	'8': {1, 0},
	'9': {2, 0},
	'4': {0, 1},
	'5': {1, 1},
	'6': {2, 1},
	'1': {0, 2},
	'2': {1, 2},
	'3': {2, 2},
	'_': {0, 3},
	'0': {1, 3},
	'A': {2, 3},
}

var directionalKeypad = map[byte]Point{
	'_': {0, 0},
	'^': {1, 0},
	'A': {2, 0},
	'<': {0, 1},
	'v': {1, 1},
	'>': {2, 1},
}

func generateSequenceFromTo(fromKey, toKey byte, keypad map[byte]Point) string {
	from, to := keypad[fromKey], keypad[toKey]
	dx, dy := to.x-from.x, to.y-from.y

	var horizontal string
	if dx > 0 {
		horizontal = strings.Repeat(">", dx)
	} else {
		horizontal = strings.Repeat("<", -dx)
	}
	var vertical string
	if dy > 0 {
		vertical = strings.Repeat("v", dy)
	} else {
		vertical = strings.Repeat("^", -dy)
	}

	gap := keypad['_']
	if dx > 0 && (gap.x != from.x || gap.y != to.y) {
		return vertical + horizontal + "A"
	}
	if gap.x != to.x || gap.y != from.y {
		return horizontal + vertical + "A"
	}
	return vertical + horizontal + "A"
}

func generateSubSequences(code string, keypad map[byte]Point) []string {
	seq := "A" + code
	var result []string
	for i := 0; i < len(seq)-1; i++ {
		result = append(result, generateSequenceFromTo(seq[i], seq[i+1], keypad))
	}
	return result
}

func calcSequenceLength(code string, robots int) int {
	seqCounts := map[string]int{strings.Join(generateSubSequences(code, numericKeypad), ""): 1}

	for iter := 0; iter < robots; iter++ {
		subSeqCounts := map[string]int{}
		for seq, count := range seqCounts {
			for _, subSeq := range generateSubSequences(seq, directionalKeypad) {
				subSeqCounts[subSeq] += count
			}
		}
		seqCounts = subSeqCounts
	}

	var seqLength int
	for seq, count := range seqCounts {
		seqLength += len(seq) * count
	}
	return seqLength
}

func getNumericValue(code []byte) int {
	val, _ := strconv.Atoi(string(code[:len(code)-1]))
	return val
}

func main() {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "input_codes.txt")
	inputBytes, _ := os.ReadFile(filePath)
	input := bytes.Split(inputBytes, []byte("\r\n"))

	var complexity int
	for _, code := range input {
		complexity += calcSequenceLength(string(code), 25) * getNumericValue(code)
	}
	fmt.Printf("Complexity: %d\n", complexity)
}
