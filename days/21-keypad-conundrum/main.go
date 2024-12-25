package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strconv"
)

type Point struct {
	x, y int
}

type Move struct {
	dx, dy int
	key    byte
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
	'0': {1, 3},
	'A': {2, 3},
}

var directionalKeypad = map[byte]Point{
	'^': {1, 0},
	'A': {2, 0},
	'<': {0, 1},
	'v': {1, 1},
	'>': {2, 1},
}

var directionalKeypadPresses = map[string]string{
	"A^": "<A",
	"A>": "vA",
	"Av": "<vA",
	"A<": "v<<A",
	"^A": ">A",
	"^>": "v>A",
	"^<": "v<A",
	"^v": "vA",
	"vA": "^>A",
	"v>": ">A",
	"v<": "<A",
	"v^": "^A",
	">A": "^A",
	">^": "<^A",
	">v": "<A",
	"><": "<<A",
	"<A": ">>^A",
	"<^": ">^A",
	"<v": ">A",
	"<>": ">>A",
}

var cache = make(map[string]string)

func dfsFromKeyToKey(from, to byte, keypad *(map[byte]Point)) []string {
	start, end := (*keypad)[from], (*keypad)[to]
	// visited := make(map[Point]bool)
	var paths []string

	var dfs func(current Point, path string)
	dfs = func(current Point, path string) {
		if current == end {
			paths = append(paths, path)
			return
		}

		if len(path) > 5 {
			return
		}

		// visited[current] = true
		// 379A  -->>  ^A<<^^A>>AvvvA
		//       -->>  <A>Av<<AA>^AA>AvAA^A<vAAA>^A
		//       -->>  <v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
		for _, move := range []Move{{-1, 0, '<'}, {0, 1, 'v'}, {1, 0, '>'}, {0, -1, '^'}} {
			next := Point{current.x + move.dx, current.y + move.dy}
			// if !visited[next] && isValid(next, keypad) {
			if isValid(next, keypad) {
				dfs(next, path+string(move.key))
			}
		}
		// visited[current] = false
	}

	dfs(start, "")
	slices.SortFunc(paths, func(a, b string) int {
		return len(a) - len(b)
	})
	return paths
}

func getKeyPressSequence(code string, robots int, maxRobots int) string {
	seq := "A" + code
	result := ""
	// if val, ok := cache[strconv.Itoa(robots)+" "+code]; ok {
	// 	return val
	// }
	for i := 0; i < len(seq)-1; i++ {
		if robots > 0 {
			// if val, ok := cache[strconv.Itoa(robots)+" "+seq[i:i+2]]; ok {
			// 	result += val
			// } else {
			if seq[i] == seq[i+1] {
				result += "A"
				// cache[strconv.Itoa(robots)+" "+seq[i:i+2]] = "A"
			} else {
				result += directionalKeypadPresses[seq[i:i+2]]
				// cache[strconv.Itoa(robots)+" "+seq[i:i+2]] = directionalKeypadPresses[seq[i:i+2]]
			}
			// }
		} else {
			// from, to := (*keypad)[seq[i]], (*keypad)[seq[i+1]]
			// horizontalDist := abs(from.x - to.x)
			// var horizontalKeyPresses string
			// if to.x > from.x {
			// 	horizontalKeyPresses = strings.Repeat(">", horizontalDist)
			// } else {
			// 	horizontalKeyPresses = strings.Repeat("<", horizontalDist)
			// }
			// verticalDist := abs(from.y - to.y)
			// var verticalKeyPresses string
			// if to.y > from.y {
			// 	verticalKeyPresses = strings.Repeat("v", verticalDist)
			// } else {
			// 	verticalKeyPresses = strings.Repeat("^", verticalDist)
			// }

			// result += horizontalKeyPresses + verticalKeyPresses + "A"
			paths := dfsFromKeyToKey(seq[i], seq[i+1], &numericKeypad)
			minPath, min := "", math.MaxInt
			for _, path := range paths {
				// expandedPath := getKeyPressSequence(getKeyPressSequence(path, &directionalKeypad), &directionalKeypad)
				expandedPath := getKeyPressSequence(path, 1, maxRobots)
				// fmt.Println(expandedPath)
				if len(expandedPath) < min {
					min = len(expandedPath)
					minPath = path
				}
			}
			fmt.Println()
			result += minPath + "A"
			// result += paths[0] + "A"
		}
	}
	if robots < maxRobots {
		result = getKeyPressSequence(result, robots+1, maxRobots)
	}
	// cache[strconv.Itoa(robots)+" "+code] = result
	return result
}

func getNumericValue(code []byte) int {
	val, _ := strconv.Atoi(string(code[:len(code)-1]))
	return val
}

func isValid(pos Point, keypad *(map[byte]Point)) bool {
	for _, position := range *keypad {
		if pos == position {
			return true
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "input_codes.txt")
	inputBytes, _ := os.ReadFile(filePath)
	input := bytes.Split(inputBytes, []byte("\r\n"))

	var complexity int
	for _, code := range input {
		complexity += len(getKeyPressSequence(string(code), 0, 2)) * getNumericValue(code)
	}
	fmt.Printf("Complexity: %d\n", complexity)
}
