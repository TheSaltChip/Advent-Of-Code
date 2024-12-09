package day7

import (
	"2024/util"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Part1() {
	scanner, file := util.OpenFileAndScanner("day7/day7input.txt")

	defer file.Close()

	var validSums []int

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

		var usedOps []string

		usedOps, err = Check(sum, values[0], values, 1, usedOps)

		if err == nil {
			//fmt.Printf("%d %s\n", values, usedOps)
			validSums = append(validSums, sum)
		}
	}

	sum := 0

	for _, validSum := range validSums {
		sum += validSum
	}

	fmt.Println(sum)
}

func Check(sum int, tempSum int, values []int, i int, operators []string) ([]string, error) {
	multOperators, mulErr := Mult(sum, tempSum, values, i, operators)

	if mulErr == nil {
		return multOperators, nil
	}

	addOperators, addErr := Add(sum, tempSum, values, i, operators)
	if addErr == nil {
		return addOperators, nil
	}

	return nil, errors.New(mulErr.Error() + " " + addErr.Error())
}

func Add(sum int, tempSum int, values []int, i int, operators []string) ([]string, error) {
	if i > len(values)-1 {
		return nil, errors.New("out of range")
	}

	tempSum += values[i]

	if tempSum == sum {
		return append(operators, "+"), nil
	}

	if tempSum > sum {
		return nil, errors.New("over sum")
	}

	return Check(sum, tempSum, values, i+1, append(operators, "+"))
}

func Mult(sum int, tempSum int, values []int, i int, operators []string) ([]string, error) {
	if i > len(values)-1 {
		return nil, errors.New("out of range")
	}

	tempSum *= values[i]

	if tempSum == sum {
		return append(operators, "*"), nil
	}

	if tempSum > sum {
		return nil, errors.New("over sum")
	}

	return Check(sum, tempSum, values, i+1, append(operators, "*"))
}

func Part2() {
	scanner, file := util.OpenFileAndScanner("day7/day7input.txt")

	defer file.Close()

	var validSums []int

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

		var usedOps []string

		usedOps, err = Check2(sum, values[0], values, 1, usedOps)

		if err == nil {
			//fmt.Printf("%d %s\n", values, usedOps)
			validSums = append(validSums, sum)
		}
	}

	sum := 0

	for _, validSum := range validSums {
		sum += validSum
	}

	fmt.Println(sum)
}

func Check2(sum int, tempSum int, values []int, i int, operators []string) ([]string, error) {
	tempOps, mulErr := Mult2(sum, tempSum, values, i, operators)

	if mulErr == nil {
		return tempOps, nil
	}

	tempOps, addErr := Add2(sum, tempSum, values, i, operators)
	if addErr == nil {
		return tempOps, nil
	}

	tempOps, concatErr := Concat(sum, tempSum, values, i, operators)

	if concatErr == nil {
		return tempOps, nil
	}

	return nil, errors.New(mulErr.Error() + " " + addErr.Error())
}

func Add2(sum int, tempSum int, values []int, i int, operators []string) ([]string, error) {
	if i > len(values)-1 {
		return nil, errors.New("out of range")
	}

	tempSum += values[i]

	if tempSum == sum {
		return append(operators, "+"), nil
	}

	if tempSum > sum {
		return nil, errors.New("over sum")
	}

	return Check2(sum, tempSum, values, i+1, append(operators, "+"))
}

func Mult2(sum int, tempSum int, values []int, i int, operators []string) ([]string, error) {
	if i > len(values)-1 {
		return nil, errors.New("out of range")
	}

	tempSum *= values[i]

	if tempSum == sum {
		return append(operators, "*"), nil
	}

	if tempSum > sum {
		return nil, errors.New("over sum")
	}

	return Check2(sum, tempSum, values, i+1, append(operators, "*"))
}

func Concat(sum int, tempSum int, values []int, i int, operators []string) ([]string, error) {
	if i > len(values)-1 {
		return nil, errors.New("out of range")
	}

	tempSum, _ = strconv.Atoi(strconv.Itoa(tempSum) + strconv.Itoa(values[i]))

	if tempSum == sum {
		return append(operators, "||"), nil
	}

	if tempSum > sum {
		return nil, errors.New("over sum")
	}

	return Check2(sum, tempSum, values, i+1, append(operators, "||"))
}
