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

func is_guard_outside_map(x, y int, input [][]byte) bool {
	return x < 0 || x >= len(input) || y < 0 || y >= len(input[x])
}

func turn_right(dir_x *int, dir_y *int) {
	cos_90, sin_90 := 0, -1
	*dir_x, *dir_y = *dir_x*cos_90-(-*dir_y)*sin_90, -(*dir_x*sin_90 + (-*dir_y)*cos_90)
}

func check_and_turn_right(x int, y int, dir_x *int, dir_y *int, input [][]byte) (has_turned bool) {
	new_guard_x, new_guard_y := x+*dir_x, y+*dir_y
	if !is_guard_outside_map(new_guard_x, new_guard_y, input) && input[new_guard_y][new_guard_x] == '#' {
		turn_right(dir_x, dir_y)

		new_guard_x, new_guard_y = x+*dir_x, y+*dir_y
		if !is_guard_outside_map(new_guard_x, new_guard_y, input) && input[new_guard_y][new_guard_x] == '#' {
			turn_right(dir_x, dir_y)
		}

		return true
	}
	return false
}

func move(x *int, y *int, dir_x int, dir_y int, input [][]byte) {
	input[*y][*x] = 'X'
	*x += dir_x
	*y += dir_y
	//
	// DRAW MOVEMENTS ON THE MAP
	//
	// time.Sleep(10 * time.Millisecond)
	// func() {
	// 	defer func() {
	// 		if r := recover(); r != nil {
	// 		}
	// 	}()
	// fmt.Printf("-----\n|%s|\n|%s|\n|%s|\n|%s|\n|%s|\n",
	// 	input[*y-2][*x-4:*x+5], input[*y-1][*x-4:*x+5], input[*y][*x-4:*x+5], input[*y+1][*x-4:*x+5], input[*y+2][*x-4:*x+5])
	// }()
}

func move2(x *int, y *int, dir_x int, dir_y int, input [][]byte, has_turned bool) (has_cycle bool) {
	if has_turned && input[*y][*x] == '+' {
		if input[*y+dir_y][*x+dir_x] == '|' && dir_y != 0 || input[*y+dir_y][*x+dir_x] == '-' && dir_x != 0 {
			return true
		}
	}

	if has_turned {
		input[*y][*x] = '+'
	} else if dir_x == 0 {
		input[*y][*x] = '|'
	} else if dir_y == 0 {
		input[*y][*x] = '-'
	}
	*x += dir_x
	*y += dir_y

	return false
}

func main() {
	filePath := filepath.Join(Must(os.Getwd()), "input_map.txt")
	input := bytes.Split(Must((os.ReadFile(filePath))), []byte("\r\n"))
	m := make([][]byte, len(input))
	for i := range input {
		m[i] = make([]byte, len(input[i]))
		copy(m[i], input[i])
	}

	var initial_guard_x, initial_guard_y int
	for y := range m {
		x := bytes.IndexByte(m[y], '^')
		if x != -1 {
			initial_guard_x, initial_guard_y = x, y
			break
		}
	}

	// challenge 1

	guard_x, guard_y := initial_guard_x, initial_guard_y
	guard_dir_x, guard_dir_y := 0, -1
	for !is_guard_outside_map(guard_x, guard_y, m) {
		check_and_turn_right(guard_x, guard_y, &guard_dir_x, &guard_dir_y, m)
		move(&guard_x, &guard_y, guard_dir_x, guard_dir_y, m)
	}

	num_of_Xs := 0
	for _, line := range m {
		num_of_Xs += bytes.Count(line, []byte{'X'})
	}
	fmt.Printf("Number of unique positions visited: %d\n", num_of_Xs)

	// challenge 2

	num_of_cycles := 0
	for obstruction_y := range input {
		for obstruction_x := range input[obstruction_y] {
			if input[obstruction_y][obstruction_x] != '.' {
				continue
			}

			m2 := make([][]byte, len(input))
			for i := range input {
				m2[i] = make([]byte, len(input[i]))
				copy(m2[i], input[i])
			}
			m2[obstruction_y][obstruction_x] = '#'
			// m2[6][3] = '#'

			guard_x, guard_y := initial_guard_x, initial_guard_y
			guard_dir_x, guard_dir_y = 0, -1
			for !is_guard_outside_map(guard_x, guard_y, m2) {
				has_turned := check_and_turn_right(guard_x, guard_y, &guard_dir_x, &guard_dir_y, m2)
				is_cycle := move2(&guard_x, &guard_y, guard_dir_x, guard_dir_y, m2, has_turned)

				if is_cycle {
					num_of_cycles++
					break
				}
			}
		}
	}
	fmt.Printf("Number of unique positions obstructed to create a cycle: %d\n", num_of_cycles)
}
