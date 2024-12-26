package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
)

type Point struct {
	x, y int
}

type Direction = Point

// State represents the Reindeer's current position, facing direction, and costs
type State struct {
	Point     Point
	Direction Direction
	Cost      int
	Parent    *State
}

func tracePaths(endStates [](*State), points map[Point]bool) map[Point]bool {
	for _, endState := range endStates {
		points[(*endState).Point] = true
		if (*endState).Parent == nil {
			continue
		}
		tracePaths([]*State{(*endState).Parent}, points)
	}
	return points
}

func AStar(maze [][]byte, start, end Point) (int, map[Point]bool) {
	queue := []State{{
		Point:     start,
		Direction: Point{1, 0},
		Cost:      0,
		Parent:    nil,
	}}
	visited := make(map[Point]map[Direction]int)
	var minCost = math.MaxInt
	var pathEnds = [](*State){}

	for len(queue) > 0 {
		currentState := queue[0]
		queue = queue[1:]

		if _, ok := visited[currentState.Point]; !ok {
			visited[currentState.Point] = make(map[Direction]int)
		}
		if cost, ok := visited[currentState.Point][currentState.Direction]; ok && cost < currentState.Cost {
			continue
		} else {
			visited[currentState.Point][currentState.Direction] = currentState.Cost
		}

		if currentState.Point == end {
			if currentState.Cost == minCost {
				minCost = currentState.Cost
				pathEnds = append(pathEnds, &currentState)
			} else if currentState.Cost < minCost {
				minCost = currentState.Cost
				pathEnds = [](*State){&currentState}
			}
			continue
		}

		for _, dir := range []Point{{1, 0}, {0, -1}, {-1, 0}, {0, 1}} {
			if dir == currentState.Direction {
				newPoint := Point{currentState.Point.x + dir.x, currentState.Point.y + dir.y}
				if newPoint.y >= 0 && newPoint.y < len(maze) && newPoint.x >= 0 && newPoint.x < len(maze[0]) && maze[newPoint.y][newPoint.x] != '#' {
					queue = append(queue, State{
						Point:     newPoint,
						Direction: currentState.Direction,
						Cost:      currentState.Cost + 1,
						Parent:    &currentState,
					})
				}
			} else if !(dir.x == -currentState.Direction.x && dir.y == -currentState.Direction.y) {
				queue = append(queue, State{
					Point:     currentState.Point,
					Direction: dir,
					Cost:      currentState.Cost + 1000,
					Parent:    &currentState,
				})
			}
		}
	}
	trace := tracePaths(pathEnds, make(map[Point]bool))

	return minCost, trace
}

func main() {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "input_maze.txt")
	inputBytes, _ := os.ReadFile(filePath)
	input := bytes.Split(inputBytes, []byte("\r\n"))

	var startPos, endPos Point
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == 'S' {
				startPos = Point{x, y}
			}
			if input[y][x] == 'E' {
				endPos = Point{x, y}
			}
		}
	}

	minScore, trace := AStar(input, endPos, startPos)

	var bestPathTiles int
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if found, ok := trace[Point{x, y}]; found && ok {
				bestPathTiles++
			}
		}
	}
	fmt.Printf("Lowest score: %d\n", minScore)
	fmt.Printf("Number of tiles on best paths: %d\n", bestPathTiles)
}
