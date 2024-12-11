package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

type Key struct {
	rock       int
	iterations int
}

var cacheDict map[Key]int = make(map[Key]int)

func countRocks(rock int, iterations int) int {
	cachedCount, hasKey := cacheDict[Key{rock, iterations}]
	if hasKey {
		return cachedCount
	}

	Cache := func(count int) int {
		cacheDict[Key{rock, iterations}] = count
		return count
	}

	if iterations == 0 {
		return 1
	}

	if rock == 0 {
		return Cache(countRocks(1, iterations-1))
	}

	rockString := strconv.Itoa(rock)
	if len(rockString)%2 == 0 {
		leftHalf := Must(strconv.Atoi(rockString[:len(rockString)/2]))
		rightHalf := Must(strconv.Atoi(rockString[len(rockString)/2:]))
		return Cache(countRocks(leftHalf, iterations-1) + countRocks(rightHalf, iterations-1))
	}

	return Cache(countRocks(rock*2024, iterations-1))
}

func main() {
	filePath := filepath.Join(Must(os.Getwd()), "input_rocks.txt")
	line := strings.Split(string(Must((os.ReadFile(filePath)))), " ")

	var rockCount25, rockCount75 int
	for _, rock := range line {
		rockCount25 += countRocks(Must(strconv.Atoi(rock)), 25)
		rockCount75 += countRocks(Must(strconv.Atoi(rock)), 75)
	}
	fmt.Printf("Number of rocks after 25 iterations: %d\n", rockCount25)
	fmt.Printf("Number of rocks after 75 iterations: %d\n", rockCount75)
}
