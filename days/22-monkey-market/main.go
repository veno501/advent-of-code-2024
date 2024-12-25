package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func calcNextSecret(secret, iterations int) (int, []int, []int) {
	prices, changes := make([]int, iterations), make([]int, iterations)
	for i := 0; i < iterations; i++ {
		prevSecret := secret
		secret = prune(mix(secret, secret<<6))
		secret = prune(mix(secret, secret>>5))
		secret = prune(mix(secret, secret<<11))
		prices[i] = secret % 10
		changes[i] = secret%10 - prevSecret%10
	}
	return secret, prices, changes
}

func mix(secret, n int) int {
	return secret ^ n
}

func prune(secret int) int {
	return secret % 16777216
}

func sliceToString(slice []int) string {
	// var strs []string
	// for _, n := range slice {
	// 	strs[strconv.Itoa(n)]
	// }
	// return strings.Join(strs, ",")
	var str string
	for _, n := range slice {
		str += strconv.Itoa(n)
	}
	return str
}

func main() {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "input_secrets.txt")
	inputBytes, _ := os.ReadFile(filePath)
	input := strings.Split(string(inputBytes), "\r\n")

	pricesForPatterns := make(map[string][]int)

	var sumLastSecrets int
	for _, secretStr := range input {
		secret, _ := strconv.Atoi(secretStr)
		secret, prices, changes := calcNextSecret(secret, 2000)
		sumLastSecrets += secret

		visited := make(map[string]bool)
		for i := 3; i < len(changes); i++ {
			pattern := sliceToString(changes[i-3 : i+1])
			if _, found := visited[pattern]; !found {
				if _, ok := pricesForPatterns[pattern]; !ok {
					pricesForPatterns[pattern] = []int{}
				}
				pricesForPatterns[pattern] = append(pricesForPatterns[pattern], prices[i])
				visited[pattern] = true
			}
		}
	}
	maxPriceSum := 0
	for _, prices := range pricesForPatterns {
		priceSum := 0
		for _, price := range prices {
			priceSum += price
		}
		if priceSum > maxPriceSum {
			maxPriceSum = priceSum
		}
	}
	fmt.Printf("Sum of the 2000th secret number for each buyer: %d\n", sumLastSecrets)
	fmt.Printf("Maximum sum of the 2000th secret number for each buyer: %d\n", maxPriceSum)
}
