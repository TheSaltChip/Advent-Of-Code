package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func StripBOM(str string) string {
	return strings.TrimPrefix(str, "\ufeff")
}

func ToIntArray(str string, sep string) ([]int, error) {
	elems := strings.Split(str, sep)

	ints := make([]int, len(elems))
	var err error

	for i := 0; i < len(elems); i++ {
		ints[i], err = strconv.Atoi(elems[i])

		if err != nil {
			return nil, err
		}
	}

	return ints, nil
}

func OpenFileAndScanner(path string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file), file
}
