package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

type Point struct {
	x int
	y int
}

func tryMove(warehouse [][]byte, pos Point, dir Point) bool {
	if checkCanMove(warehouse, pos, dir) {
		move(warehouse, pos, dir)
		return true
	}
	return false
}

func checkCanMove(warehouse [][]byte, pos Point, dir Point) bool {
	newPos := Point{pos.x + dir.x, pos.y + dir.y}
	if warehouse[newPos.y][newPos.x] == 'O' {
		return checkCanMove(warehouse, newPos, dir)
	}
	if warehouse[newPos.y][newPos.x] == '[' {
		if dir.y != 0 {
			return checkCanMove(warehouse, newPos, dir) && checkCanMove(warehouse, Point{newPos.x + 1, newPos.y}, dir)
		}
		return checkCanMove(warehouse, newPos, dir)
	}
	if warehouse[newPos.y][newPos.x] == ']' {
		if dir.y != 0 {
			return checkCanMove(warehouse, newPos, dir) && checkCanMove(warehouse, Point{newPos.x - 1, newPos.y}, dir)
		}
		return checkCanMove(warehouse, newPos, dir)
	}
	if warehouse[newPos.y][newPos.x] == '.' {
		return true
	}
	return false
}

func move(warehouse [][]byte, pos Point, dir Point) {
	newPos := Point{pos.x + dir.x, pos.y + dir.y}
	if warehouse[newPos.y][newPos.x] == 'O' {
		move(warehouse, newPos, dir)
	}
	if warehouse[newPos.y][newPos.x] == '[' {
		if dir.y != 0 {
			move(warehouse, Point{newPos.x + 1, newPos.y}, dir)
		}
		move(warehouse, newPos, dir)
	}
	if warehouse[newPos.y][newPos.x] == ']' {
		if dir.y != 0 {
			move(warehouse, Point{newPos.x - 1, newPos.y}, dir)
		}
		move(warehouse, newPos, dir)
	}
	warehouse[newPos.y][newPos.x] = warehouse[pos.y][pos.x]
	warehouse[pos.y][pos.x] = '.'
}

func main() {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "input_warehouse.txt")
	inputBytes, _ := os.ReadFile(filePath)
	input := bytes.Split(inputBytes, []byte("\r\n\r\n"))
	warehouse := bytes.Split(input[0], []byte("\r\n"))

	wideWarehouse := make([][]byte, len(warehouse))
	for y := 0; y < len(warehouse); y++ {
		wideWarehouse[y] = make([]byte, len(warehouse[y])*2)
		for x := 0; x < len(warehouse[y]); x++ {
			firstChar := warehouse[y][x]
			if firstChar == '@' {
				wideWarehouse[y][x*2] = '@'
				wideWarehouse[y][x*2+1] = '.'
			} else if firstChar == 'O' {
				wideWarehouse[y][x*2] = '['
				wideWarehouse[y][x*2+1] = ']'
			} else {
				wideWarehouse[y][x*2] = firstChar
				wideWarehouse[y][x*2+1] = firstChar
			}
		}
	}

	var rPos Point
	for y := 0; y < len(warehouse); y++ {
		for x := 0; x < len(warehouse[y]); x++ {
			if warehouse[y][x] == '@' {
				rPos = Point{x, y}
			}
		}
	}
	var rWidePos Point
	for y := 0; y < len(wideWarehouse); y++ {
		for x := 0; x < len(wideWarehouse[y]); x++ {
			if wideWarehouse[y][x] == '@' {
				rWidePos = Point{x, y}
			}
		}
	}

	for _, instruction := range input[1] {
		var dir Point
		if instruction == '^' {
			dir = Point{0, -1}
		} else if instruction == 'v' {
			dir = Point{0, 1}
		} else if instruction == '<' {
			dir = Point{-1, 0}
		} else if instruction == '>' {
			dir = Point{1, 0}
		} else {
			continue
		}

		if tryMove(warehouse, rPos, dir) {
			rPos = Point{rPos.x + dir.x, rPos.y + dir.y}
		}

		if tryMove(wideWarehouse, rWidePos, dir) {
			rWidePos = Point{rWidePos.x + dir.x, rWidePos.y + dir.y}
		}
	}

	for y := 0; y < len(warehouse); y++ {
		for x := 0; x < len(warehouse[y]); x++ {
			if warehouse[y][x] == '.' {
				fmt.Printf("%c", warehouse[y][x])
			} else {
				fmt.Printf("%c", warehouse[y][x])
			}
		}
		fmt.Println()
	}
	for y := 0; y < len(wideWarehouse); y++ {
		for x := 0; x < len(wideWarehouse[y]); x++ {
			if wideWarehouse[y][x] == '.' {
				fmt.Printf("%c", wideWarehouse[y][x])
			} else {
				fmt.Printf("%c", wideWarehouse[y][x])
			}
		}
		fmt.Println()
	}

	coordsSum := 0
	for y := 0; y < len(warehouse); y++ {
		for x := 0; x < len(warehouse[y]); x++ {
			if warehouse[y][x] == 'O' {
				coordsSum += 100*y + x
			}
		}
	}
	fmt.Printf("Sum of the coordinates of all the boxes: %d\n", coordsSum)

	wideCoordsSum := 0
	for y := 0; y < len(wideWarehouse); y++ {
		for x := 0; x < len(wideWarehouse[y]); x++ {
			if wideWarehouse[y][x] == '[' {
				wideCoordsSum += 100*y + x
			}
		}
	}
	fmt.Printf("Sum of the coordinates of all the wide boxes in the wide warehouse: %d\n", wideCoordsSum)
}
