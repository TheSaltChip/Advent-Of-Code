﻿package day3

import (
	"2024/util"
	"bufio"
	"log"
	"os"
	"regexp"
)

func Part1() {
	file, err := os.Open("day3/day3input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		regex, err := regexp.Compile(`mul\(\d+,\d+\)`)

		if err != nil {
			log.Fatal(err)
		}

		results := regex.FindAllStringIndex(line, -1)

		for _, res := range results {
			numStrs := line[res[0]+4 : res[1]-1]
			nums, err := util.ToIntArray(numStrs, ",")

			if err != nil {
				log.Fatal(err)
			}

			sum += nums[0] * nums[1]
		}
	}

	log.Println(sum)
}

/*
for {
				if dontI == len(dontIndecies) && doI == len(doIndecies) {

				}
				log.Println(res, dontIndecies[dontI], doIndecies[doI])
				if res[0] > dontIndecies[dontI][0] && res[0] > doIndecies[doI][0] {
					if doI != len(doIndecies) && dontI != len(dontIndecies) {
						doI++
						dontI++
						continue
					} else if doI == len(doIndecies) && dontI != len(dontIndecies) {
						if dontIndecies[dontI][0] > doIndecies[doI][0] {
							do = false
							break
						} else {
							dontI++
						}
					} else if doI != len(doIndecies) && dontI == len(dontIndecies) {
						if doIndecies[doI][0] > dontIndecies[dontI][0] {
							do = true
							break
						} else {
							doI++
						}
					} else {
						if doIndecies[doI][0] > dontIndecies[dontI][0] {
							dontI++
						} else {
							doI++
						}
					}
				} else if (res[0] > doIndecies[doI][0]) && (res[0] < dontIndecies[dontI][0]) {
					do = true
					break
				} else if (res[0] < doIndecies[doI][0]) && (res[0] > dontIndecies[dontI][0]) {
					do = false
					continue outer
				} else {
					break
				}
			}
*/

func Part2() {
	file, err := os.Open("day3/day3input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := util.StripBOM(scanner.Text())

		regexMul, err := regexp.Compile(`mul\(\d+,\d+\)`)
		regexDo, err := regexp.Compile(`do\(\)`)
		regexDont, err := regexp.Compile(`don't\(\)`)

		if err != nil {
			log.Fatal(err)
		}

		resultIndecies := regexMul.FindAllStringIndex(line, -1)
		doIndecies := regexDo.FindAllStringIndex(line, -1)
		dontIndecies := regexDont.FindAllStringIndex(line, -1)

		doI, dontI := 0, 0

		do := true

		for _, res := range resultIndecies {

			for _, dontIndecy := range dontIndecies {
				if res[0] > dontIndecy[0] {
					dontI = dontIndecy[0]
				}
			}

			for _, doIndecy := range doIndecies {
				if res[0] > doIndecy[0] {
					doI = doIndecy[0]
				}
			}

			do = doI >= dontI

			if !do {
				continue
			}

			numStrs := line[res[0]+4 : res[1]-1]
			nums, err := util.ToIntArray(numStrs, ",")
			log.Printf("%d < %d > %d | %s \n", doI, res[0], dontI, numStrs)

			if err != nil {
				log.Fatal(err)
			}

			sum += nums[0] * nums[1]
		}
	}

	log.Println(sum)
}