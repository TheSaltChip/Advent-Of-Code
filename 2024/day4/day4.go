package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func Part1() {
	file, err := os.Open("day4/day4input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([][]rune, 0, 24)
	regXmas := regexp.MustCompile("(XMAS)")
	regSamx := regexp.MustCompile("(SAMX)")
	sumXmas := 0

	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}

	copyMatrix := make([][]rune, len(matrix))
	copy(copyMatrix, matrix)

	rowsStrings := createString(matrix)

	for _, str := range rowsStrings {
		sumXmas += len(regXmas.FindAllString(str, -1)) + len(regSamx.FindAllString(str, -1))
	}

	n := len(copyMatrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			copyMatrix[i][j], copyMatrix[j][i] = copyMatrix[j][i], copyMatrix[i][j]
		}
	}

	columnsStrings := createString(copyMatrix)

	for _, str := range columnsStrings {
		sumXmas += len(regXmas.FindAllString(str, -1)) + len(regSamx.FindAllString(str, -1))
	}

	var diagonals []string

	for i := 0; i < n; i++ {
		strBPosR := strings.Builder{}
		strBNegR := strings.Builder{}
		strBPosL := strings.Builder{}
		strBNegL := strings.Builder{}
		for j := 0; j < n; j++ {
			strBPosR.WriteRune(matrix[i+j][j])
			strBPosL.WriteRune(matrix[j][n-j-i-1])

			if i > 0 {
				strBNegR.WriteRune(matrix[j][j+i])
				strBNegL.WriteRune(matrix[j+i][n-j-1])
			}
			if i+j == len(matrix)-1 {
				break
			}
		}
		if strBPosR.Len() < 3 {
			continue
		}

		diagonals = append(diagonals, strBPosR.String())
		diagonals = append(diagonals, strBNegR.String())
		diagonals = append(diagonals, strBPosL.String())
		diagonals = append(diagonals, strBNegL.String())
	}

	for _, diagonal := range diagonals {
		sumXmas += len(regXmas.FindAllString(diagonal, -1)) + len(regSamx.FindAllString(diagonal, -1))
	}

	fmt.Println(sumXmas)
}

func createString(matrix [][]rune) []string {
	str := strings.Builder{}

	var strs []string

	for _, runes := range matrix {
		str.Reset()
		for _, r := range runes {
			str.WriteRune(r)
		}
		strs = append(strs, str.String())
	}

	return strs
}

func Part2() {
	file, err := os.Open("day4/day4input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([][]rune, 0, 24)
	sumXmas := 0

	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}

	for i := 0; i < len(matrix)-2; i++ {
		for j := 0; j < len(matrix[i])-2; j++ {
			if !((matrix[i][j] == 'M' && matrix[i+1][j+1] == 'A' && matrix[i+2][j+2] == 'S') ||
				(matrix[i][j] == 'S' && matrix[i+1][j+1] == 'A' && matrix[i+2][j+2] == 'M')) {
				continue
			}

			if !((matrix[i][j+2] == 'M' && matrix[i+1][j+1] == 'A' && matrix[i+2][j] == 'S') ||
				(matrix[i][j+2] == 'S' && matrix[i+1][j+1] == 'A' && matrix[i+2][j] == 'M')) {
				continue
			}

			sumXmas++
		}
	}

	fmt.Println(sumXmas)
}
