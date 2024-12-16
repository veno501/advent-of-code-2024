package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/quartercastle/vector"
)

type vec = vector.Vector

type Robot struct {
	pos vec
	vel vec
}

// func drawPicture(robots []Robot, bounds vec) {
// 	picture := make([][]byte, int(bounds.Y()))
// 	for y := 0; y < int(bounds.Y()); y++ {
// 		picture[y] = make([]byte, int(bounds.X()))
// 		for x := 0; x < int(bounds.X()); x++ {
// 			picture[y][x] = '.'
// 		}
// 	}
// 	for _, robot := range robots {
// 		x, y := int(robot.pos.X()), int(robot.pos.Y())
// 		num, _ := strconv.Atoi(string(picture[y][x]))
// 		num += 1
// 		picture[y][x] = strconv.Itoa(num)[0]
// 	}
// 	for y := 0; y < int(bounds.Y()); y++ {
// 		for x := 0; x < int(bounds.X()); x++ {
// 			fmt.Printf("%c", picture[y][x])
// 		}
// 		fmt.Println()
// 	}
// }

func main() {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "input_robots.txt")
	inputBytes, _ := os.ReadFile(filePath)
	input := strings.Split(string(inputBytes), "\r\n")

	robots := make([]Robot, len(input))
	for i, line := range input {
		split := strings.Split(line[2:], " v=")
		splitPos := strings.Split(split[0], ",")
		splitVel := strings.Split(split[1], ",")
		posX, _ := strconv.Atoi(splitPos[0])
		posY, _ := strconv.Atoi(splitPos[1])
		velX, _ := strconv.Atoi(splitVel[0])
		velY, _ := strconv.Atoi(splitVel[1])
		robots[i] = Robot{pos: vec{float64(posX), float64(posY)}, vel: vec{float64(velX), float64(velY)}}
	}

	seconds := 100
	quadrantCounts := []int{0, 0, 0, 0}

	for iter := 1; iter <= 7892; iter++ {
		bounds := vec{101.0, 103.0}
		for i := range robots {
			robots[i].pos = robots[i].pos.Add(robots[i].vel.Scale(1.0))
			robots[i].pos = vec{math.Mod(robots[i].pos.X(), bounds.X()), math.Mod(robots[i].pos.Y(), bounds.Y())}
			if robots[i].pos.X() < 0.0 {
				robots[i].pos = vec{robots[i].pos.X() + bounds.X(), robots[i].pos.Y()}
			}
			if robots[i].pos.Y() < 0.0 {
				robots[i].pos = vec{robots[i].pos.X(), robots[i].pos.Y() + bounds.Y()}
			}
		}

		// drawPicture(robots, bounds)
		// fmt.Println("ITERATION:", iter)
		// time.Sleep(10 * time.Millisecond)

		if iter == seconds {
			for _, robot := range robots {
				x, y := int(robot.pos.X()), int(robot.pos.Y())
				if x < int(bounds.X()/2.0) && y < int(bounds.Y()/2.0) {
					quadrantCounts[0] += 1
				} else if x > int(bounds.X()/2.0) && y < int(bounds.Y()/2.0) {
					quadrantCounts[1] += 1
				} else if x > int(bounds.X()/2.0) && y > int(bounds.Y()/2.0) {
					quadrantCounts[2] += 1
				} else if x < int(bounds.X()/2.0) && y > int(bounds.Y()/2.0) {
					quadrantCounts[3] += 1
				}
			}
		}
	}

	safetyFactor := quadrantCounts[0] * quadrantCounts[1] * quadrantCounts[2] * quadrantCounts[3]
	fmt.Printf("Safety factor after 100 seconds: %d\n", safetyFactor)
	fmt.Print("Christmas tree spotted at 7892 seconds\n")
}
