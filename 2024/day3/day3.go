package day3

import (
	"2024/util"
	"bufio"
	"github.com/adam-lavrik/go-imath/ix"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
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

func Part2() {
	file, err := os.Open("day3/day3input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	do := true

	for scanner.Scan() {
		line := scanner.Text()

		regexMul, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
		regexDo, err := regexp.Compile(`do\(\)`)
		regexDont, err := regexp.Compile(`don't\(\)`)

		if err != nil {
			log.Fatal(err)
		}

		resultIndecies := regexMul.FindAllStringIndex(line, -1)
		doIndecies := regexDo.FindAllStringIndex(line, -1)
		dontIndecies := regexDont.FindAllStringIndex(line, -1)

		combinedMap := make(map[int][]string)

		maxI := ix.Maxs(len(resultIndecies), len(doIndecies), len(dontIndecies))

		for i := 0; i < maxI; i++ {
			if i < len(resultIndecies) {
				res := resultIndecies[i]
				combinedMap[res[0]] = append(combinedMap[res[0]], line[res[0]:res[1]])
			}
			if i < len(doIndecies) {
				res := doIndecies[i]
				combinedMap[res[0]] = append(combinedMap[res[0]], line[res[0]:res[1]])
			}
			if i < len(dontIndecies) {
				res := dontIndecies[i]
				combinedMap[res[0]] = append(combinedMap[res[0]], line[res[0]:res[1]])
			}
		}

		keys := make([]int, 0, len(combinedMap))
		for k := range combinedMap {
			keys = append(keys, k)
		}

		sort.Ints(keys)

		for _, key := range keys {
			if do && regexMul.MatchString(combinedMap[key][0]) {
				numStrs := line[key+4 : key+4+strings.Index(line[key+4:], ")")]
				nums, err := util.ToIntArray(numStrs, ",")

				if err != nil {
					log.Fatal(err)
				}

				sum += nums[0] * nums[1]
			} else if regexDo.MatchString(combinedMap[key][0]) {
				do = true
			} else if regexDont.MatchString(combinedMap[key][0]) {
				do = false
			}
		}
	}

	log.Println(sum)
}
