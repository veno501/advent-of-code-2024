package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type key = []byte
type lock = []byte

func main() {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "input_locks.txt")
	inputBytes, _ := os.ReadFile(filePath)
	input := strings.Split(string(inputBytes), "\r\n\r\n")

	var keys []key
	var locks []lock
	for _, blockStr := range input {
		block := strings.Split(blockStr, "\r\n")
		blockSize := len(block[0])
		cols := make([]byte, blockSize)
		for x := 0; x < blockSize; x++ {
			for y := 0; y < blockSize; y++ {
				if block[y+1][x] == '#' {
					cols[x] |= 1 << y
				}
			}
		}
		if block[0][0] == '#' {
			locks = append(locks, cols)
		} else {
			keys = append(keys, cols)
		}
	}

	var count int
	for _, key := range keys {
		for _, lock := range locks {
			matches := true
			for x := 0; x < len(key); x++ {
				if x == 2 {
					fmt.Printf("%08b\n", lock[x])
					fmt.Printf("%08b\n", key[x])
				}
				if key[x]&lock[x] != 0 {
					matches = false
					break
				}
			}
			if matches {
				count += 1
			}
		}
	}
	fmt.Printf("Number of lock/key pairs that fit together: %d\n", count)
}
