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

type Permutation = []rune

func evaluateExpression(terms []int, operators Permutation) int {
	t := slices.Clone(terms)
	result := t[0]
	for j := range operators {
		switch operators[j] {
		case '+':
			result += t[j+1]
			break
		case '*':
			result *= t[j+1]
			break
		case '|':
			result = Must(strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(t[j+1])))
		}
	}
	return result
}

func generatePermutation(set_length int, tokens []rune) []Permutation {
	if set_length == 0 {
		return []Permutation{{}}
	}
	permutations := make([]Permutation, 0, set_length*set_length)

	sub_permutations := generatePermutation(set_length-1, tokens)
	for _, token := range tokens {
		for _, sub_permutation := range sub_permutations {
			permutations = append(permutations, append([]rune{token}, sub_permutation...))
		}
	}
	return permutations
}

func main() {
	filePath := filepath.Join(Must(os.Getwd()), "input_equations.txt")
	input := strings.Split(string(Must((os.ReadFile(filePath)))), "\r\n")

	summed_test_values_with_valid_permutations := 0
	summed_test_values_with_valid_permutations_with_concatenation := 0
	for _, line := range input {
		lineSplit := strings.Split(line, ": ")
		testValue := Must(strconv.Atoi(lineSplit[0]))

		var terms []int
		for _, s := range strings.Split(lineSplit[1], " ") {
			terms = append(terms, Must(strconv.Atoi(s)))
		}

		// challenge 1

		for _, permutation := range generatePermutation(len(terms)-1, []rune{'+', '*'}) {
			if evaluateExpression(terms, permutation) == testValue {

				// fmt.Print(testValue, ": ", terms[0])
				// for i := range permutation {
				// 	fmt.Print(" ", string(permutation[i]), " ", terms[i+1])
				// }
				// fmt.Println()

				// for _, p := range generatePermutation(len(terms) - 1) {
				// 	fmt.Print(testValue, ": ", terms[0])
				// 	for i := range p {
				// 		fmt.Print(" ", string(p[i]), " ", terms[i+1])
				// 	}
				// 	fmt.Println("(", len(generatePermutation(len(terms)-1)), ")")
				// }

				summed_test_values_with_valid_permutations += testValue
				break
			}
		}

		// challenge 2

		for _, permutation := range generatePermutation(len(terms)-1, []rune{'+', '*', '|'}) {
			if evaluateExpression(terms, permutation) == testValue {

				// fmt.Print(testValue, ": ", terms[0])
				// for i := range permutation {
				// 	fmt.Print(" ", string(permutation[i]), " ", terms[i+1])
				// }
				// fmt.Println()

				// for _, p := range generatePermutation(len(terms) - 1) {
				// 	fmt.Print(testValue, ": ", terms[0])
				// 	for i := range p {
				// 		fmt.Print(" ", string(p[i]), " ", terms[i+1])
				// 	}
				// 	fmt.Println("(", len(generatePermutation(len(terms)-1)), ")")
				// }

				summed_test_values_with_valid_permutations_with_concatenation += testValue
				break
			}
		}
	}
	fmt.Printf("Sum of test values of equations with valid {+, *} operator permutations: %d\n", summed_test_values_with_valid_permutations)
	fmt.Printf("Sum with valid {+, *, || (concatenation)} operator permutations: %d\n", summed_test_values_with_valid_permutations_with_concatenation)
}
