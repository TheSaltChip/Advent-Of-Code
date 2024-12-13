package day8

import (
	"2024/util"
	"fmt"
	"github.com/adam-lavrik/go-imath/ix"
)

type position struct {
	X, Y int
}

func (p position) Distance(pos position) position {
	return position{ix.Abs(p.X - pos.X), ix.Abs(p.Y - pos.Y)}
}

func (p position) Reflect(pos position) position {
	return position{p.X + (p.X - pos.X), p.Y + (p.Y - pos.Y)}
}

func (p position) ReflectWithFactor(pos position, factor int) position {
	return position{p.X + (p.X-pos.X)*factor, p.Y + (p.Y-pos.Y)*factor}
}

func (p position) AddAssign(pos position) position {
	return position{p.X + pos.X, p.Y + pos.Y}
}
func (p position) SubAssign(pos position) position {
	return position{p.X - pos.X, p.Y - pos.Y}
}

func (p position) Within(b bounds) bool {
	return b.minPos.X <= p.X &&
		b.minPos.Y <= p.Y &&
		b.maxPos.X >= p.X &&
		b.maxPos.Y >= p.Y
}

type bounds struct {
	minPos, maxPos position
}

func Part1() {
	scanner, file := util.OpenFileAndScanner("day8/day8input.txt")

	defer file.Close()

	var matrix [][]rune

	for i := 0; scanner.Scan(); i++ {
		line := util.StripBOM(scanner.Text())
		matrix = append(matrix, []rune(line))
	}

	var antennaMap = make(map[rune][]position)

	for x, column := range matrix {
		for y, item := range column {
			if item == '.' {
				continue
			}

			if antennaMap[item] == nil {
				antennaMap[item] = []position{{x, y}}
			} else {
				antennaMap[item] = append(antennaMap[item], position{x, y})
			}
		}
	}

	validArea := bounds{minPos: position{X: 0, Y: 0}, maxPos: position{X: len(matrix[0]) - 1, Y: len(matrix) - 1}}
	distinctAntiNodeCount := make(map[position]struct{})

	for _, positions := range antennaMap {
		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				dist0 := positions[i].Reflect(positions[j])
				dist1 := positions[j].Reflect(positions[i])

				if dist0.Within(validArea) {
					distinctAntiNodeCount[dist0] = struct{}{}
				}

				if dist1.Within(validArea) {
					distinctAntiNodeCount[dist1] = struct{}{}
				}
			}
		}
	}

	fmt.Println(len(distinctAntiNodeCount))
}

func Part2() {
	scanner, file := util.OpenFileAndScanner("day8/day9input.txt")

	defer file.Close()

	var matrix [][]rune

	for i := 0; scanner.Scan(); i++ {
		line := util.StripBOM(scanner.Text())
		matrix = append(matrix, []rune(line))
	}

	var antennaMap = make(map[rune][]position)

	for x, column := range matrix {
		for y, item := range column {
			if item == '.' {
				continue
			}

			if antennaMap[item] == nil {
				antennaMap[item] = []position{{x, y}}
			} else {
				antennaMap[item] = append(antennaMap[item], position{x, y})
			}
		}
	}

	validArea := bounds{minPos: position{X: 0, Y: 0}, maxPos: position{X: len(matrix[0]) - 1, Y: len(matrix) - 1}}
	distinctAntiNodeCount := make(map[position]struct{})

	fmt.Println(validArea)

	for _, positions := range antennaMap {
		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				dist0 := positions[i].Reflect(positions[j])
				for k := 1; dist0.Within(validArea); k++ {
					//	if matrix[dist0.X][dist0.Y] == '.' {
					matrix[dist0.X][dist0.Y] = '#'
					//	}
					distinctAntiNodeCount[dist0] = struct{}{}
					dist0 = positions[i].ReflectWithFactor(positions[j], k+1)
				}

				dist1 := positions[j].Reflect(positions[i])
				for k := 1; dist1.Within(validArea); k++ {
					//if matrix[dist1.X][dist1.Y] == '.' {
					matrix[dist1.X][dist1.Y] = '#'
					//}
					distinctAntiNodeCount[dist1] = struct{}{}
					dist1 = positions[j].ReflectWithFactor(positions[i], k+1)
				}
			}
		}
	}

	//count := 0
	//for _, runes := range matrix {
	//	for _, r := range runes {
	//		fmt.Print(string(r) + "")
	//		if r == '#' {
	//			count++
	//		}
	//	}
	//	fmt.Println()
	//}
	//fmt.Println(count)

	for _, val := range antennaMap {
		if len(val) > 1 {
			for _, item := range val {
				distinctAntiNodeCount[item] = struct{}{}
			}
		}
	}

	fmt.Println(len(distinctAntiNodeCount))
}
