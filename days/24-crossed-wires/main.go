package main

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type Gate struct {
	in1, in2  string
	operation string
	out       string
}

func isXOrY(wire string) bool {
	significance, _ := strconv.Atoi(wire[1:])
	return (wire[0] == 'x' || wire[0] == 'y') && significance != 0
}

func main() {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "input_gates.txt")
	inputBytes, _ := os.ReadFile(filePath)
	input := strings.Split(string(inputBytes), "\r\n\r\n")

	wires := make(map[string]int)
	for _, str := range strings.Split(input[0], "\r\n") {
		split := strings.Split(str, ": ")
		wires[split[0]], _ = strconv.Atoi(split[1])
	}
	precursorGates := make(map[string]Gate)
	for _, str := range strings.Split(input[1], "\r\n") {
		split := strings.Split(str, " ")
		gate := Gate{in1: split[0], in2: split[2], operation: split[1], out: split[4]}
		precursorGates[gate.out] = gate
	}

	var calcGateOutput func(gate Gate) int
	calcGateOutput = func(gate Gate) int {
		inVal1, found := wires[gate.in1]
		if !found {
			inVal1 = calcGateOutput(precursorGates[gate.in1])
		}
		inVal2, found := wires[gate.in2]
		if !found {
			inVal2 = calcGateOutput(precursorGates[gate.in2])
		}
		outVal := inVal1
		if gate.operation == "AND" {
			outVal &= inVal2
		} else if gate.operation == "OR" {
			outVal |= inVal2
		} else if gate.operation == "XOR" {
			outVal ^= inVal2
		}
		wires[gate.out] = outVal
		return outVal
	}

	// challenge 1

	for _, gate := range precursorGates {
		calcGateOutput(gate)
	}
	var zWireVals int64
	for wireId, wireVal := range wires {
		if wireId[0] == 'z' {
			significance, _ := strconv.Atoi(wireId[1:])
			zWireVals += int64(wireVal) << significance
		}
	}
	fmt.Printf("z-wire values: %d\n", zWireVals)

	// challenge 2

	swappedWires := make(map[string]bool)
	for wireName, precursorGate := range precursorGates {
		if wireName[0] == 'z' {
			val, _ := strconv.Atoi(wireName[1:])
			if precursorGate.operation != "XOR" && val != 45 {
				swappedWires[wireName] = true
			}
		} else if !isXOrY(precursorGate.in1) && !isXOrY(precursorGate.in2) && precursorGate.in1[0] != precursorGate.in2[0] && precursorGate.operation == "XOR" {
			swappedWires[wireName] = true
		}

		if precursorGate.operation == "XOR" && isXOrY(precursorGate.in1) && isXOrY(precursorGate.in2) && precursorGate.in1[0] != precursorGate.in2[0] {
			isValid := false
			for _, dp := range precursorGates {
				if dp.operation == "XOR" && (dp.in1 == wireName || dp.in2 == wireName) {
					isValid = true
				}
			}
			if !isValid {
				swappedWires[wireName] = true
			}
		}

		if precursorGate.operation == "AND" && isXOrY(precursorGate.in1) && isXOrY(precursorGate.in2) && precursorGate.in1[0] != precursorGate.in2[0] {
			isValid := false
			for _, dp := range precursorGates {
				if dp.operation == "OR" && (dp.in1 == wireName || dp.in2 == wireName) {
					isValid = true
				}
			}
			if !isValid {
				swappedWires[wireName] = true
			}
		}
	}
	keys := slices.Collect(maps.Keys(swappedWires))
	slices.Sort(keys)
	fmt.Printf("Wires that were swapped: %s\n", strings.Join(keys, ","))
}
