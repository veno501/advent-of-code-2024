package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func calculateCost(ax, bx, ay, by, px, py int) int {
	A := mat.NewDense(2, 2, []float64{float64(ax), float64(bx), float64(ay), float64(by)})
	Y := mat.NewDense(2, 1, []float64{float64(px), float64(py)})
	A.Inverse(A)
	var X mat.Dense
	X.Mul(A, Y)

	ta := X.At(0, 0)
	tb := X.At(1, 0)

	if math.Abs(math.Round(ta)-ta) < 0.001 {
		ta = math.Round(ta)
	} else {
		return 0
	}
	if math.Abs(math.Round(tb)-tb) < 0.001 {
		tb = math.Round(tb)
	} else {
		return 0
	}

	return int(ta)*3 + int(tb)*1
}

func main() {
	filePath := filepath.Join(Must(os.Getwd()), "input_machines.txt")
	input := strings.Split(string(Must((os.ReadFile(filePath)))), "\r\n\r\n")

	var costSum, costSumBig int
	for _, machine := range input {
		split := strings.Split(machine, "\r\n")
		a := strings.Split(split[0], "X+")[1]
		ax, ay := Must(strconv.Atoi(strings.Split(a, ", Y+")[0])), Must(strconv.Atoi(strings.Split(a, ", Y+")[1]))
		b := strings.Split(split[1], "X+")[1]
		bx, by := Must(strconv.Atoi(strings.Split(b, ", Y+")[0])), Must(strconv.Atoi(strings.Split(b, ", Y+")[1]))
		prize := strings.Split(split[2], "X=")[1]
		px, py := Must(strconv.Atoi(strings.Split(prize, ", Y=")[0])), Must(strconv.Atoi(strings.Split(prize, ", Y=")[1]))

		costSum += calculateCost(ax, bx, ay, by, px, py)
		costSumBig += calculateCost(ax, bx, ay, by, px+10000000000000, py+10000000000000)
	}
	fmt.Printf("Total cost: %d\n", costSum)
	fmt.Printf("Total cost (big): %d\n", costSumBig)
}
