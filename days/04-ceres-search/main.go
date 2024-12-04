package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

type DiagonalDirection string

const (
	〵 DiagonalDirection = "〵"
	〳 DiagonalDirection = "〳"
)

func check_for_word(graph [][]byte, word []byte, x int, y int, dir_x int, dir_y int) bool {
	for i := 0; i < len(word); i++ {
		X := x + i*dir_x
		Y := y + i*dir_y
		if X < 0 || X >= len(graph) || Y < 0 || Y >= len(graph[X]) {
			return false
		}
		if graph[X][Y] != word[i] {
			return false
		}
	}
	return true
}

func find_all_xmas(graph [][]byte) (xmas_count int, mas_count int) {
	xmas := []byte("XMAS")
	mas := []byte("MAS")

	for x := range graph {
		for y := range graph[x] {
			if graph[x][y] == xmas[0] {
				for _, dir := range [][]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}} {
					if check_for_word(graph, xmas, x, y, dir[0], dir[1]) {
						xmas_count += 1
					}
				}
			}

			if graph[x][y] == mas[1] {
				if (check_for_word(graph, mas, x-1, y+1, 1, -1) || check_for_word(graph, mas, x+1, y-1, -1, 1)) &&
					(check_for_word(graph, mas, x+1, y+1, -1, -1) || check_for_word(graph, mas, x-1, y-1, 1, 1)) {
					mas_count += 1
				}
			}
		}
	}
	return xmas_count, mas_count
}

func main() {
	filePath := filepath.Join(Must(os.Getwd()), "input_graph.txt")
	graph := bytes.Split(Must(os.ReadFile(filePath)), []byte("\r\n"))

	xmas_count, mas_count := find_all_xmas(graph)
	fmt.Printf("XMAS appears %d times\n", xmas_count)
	fmt.Printf("MAS appears %d times\n", mas_count)
}
