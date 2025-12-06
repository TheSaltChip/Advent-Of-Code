package day3

import "core:fmt"
import "core:mem"
import "core:os"
import "core:strings"

main :: proc() {
	data, succ := os.read_entire_file(os.args[2])

	if !succ {
		fmt.eprintfln("Failed to read file", os.args[2])
		os.exit(1)
	}

	switch (os.args[1]) {
	case "1":
		part2(string(data), 2)
	case "2":
		part2(string(data), 12)
	}
}

part1 :: proc(input: string) {
	lines := strings.split(input, "\n")
	sum := 0
	for line in lines {
		collected := [2]int{}
		line_sum := 0

		for num_rune, index in line {
			num := int(num_rune - '0')

			if collected[0] == 0 ||
			   (collected[1] == 0 && num > collected[0] && index < len(line) - 2) {
				collected[0] = num
			} else if collected[1] == 0 {
				collected[1] = num
			} else {
				if num > collected[0] && index < len(line) - 2 {
					collected[0] = num
					collected[1] = 0
				} else if num > collected[1] {
					collected[1] = num
				}
			}
		}

		for digit in collected {
			line_sum = line_sum * 10 + digit
		}

		sum += line_sum
	}

	fmt.println(sum)
}

part2 :: proc(input: string, pattern_length: int) {
	sum := 0
	it := input
	collected := make([]int, pattern_length)
	defer delete(collected)
	for line in strings.split_lines_iterator(&it) {
		mem.zero_slice(collected[:])
		for num_rune, index in line {
			num := int(num_rune - '0')
			rem_in_line := len(line) - 1 - index
			for c, i in collected {
				if num > c && (pattern_length - 1 - i) <= rem_in_line {
					collected[i] = num
					mem.zero_slice(collected[i + 1:])
					break
				}
			}
		}
		line_sum := 0
		for digit in collected {
			line_sum = line_sum * 10 + digit
		}
		sum += line_sum
	}
	fmt.println(sum)
}
