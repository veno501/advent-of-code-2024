package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Node struct {
	point    Point
	priority int
}

var directions = []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func dijkstra(grid [][]byte, start, end Point) int {
	n := len(grid)
	m := len(grid[0])
	distances := make([][]int, n)
	for i := range distances {
		distances[i] = make([]int, m)
		for j := range distances[i] {
			distances[i][j] = math.MaxInt
		}
	}
	distances[start.x][start.y] = 0

	queue := []Node{{point: start, priority: 0}}

	for len(queue) > 0 {
		slices.SortFunc(queue, func(i, j Node) int {
			return i.priority - j.priority
		})

		current := queue[0]
		queue = queue[1:]
		currPoint := current.point

		if currPoint == end {
			return distances[end.x][end.y]
		}

		for _, dir := range directions {
			nx, ny := currPoint.x+dir.x, currPoint.y+dir.y
			if nx >= 0 && nx < n && ny >= 0 && ny < m && grid[nx][ny] == '.' {
				newDist := distances[currPoint.x][currPoint.y] + 1
				if newDist < distances[nx][ny] {
					distances[nx][ny] = newDist
					queue = append(queue, Node{point: Point{nx, ny}, priority: newDist})
				}
			}
		}
	}

	return -1
}

func bytesFalling(grid *([][]byte), coords []Point, amount int) {
	for i := 0; i < len(coords) && i < amount; i++ {
		(*grid)[coords[i].y][coords[i].x] = '#'
	}
}

func main() {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "input_bytes.txt")
	inputBytes, _ := os.ReadFile(filePath)
	input := bytes.Split(inputBytes, []byte("\r\n"))

	gridSize := 71
	initialBytesToFall := 1024

	coords := make([]Point, len(input))
	for i := 0; i < len(input); i++ {
		split := strings.Split(string(input[i]), ",")
		X, _ := strconv.Atoi(split[0])
		Y, _ := strconv.Atoi(split[1])
		coords[i] = Point{x: X, y: Y}
	}

	for i := initialBytesToFall; i < len(coords); i++ {

		grid := make([][]byte, gridSize)
		for y := 0; y < len(grid); y++ {
			grid[y] = make([]byte, gridSize)
			for x := 0; x < len(grid[y]); x++ {
				grid[y][x] = '.'
			}
		}
		bytesFalling(&grid, coords, i)

		minPath := dijkstra(grid, Point{x: 0, y: 0}, Point{x: gridSize - 1, y: gridSize - 1})

		// challenge 1

		if i == initialBytesToFall {
			fmt.Printf("Shortest path to the exit: %d\n", minPath)
		}

		// challenge 2

		if minPath == -1 {
			fmt.Printf("Coordinates of the first byte that will prevent the exit from being reachable: %d,%d\n", coords[i-1].x, coords[i-1].y)
			break
		}
	}
}
