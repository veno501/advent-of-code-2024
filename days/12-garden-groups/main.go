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

type Point struct {
	x int
	y int
}

func isBorder(farm [][]byte, pos Point, prevPos Point) bool {
	if pos == prevPos || prevPos.x < 0 || prevPos.x >= len(farm[0]) || prevPos.y < 0 || prevPos.y >= len(farm) {
		return false
	}
	return pos.x < 0 || pos.x >= len(farm[0]) || pos.y < 0 || pos.y >= len(farm) || farm[pos.y][pos.x] != farm[prevPos.y][prevPos.x]
}

func traverseGarden(farm [][]byte, visited map[Point]bool, pos Point, prevPos Point, straightLinesOnly bool) (area int, perimeter int) {
	is_visited, ok := visited[pos]
	if isBorder(farm, pos, prevPos) {
		var checkPos, prevCheckPos Point
		if pos.x == prevPos.x {
			checkPos = Point{pos.x - 1, pos.y}
			prevCheckPos = Point{prevPos.x - 1, prevPos.y}
		}
		if pos.y == prevPos.y {
			checkPos = Point{pos.x, pos.y - 1}
			prevCheckPos = Point{prevPos.x, prevPos.y - 1}
		}
		if straightLinesOnly && isBorder(farm, checkPos, prevCheckPos) && farm[prevCheckPos.y][prevCheckPos.x] == farm[prevPos.y][prevPos.x] {
			return 0, 0
		}
		return 0, 1
	}
	if ok && is_visited {
		return 0, 0
	}
	visited[pos] = true
	var areaSum, perimeterSum int
	for _, dir := range []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		a, p := traverseGarden(farm, visited, Point{pos.x + dir.x, pos.y + dir.y}, pos, straightLinesOnly)
		areaSum += a
		perimeterSum += p
	}
	return 1 + areaSum, 0 + perimeterSum
}

func main() {
	filePath := filepath.Join(Must(os.Getwd()), "input_gardens.txt")
	input := bytes.Split(Must(os.ReadFile(filePath)), []byte("\r\n"))

	visitedForPerimiter, visitedForStraightLines := map[Point]bool{}, map[Point]bool{}
	var priceOfPerimeter, priceOfStraightLines int
	for y := range input {
		for x := range input[y] {
			a, p := traverseGarden(input, visitedForPerimiter, Point{x, y}, Point{x, y}, false)
			priceOfPerimeter += a * p
			_, ok := visitedForStraightLines[Point{x, y}]
			if ok {
				continue
			}
			a, p = traverseGarden(input, visitedForStraightLines, Point{x, y}, Point{x, y}, true)
			fmt.Println(a, p)
			priceOfStraightLines += a * p
		}
	}

	fmt.Printf("Price of fences for all gardens: %d\n", priceOfPerimeter)
	fmt.Printf("Price of fences for gardens, counting straight lines only: %d\n", priceOfStraightLines)
}
