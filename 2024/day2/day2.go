package day2

import (
	"2024/util"
	"bufio"
	"fmt"
	"github.com/adam-lavrik/go-imath/ix"
	"os"
)

func readFile() *os.File {
	file, err := os.Open("day2/day2input.txt")

	if err != nil {
		panic(err)
	}

	return file
}

func Part1() {
	file := readFile()

	numSafeRapport := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := util.StripBOM(scanner.Text())
		ints, err := util.ToIntArray(line, " ")

		if err != nil {
			panic(err)
		}

		if isSafe(ints) {
			numSafeRapport++
		}
	}

	fmt.Println(numSafeRapport)
}

func isSafe(ints []int) bool {
	var increasing bool
	for i := range ints {
		curr := ints[i]
		next := ints[i+1]
		diff := ix.Abs(next - curr)
		if i == 0 {
			increasing = curr < next
		}

		if diff < 1 || diff > 3 ||
			(increasing && curr > next) ||
			(!increasing && curr < next) {
			return false
		}

		if i == len(ints)-2 {
			return true
		}
	}

	return false
}

func Part2() {
	file := readFile()

	numSafeRapport := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := util.StripBOM(scanner.Text())
		ints, err := util.ToIntArray(line, " ")

		if err != nil {
			panic(err)
		}

		if isSafe(ints) {
			numSafeRapport++
		} else {
			for deleteIndex := 0; deleteIndex < len(ints); deleteIndex++ {
				copyInts := make([]int, len(ints))
				copy(copyInts, ints)

				if deleteIndex == len(copyInts)-1 {
					copyInts = copyInts[:deleteIndex]
				} else {
					copyInts = append(copyInts[:deleteIndex], copyInts[deleteIndex+1:]...)
				}

				if isSafe(copyInts) {
					numSafeRapport++
					break
				}
			}
		}
	}

	fmt.Println(numSafeRapport)
}

func oldBrokenPart2() {
	file := readFile()

	numSafeRapport := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := util.StripBOM(scanner.Text())
		ints, err := util.ToIntArray(line, " ")

		if err != nil {
			panic(err)
		}

		var increasing bool
		levelsToRemove := 1
		for i := 0; i < len(ints)-1; i++ {
			curr := ints[i]
			next := ints[i+1]
			diff := ix.Abs(next - curr)
			levelRemoved := false

			if i == 0 {
				increasing = curr < next
			}

			if diff < 1 || diff > 3 {
				if levelsToRemove == 0 {
					break
				}

				if i != 0 && i != len(ints)-2 {
					newDiff := ix.Abs(ints[i-1] - ints[i+1])
					if newDiff < 1 || newDiff > 3 {
						break
					}
				}

				levelsToRemove--
				levelRemoved = true
			}

			if levelRemoved {
				if i == 0 {
					increasing = next < ints[i+2]
				} else if i == 1 {
					increasing = ints[i-1] < next
				}
			}

			if ((increasing && curr > next) ||
				(!increasing && curr < next)) && !levelRemoved {

				if levelsToRemove == 0 {
					break
				}

				if i != 0 && i != len(ints)-2 {
					curr = ints[i-1]
					next = ints[i+1]

					increasing = curr < next

					if (increasing && curr > next) ||
						(!increasing && curr < next) {
						break
					}
				}

				if i == 0 {
					curr = next
					next = ints[i+2]
					increasing = curr < next
				}

				levelsToRemove--
				levelRemoved = true
			}

			if i == len(ints)-2 {
				numSafeRapport++
				break
			}
		}
	}

	fmt.Println(numSafeRapport)
}
