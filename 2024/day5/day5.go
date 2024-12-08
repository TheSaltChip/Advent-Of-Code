package day5

import (
	"2024/util"
	"bufio"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func createRules(scanner *bufio.Scanner) map[string][]string {
	rules := make(map[string][]string)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		nums := strings.Split(scanner.Text(), "|")

		rules[nums[1]] = append(rules[nums[1]], nums[0])
	}
	return rules
}

func Part1() {
	scanner, file := util.OpenFileAndScanner("day5/day5input.txt")

	defer file.Close()

	rules := createRules(scanner)

	var validUpdates [][]string

	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), ",")
		var alreadyPrinted []string
		valid := true

	outer:
		for _, num := range nums {
			if rules[num] == nil {
				alreadyPrinted = append(alreadyPrinted, num)
				continue
			}

			deps := rules[num]

			// Iterate over deps
			for _, dep := range deps {
				// is dep in update
				if slices.Contains(nums, dep) {
					// is the dep already printed
					if !slices.Contains(alreadyPrinted, dep) {
						valid = false
						break outer
					}
				}
			}

			alreadyPrinted = append(alreadyPrinted, num)
			continue
		}

		if valid {
			validUpdates = append(validUpdates, alreadyPrinted)
		}
	}

	sum := 0

	for _, update := range validUpdates {
		val, err := strconv.Atoi(update[len(update)/2])

		if err != nil {
			log.Fatal(err)
		}

		sum += val
	}

	fmt.Println(sum)
}

func Part2() {
	scanner, file := util.OpenFileAndScanner("day5/day5input.txt")

	defer file.Close()

	rules := createRules(scanner)

	var inValidUpdates [][]string

	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), ",")
		var alreadyPrinted []string
		valid := true

	outer:
		for _, num := range nums {
			if rules[num] == nil {
				alreadyPrinted = append(alreadyPrinted, num)
				continue
			}

			deps := rules[num]

			// Iterate over deps
			for _, dep := range deps {
				// is dep in update
				if slices.Contains(nums, dep) {
					// is the dep already printed
					if !slices.Contains(alreadyPrinted, dep) {
						valid = false
						break outer
					}
				}
			}

			alreadyPrinted = append(alreadyPrinted, num)
			continue
		}

		if !valid {
			inValidUpdates = append(inValidUpdates, nums)
		}
	}

	sum := 0

	for _, update := range inValidUpdates {
		for i, item := range update {

			deps := rules[item]

			if deps == nil {
				continue
			}
			ind := i
			for j := i + 1; j < len(update); j++ {
				if slices.Contains(deps, update[j]) {
					update[ind], update[j] = update[j], update[ind]
					j = ind + 1
					deps = rules[update[ind]]
					// [1 2 3 4]
					//  0 1 2 3 : ind 0, j = 3
					// [4 2 3 1] : ind 0, j = 0
				}
			}
		}
	}

	for _, update := range inValidUpdates {
		val, err := strconv.Atoi(update[len(update)/2])

		if err != nil {
			log.Fatal(err)
		}

		sum += val
	}

	fmt.Println(sum)
}
