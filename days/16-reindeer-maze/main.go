package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
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
	Estimate  int
	Parent    *State
}

type Visitations struct {
	Cost    int
	Parents []State
}

func heuristic(a, b Point) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func tracePaths(endStates [](*State), points map[Point]bool) map[Point]bool {
	for _, endState := range endStates {
		points[(*endState).Point] = true
		if (*endState).Parent == nil {
			continue
		}
		tracePaths([]*State{(*endState).Parent}, points)
	}
	// for _, parent := range visitations[endState.Point][endState.Direction].Parents {
	// 	// parent, ok := parents[p]
	// 	// if !ok {
	// 	// 	return points
	// 	// }
	// 	points[parent.Point] = true
	// 	points = tracePaths(parent, visitations, parents, points)
	// }
	return points
}

func AStar(maze [][]byte, start, end Point) (int, map[Point]bool) {
	queue := []State{{
		Point:     start,
		Direction: Point{1, 0},
		Cost:      0,
		Estimate:  heuristic(start, end),
		Parent:    nil,
	}}
	// visitations := make(map[Point]map[Point]Visitations)
	// parent := make(map[State]State)
	// nodeCost := make(map[State]int)
	visited := make(map[Point]map[Direction]int)

	var minCost = math.MaxInt
	// var endState State
	var pathEnds = [](*State){}

	for len(queue) > 0 {
		slices.SortFunc(queue, func(a, b State) int {
			return (a.Cost + a.Estimate) - (b.Cost + b.Estimate)
		})

		currentState := queue[0]
		queue = queue[1:]

		// if _, found := visitations[currentState.Point]; !found {
		// 	visitations[currentState.Point] = make(map[Point]Visitations)
		// }

		// if v, visited := visitations[currentState.Point][currentState.Direction]; !visited || currentState.Cost < v.Cost {
		// 	visitations[currentState.Point][currentState.Direction] = Visitations{currentState.Cost, []State{}}
		// }
		// v, _ := visitations[currentState.Point][currentState.Direction]
		// if currentState.Cost > v.Cost {
		// 	continue
		// } else if currentState.Cost == v.Cost {
		// 	visitations[currentState.Point][currentState.Direction] = Visitations{currentState.Cost, append(slices.Clone(v.Parents), currentState)}
		// }

		if _, ok := visited[currentState.Point]; !ok {
			visited[currentState.Point] = make(map[Direction]int)
		}
		if _, ok := visited[currentState.Point][currentState.Direction]; ok {
			// if cost, ok := visited[currentState.Point][currentState.Direction]; ok && cost < currentState.Cost {
			continue
		} else {
			visited[currentState.Point][currentState.Direction] = currentState.Cost
		}

		if currentState.Point == end {
			if currentState.Cost <= minCost {
				minCost = currentState.Cost
				// endState = currentState
				pathEnds = append(pathEnds, &currentState)

				// var p State = currentState
				// for p.Parent != nil {
				// 	p = *p.Parent
				// 	visited[p.Point][p.Direction] = math.MaxInt
				// 	// pathEnds = append(pathEnds, &p)
				// }
			}
			continue
			// if minCost == -1 || currentState.Cost < minCost {
			// 	minCost = currentState.Cost
			// 	bestPaths = [][]Point{currentState.Path}
			// } else if currentState.Cost == minCost {
			// 	bestPaths = append(bestPaths, slices.Clone(currentState.Path))
			// }
			// continue
		}

		for _, dir := range []Point{{1, 0}, {0, -1}, {-1, 0}, {0, 1}} {
			if dir == currentState.Direction {
				newPoint := Point{currentState.Point.x + dir.x, currentState.Point.y + dir.y}
				if newPoint.y >= 0 && newPoint.y < len(maze) && newPoint.x >= 0 && newPoint.x < len(maze[0]) && maze[newPoint.y][newPoint.x] != '#' {
					queue = append(queue, State{
						Point:     newPoint,
						Direction: currentState.Direction,
						Cost:      currentState.Cost + 1,
						Estimate:  heuristic(newPoint, end),
						// Path:      append(slices.Clone(currentState.Path), newPoint),
						Parent: &currentState,
					})
					// parent[queue[len(queue)-1]] = currentState
				}
			} else if !(dir.x == -currentState.Direction.x && dir.y == -currentState.Direction.y) {
				queue = append(queue, State{
					Point:     currentState.Point,
					Direction: dir,
					Cost:      currentState.Cost + 1000,
					Estimate:  currentState.Estimate,
					// Path:      append(slices.Clone(currentState.Path), currentState.Point),
					Parent: &currentState,
				})
				// parent[queue[len(queue)-1]] = currentState
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

	minScore, trace := AStar(input, startPos, endPos)
	// for p := range trace {
	// 	if p == startPos || p == endPos {
	// 		continue
	// 	}
	// 	AStarSameCost(input, p, endPos)
	// }
	var bestPathTiles int
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if found, ok := trace[Point{x, y}]; found && ok {
				bestPathTiles++
				fmt.Print("O")
			} else {
				fmt.Print(string(input[y][x]))
			}
		}
		fmt.Println()
	}
	fmt.Printf("Lowest score: %d\n", minScore)
	fmt.Printf("Number of tiles on best paths: %d\n", bestPathTiles)
}
