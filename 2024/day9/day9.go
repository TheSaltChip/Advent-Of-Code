package day9

import (
	"2024/util"
	"fmt"
	"slices"
)

func Part1() {
	scanner, file := util.OpenFileAndScanner("day9/day9input.txt")

	defer file.Close()

	var fileMap = make(map[int][]int)

	scanner.Scan()
	line := util.StripBOM(scanner.Text())

	numArray, err := util.ToIntArray(line, "")

	if err != nil {
		panic(err)
	}

	fileId := 0

	// When creating fileMap
	for i, r := range numArray {
		if i%2 == 0 {
			if r == 0 {
				continue
			}
			fileMap[i] = slices.Repeat([]int{fileId}, r)
			fileId++
		} else {
			fileMap[i] = slices.Repeat([]int{-1}, r)
		}
	}

	leftIndex := 0
	rightIndex := len(fileMap) - 1

	emptyPlaces := 0
	toPlace := len(fileMap[rightIndex])

	finalList := make([]int, 0, 20000)

	for leftIndex < rightIndex {
		if leftIndex%2 == 0 {
			if len(fileMap[leftIndex]) > 0 {
				finalList = append(finalList, fileMap[leftIndex]...)
			}
			leftIndex++
		}

		if emptyPlaces == 0 {
			emptyPlaces = len(fileMap[leftIndex])
		}

		for emptyPlaces > 0 && toPlace > 0 {
			toPlace--
			emptyPlaces--
			finalList = append(finalList, fileMap[rightIndex][toPlace])
		}

		if toPlace == 0 {
			rightIndex -= 2
			toPlace = len(fileMap[rightIndex])
		}

		if emptyPlaces == 0 {
			leftIndex++
		}
	}

	maxFileId := -1
	minFileId := 100000
	for _, v := range finalList {
		if v > maxFileId {
			maxFileId = v
		}
		if v < minFileId {
			minFileId = v
		}
	}

	sum := 0

	for i, v := range finalList {
		if v >= 0 {
			contribution := i * v
			sum += contribution
		}
	}

	fmt.Println("Sum", sum)
}

func Part2() {
}

// Claude's solution
func solve(input string) int {
	nums, err := util.ToIntArray(input, "")
	if err != nil {
		panic(err)
	}

	// Build initial state
	fileBlocks := make([]int, 0, 50000)
	fileId := 0

	for i := 0; i < len(nums); i += 2 {
		// Add file blocks
		for j := 0; j < nums[i]; j++ {
			fileBlocks = append(fileBlocks, fileId)
		}
		// Add gaps
		if i+1 < len(nums) {
			for j := 0; j < nums[i+1]; j++ {
				fileBlocks = append(fileBlocks, -1)
			}
		}
		fileId++
	}

	// Start moving files right to left
	finalBlocks := make([]int, 0, len(fileBlocks))
	gapStart := 0
	fileEnd := len(fileBlocks) - 1

	for gapStart < fileEnd {
		// Find next gap
		for gapStart < len(fileBlocks) && fileBlocks[gapStart] != -1 {
			finalBlocks = append(finalBlocks, fileBlocks[gapStart])
			gapStart++
		}

		// Find rightmost file block
		for fileEnd >= 0 && fileBlocks[fileEnd] == -1 {
			fileEnd--
		}

		// Move one block if we found both gap and file
		if gapStart < fileEnd && fileBlocks[gapStart] == -1 && fileBlocks[fileEnd] != -1 {
			finalBlocks = append(finalBlocks, fileBlocks[fileEnd])
			fileEnd--
			gapStart++
		}

		//fmt.Println("Current state:", finalBlocks)
	}

	// Calculate checksum
	sum := 0
	for pos, fileId := range finalBlocks {
		if fileId != -1 {
			sum += pos * fileId
		}
	}

	return sum
}
