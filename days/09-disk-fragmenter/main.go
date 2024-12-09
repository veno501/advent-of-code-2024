package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func main() {
	filePath := filepath.Join(Must(os.Getwd()), "input_disk.txt")
	input := string(Must((os.ReadFile(filePath))))

	var disk []int
	for i, length, id := 0, len(input), 0; i < length; i += 2 {
		fileSpaces := Must(strconv.Atoi(string(input[i])))
		disk = slices.Concat(disk, slices.Repeat([]int{id}, fileSpaces))
		id++
		if i+1 >= length {
			break
		}
		freeSpaces := Must(strconv.Atoi(string(input[i+1])))
		disk = slices.Concat(disk, slices.Repeat([]int{-1}, freeSpaces))
	}
	disk2 := slices.Clone(disk)

	// challenge 1

	for i1, i2 := 0, len(disk)-1; i1 < i2; i1++ {
		if disk[i1] == -1 {
			disk[i1] = disk[i2]
			disk[i2] = -1
			for disk[i2] == -1 {
				i2--
			}
		}
	}

	// challenge 2

	for i2 := len(disk2) - 1; 0 < i2; {
		n2 := 0
		for i := i2; i >= 0 && disk2[i] == disk2[i2]; i-- {
			n2++
		}

		i1 := -1
		for i, counter := 0, 0; i < i2; i++ {
			if disk2[i] == -1 {
				counter++
			} else {
				counter = 0
			}
			if counter == n2 {
				i1 = i - (counter - 1)
				break
			}
		}

		if i1 != -1 {
			for i := 0; i < n2; i++ {
				disk2[i1+i] = disk2[i2-i]
				disk2[i2-i] = -1
			}
		}

		i2 -= n2
		if i2 < 0 {
			break
		}
		for disk2[i2] == -1 {
			i2--
		}
	}

	checksum := 0
	for i := range disk {
		if disk[i] != -1 {
			checksum += disk[i] * i
		}
	}
	fmt.Println(checksum)

	checksum2 := 0
	for i := range disk2 {
		if disk2[i] != -1 {
			checksum2 += disk2[i] * i
		}
	}
	fmt.Println(checksum2)
}
