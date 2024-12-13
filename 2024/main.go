package main

import (
	"2024/day1"
	"2024/day2"
	"2024/day3"
	"2024/day4"
	"2024/day5"
	"2024/day6"
	"2024/day7"
	"2024/day8"
	"2024/day9"
)

func main() {

	day := 9

	switch day {
	case Day9:
		day9.Part1()
		day9.Part2()
	case Day8:
		day8.Part1()
		day8.Part2()
		break
	case Day7:
		day7.Part1()
		day7.Part2()
		break
	case Day6:
		day6.Part1()
		day6.Part2()
		break
	case Day5:
		day5.Part1()
		day5.Part2()
		break
	case Day4:
		day4.Part1()
		day4.Part2()
		break
	case Day3:
		day3.Part1()
		day3.Part2()
		break
	case Day2:
		day2.Part1()
		day2.Part2()
		break
	case Day1:
		day1.Part1()
		day1.Part2()
		break
	default:
		panic("unhandled default case")
	}
}

const (
	_    = iota
	Day1 = iota
	Day2
	Day3
	Day4
	Day5
	Day6
	Day7
	Day8
	Day9
	Day10
	Day11
	Day12
	Day13
	Day14
	Day15
	Day16
	Day17
	Day18
	Day19
	Day20
	Day21
	Day22
	Day23
	Day24
)
