package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

type Point struct {
	x, y int
}

type Node struct {
	point    Point
	priority int
}

var directions = []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func getTrack(grid [][]byte) []Point {
	var start Point
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'S' {
				start = Point{x, y}
				break
			}
		}
	}

	track := []Point{start}
	x, y := start.x, start.y
	for grid[y][x] != 'E' {
		for _, dir := range directions {
			newPos := Point{x + dir.x, y + dir.y}
			if (len(track) < 2 || track[len(track)-2] != newPos) && grid[newPos.y][newPos.x] != '#' {
				track = append(track, newPos)
				x, y = newPos.x, newPos.y
				break
			}
		}
	}
	return track
}

func cheats(track []Point, maxDistance int) []int {
	saved := []int{}
	for idx1, point1 := range track {
		for idx2 := idx1 + 3; idx2 < len(track); idx2++ {
			point2 := track[idx2]
			manhattanDistance := abs(point2.x-point1.x) + abs(point2.y-point1.y)

			// if manhattan distance between the two nodes in the track is less than the amount of on-track nodes between them, it's a shortcut
			distanceSavedWithCheat := idx2 - idx1 - manhattanDistance
			if manhattanDistance <= maxDistance && distanceSavedWithCheat > 0 {
				saved = append(saved, distanceSavedWithCheat)
			}
		}
	}
	return saved
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "input_racetrack.txt")
	inputBytes, _ := os.ReadFile(filePath)
	input := bytes.Split(inputBytes, []byte("\r\n"))

	track := getTrack(input)

	// challenge 1

	var numOfCheatsThatSave100Distance int
	for _, saved := range cheats(track, 2) {
		if saved >= 100 {
			numOfCheatsThatSave100Distance++
		}
	}
	fmt.Printf("Number of cheats that save 100 distance by disabling collision for 2 distance: %d\n", numOfCheatsThatSave100Distance)

	// challenge 2

	numOfCheatsThatSave100Distance = 0
	for _, saved := range cheats(track, 20) {
		if saved >= 100 {
			numOfCheatsThatSave100Distance++
		}
	}
	fmt.Printf("Number of cheats that save 100 distance by disabling collision for 20 distance: %d\n", numOfCheatsThatSave100Distance)
}
