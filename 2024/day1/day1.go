package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func stripBOM(str string) string {
	return strings.TrimPrefix(str, "\ufeff")
}

func CreateLists() ([]int, []int, error) {
	dat, err := os.Open("day1/input1.txt")

	if err != nil {
		return nil, nil, err
	}

	defer func(dat *os.File) {
		err := dat.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(dat)

	scanner := bufio.NewScanner(dat)

	var leftSideSplice []int
	var rightSideSplice []int

	for scanner.Scan() {
		line := scanner.Text()
		elems := strings.Split(stripBOM(line), "   ")
		num1, err1 := strconv.Atoi(elems[0])

		if err1 != nil {
			return nil, nil, err
		}

		num2, err2 := strconv.Atoi(elems[1])

		if err2 != nil {
			return nil, nil, err
		}
		leftSideSplice = append(leftSideSplice, num1)
		rightSideSplice = append(rightSideSplice, num2)
	}

	return leftSideSplice, rightSideSplice, err
}
func Part1() {
	leftSideSplice, rightSideSplice, err := CreateLists()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	slices.Sort(leftSideSplice)
	slices.Sort(rightSideSplice)

	sum := 0
	for i := 0; i < len(leftSideSplice); i++ {
		sum += int(math.Abs(float64(leftSideSplice[i] - rightSideSplice[i])))
	}

	fmt.Println(sum)
}

func Part2() {
	leftSideSplice, rightSideSplice, err := CreateLists()

	if err != nil {
		panic(err)
	}

	freqMap := make(map[int]int)

	for _, val := range rightSideSplice {
		_, exists := freqMap[val]

		if exists {
			freqMap[val]++
		} else {
			freqMap[val] = 1
		}
	}

	sum := 0

	for _, val := range leftSideSplice {
		sum += freqMap[val] * val
	}

	fmt.Println(sum)
}
