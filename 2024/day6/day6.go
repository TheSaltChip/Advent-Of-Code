package day6

import (
	"2024/util"
	"fmt"
)

const (
	Up = iota
	Right
	Down
	Left
)

type Vector2Int struct {
	X, Y int
}

type GuardState struct {
	Pos         Vector2Int
	orientation int
}

func (v *Vector2Int) Add(pos Vector2Int) Vector2Int {
	return Vector2Int{v.X + pos.X, v.Y + pos.Y}
}

func Part1() {
	scanner, file := util.OpenFileAndScanner("day6/day6input.txt")

	defer file.Close()

	matrix := make([][]rune, 0, 130)

	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}

	var guardState GuardState

	for x, row := range matrix {
		for y, val := range row {
			if val == '^' {
				guardState.Pos.X = x
				guardState.Pos.Y = y
				guardState.orientation = Up
				break
			}
		}
	}

	searchDirection := Vector2Int{0, 0}
	visitedPositions := make(map[Vector2Int]struct{})
	visitedPositions[guardState.Pos] = struct{}{}
outer:
	for {
		switch guardState.orientation {
		case Up:
			searchDirection.Y = 0
			searchDirection.X = -1
			break
		case Down:
			searchDirection.Y = 0
			searchDirection.X = 1
			break
		case Left:
			searchDirection.Y = -1
			searchDirection.X = 0
			break
		case Right:
			searchDirection.Y = 1
			searchDirection.X = 0
			break
		}

		for {
			nextPos := guardState.Pos.Add(searchDirection)
			if nextPos.X < 0 || nextPos.X >= len(matrix) || nextPos.Y < 0 || nextPos.Y >= len(matrix[nextPos.X]) {
				break outer
			}

			if matrix[nextPos.X][nextPos.Y] == '#' {
				guardState.orientation = (guardState.orientation + 1) % 4
				break
			}

			guardState.Pos = nextPos

			visitedPositions[nextPos] = struct{}{}
		}
	}

	sumSteps := len(visitedPositions)

	fmt.Println(sumSteps)
}

func Part2() {

}
