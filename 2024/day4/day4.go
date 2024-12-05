package day4

import (
	"bufio"
	"log"
	"os"
	"slices"
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

	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}

	copyMatrix := make([][]rune, len(matrix))
	copy(copyMatrix, matrix)

	rowsString := createString(matrix)
	log.Println(rowsString)

	reversedRowsString := slices.Clone(rowsString)
	slices.Reverse(reversedRowsString)
	log.Println(reversedRowsString)

	n := len(copyMatrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			copyMatrix[i][j], copyMatrix[j][i] = copyMatrix[j][i], copyMatrix[i][j]
		}
	}

	columnsString := createString(copyMatrix)
	log.Println(columnsString)

	reversedColumnsString := slices.Clone(columnsString)
	slices.Reverse(reversedColumnsString)
	log.Println(reversedColumnsString)

	copy(copyMatrix, matrix)

}

func createString(matrix [][]rune) []rune {
	str := strings.Builder{}

	for _, runes := range matrix {
		for _, r := range runes {
			str.WriteRune(r)
		}
	}

	return []rune(str.String())
}