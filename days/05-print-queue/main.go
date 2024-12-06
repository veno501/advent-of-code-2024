package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func main() {
	filePath := filepath.Join(Must(os.Getwd()), "input_pages.txt")
	input := strings.Split(string(Must((os.ReadFile(filePath)))), "\r\n\r\n")
	rules := strings.Split(input[0], "\r\n")
	updates := strings.Split(input[1], "\r\n")

	var correctlyOrderedUpdatesResult, incorrectlyOrderedUpdatesResult int
	for _, update := range updates {
		pages := strings.Split(update, ",")

		sortedPages := slices.SortedFunc(slices.Values(pages), func(a, b string) int {
			var numOfPagesAfterA, numOfPagesAfterB int
			for _, rule := range rules {
				split := strings.Split(string(rule), "|")
				before, after := split[0], split[1]

				if slices.Contains(pages, after) {
					if before == a {
						numOfPagesAfterA++
					}
					if before == b {
						numOfPagesAfterB++
					}
				}
			}
			return numOfPagesAfterB - numOfPagesAfterA
		})

		if slices.Equal(sortedPages, pages) {
			correctlyOrderedUpdatesResult += Must(strconv.Atoi(sortedPages[len(sortedPages)/2]))
		} else {
			incorrectlyOrderedUpdatesResult += Must(strconv.Atoi(sortedPages[len(sortedPages)/2]))
		}
	}

	fmt.Printf("Summed middle pages of correctly-ordered updates: %d\n", correctlyOrderedUpdatesResult)
	fmt.Printf("Summed middle pages of incorrectly-ordered updates after ordering: %d\n", incorrectlyOrderedUpdatesResult)
}
