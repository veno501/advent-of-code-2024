package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

type Point struct {
	y, x int
}

var directions = []Point{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func find_paths(topography [][]int, x int, y int) (score int, rating int) {
	length_y, length_x := len(topography), len(topography[0])

	var reachable_endpoints []Point

	var dfs func([]Point, int, int, map[Point]bool)
	dfs = func(path []Point, x int, y int, visited map[Point]bool) {
		current_height := topography[y][x]

		if current_height == 9 {
			if !slices.Contains(reachable_endpoints, Point{y, x}) {
				reachable_endpoints = append(reachable_endpoints, Point{y, x})
				score += 1
			}
			rating += 1
			return
		}

		for _, dir := range directions {
			new_x, new_y := x+dir.x, y+dir.y

			if new_x >= 0 && new_x < length_x && new_y >= 0 && new_y < length_y {
				next_height := topography[new_y][new_x]
				next_point := Point{new_y, new_x}

				if next_height == current_height+1 && !visited[next_point] {
					visited[next_point] = true
					dfs(append(path, next_point), new_x, new_y, visited)
					delete(visited, next_point)
				}
			}
		}
	}

	start := Point{y, x}
	visited := map[Point]bool{start: true}
	dfs([]Point{start}, x, y, visited)

	return score, rating
}

func main() {
	filePath := filepath.Join(Must(os.Getwd()), "input_map.txt")
	input := strings.Split(string(Must((os.ReadFile(filePath)))), "\r\n")

	topography := make([][]int, len(input))
	for y := range input {
		topography[y] = make([]int, len(input[y]))
		for x := range input[y] {
			topography[y][x] = Must(strconv.Atoi(string(input[y][x])))
		}
	}

	var score, rating int
	for y := range topography {
		for x := range topography[y] {
			if topography[y][x] == 0 {
				s, r := find_paths(topography, x, y)
				fmt.Println(s, r)
				score += s
				rating += r
			}
		}
	}
	fmt.Printf("Sum of all scores: %d\n", score)
	fmt.Printf("Sum of all ratings: %d\n", rating)
}
