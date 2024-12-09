package day7

import (
	"2024/util"
	"log"
	"strconv"
	"strings"
)

func Part1() {
	scanner, file := util.OpenFileAndScanner("day7/day7.txt")

	defer file.Close()

	sumAndValues := make(map[int][]int)

	for scanner.Scan() {
		line := util.StripBOM(scanner.Text())

		temp := strings.Split(line, ":")

		sum, err := strconv.Atoi(temp[0])

		if err != nil {
			log.Fatal(err)
		}

		values, err := util.ToIntArray(temp[1], " ")

		if err != nil {
			log.Fatal(err)
		}

		sumAndValues[sum] = values
	}

}

func Part2() {}
