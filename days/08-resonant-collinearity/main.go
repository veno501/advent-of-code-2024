package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/quartercastle/vector"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

type vec = vector.Vector

func can_add_antinode(antinode_positions []vec, map_bounds vec, antinode vec) bool {
	return antinode.X() >= 0 && antinode.X() < map_bounds.X() && antinode.Y() >= 0 && antinode.Y() < map_bounds.Y() &&
		!slices.ContainsFunc(antinode_positions, func(v vec) bool { return v.Equal(antinode) })
}

func main() {
	filePath := filepath.Join(Must(os.Getwd()), "input_map.txt")
	input := strings.Split(string(Must((os.ReadFile(filePath)))), "\r\n")

	node_positions := map[rune]([]vec){}
	for y, line := range input {
		for x, char := range line {
			if char == '.' {
				continue
			}

			node_positions[char] = append(node_positions[char], vec{float64(x), float64(y)})
		}
	}

	var antinode_positions []vec
	map_bounds := vec{float64(len(input[0])), float64(len(input))}
	for _, positions := range node_positions {
		for i := range positions {
			for j := i + 1; j < len(positions); j++ {
				v1, v2 := positions[i], positions[j]
				for d := 0; d < 100; d++ {
					antinode1 := v1.Add((v1.Sub(v2)).Scale(float64(d)))
					antinode2 := v2.Add((v2.Sub(v1)).Scale(float64(d)))
					if can_add_antinode(antinode_positions, map_bounds, antinode1) {
						antinode_positions = append(antinode_positions, antinode1)
					}
					if can_add_antinode(antinode_positions, map_bounds, antinode2) {
						antinode_positions = append(antinode_positions, antinode2)
					}
				}
			}
		}
	}

	fmt.Printf("Number of unique positions of antinodes: %d\n", len(antinode_positions))
}
