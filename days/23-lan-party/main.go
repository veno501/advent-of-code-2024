package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func createPowerSet(set []string) [][]string {
	powerSet := [][]string{}
	for i := 0; i < (1 << len(set)); i++ {
		subset := []string{}
		for j := 0; j < len(set); j++ {
			if i&(1<<j) > 0 {
				subset = append(subset, set[j])
			}
		}
		powerSet = append(powerSet, subset)
	}
	return powerSet
}

func main() {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "input_network.txt")
	inputBytes, _ := os.ReadFile(filePath)
	input := strings.Split(string(inputBytes), "\r\n")

	neighbors := make(map[string][]string)
	for _, line := range input {
		split := strings.Split(line, "-")
		PC1, PC2 := split[0], split[1]
		neighbors[PC1] = append(neighbors[PC1], PC2)
		neighbors[PC2] = append(neighbors[PC2], PC1)
	}

	// challenge 1

	cliquesSet := make(map[string]bool)
	for node := range neighbors {
		for _, adjacentNode := range neighbors[node] {
			for _, adjacentNodeToAdjacentNode := range neighbors[adjacentNode] {
				if slices.Contains(neighbors[node], adjacentNodeToAdjacentNode) {
					nodes := []string{node, adjacentNode, adjacentNodeToAdjacentNode}
					slices.Sort(nodes)
					cliquesSet[strings.Join(nodes, ",")] = true
				}
			}
		}
	}
	var countCliques int
	for key := range cliquesSet {
		nodes := strings.Split(key, ",")
		if strings.HasPrefix(nodes[0], "t") || strings.HasPrefix(nodes[1], "t") || strings.HasPrefix(nodes[2], "t") {
			countCliques += 1
		}
	}
	fmt.Printf("Number of cliques of three interconnected computers: %d\n", countCliques)

	// challenge 2

	fans := []string{}
	for node := range neighbors {
		for _, subset := range createPowerSet(neighbors[node]) {
			fan := append(subset, node)
			slices.Sort(fan)
			fans = append(fans, strings.Join(fan, ","))
		}
	}

	fanCounts := make(map[string]int)
	for _, fan := range fans {
		fanCounts[fan] += 1
	}

	maxClique := ""
	maxCliqueLength := 0
	for fan, count := range fanCounts {
		nodes := strings.Split(fan, ",")
		if len(nodes) == count && len(nodes) > maxCliqueLength {
			maxCliqueLength = len(nodes)
			maxClique = fan
		}
	}
	fmt.Printf("Largest clique of interconnected computers in the network: %s\n", maxClique)
}
