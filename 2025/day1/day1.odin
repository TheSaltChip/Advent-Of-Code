package day1

import "core:fmt"
import "core:os"
import "core:strconv"
import "core:strings"

main :: proc() {
	args := os.args
	data, ok := os.read_entire_file(args[1], context.allocator)

	if !ok {
		return
	}

	defer delete(data, context.allocator)

	it := string(data)

	start := 50
	times_0 := 0

	for line in strings.split_lines_iterator(&it) {
		fmt.println(line, start)
		command := line[0]
		amount, ok := strconv.parse_int(line[1:])

		if !ok {
			fmt.println("Failed to parse line: ", line[1:])
			os.exit(-1)
		}

		if command == 'L' {
			start -= amount
		} else if command == 'R' {
			start += amount
		}

		start %= 100

		if start < 0 {
			start += 100
		}

		if start == 0 {
			times_0 += 1
		}
	}

	fmt.println(times_0)
}

part1 :: proc () {args := os.args
	data, ok := os.read_entire_file(args[1], context.allocator)

	if !ok {
		return
	}

	defer delete(data, context.allocator)

	it := string(data)

	start := 50
	times_0 := 0

	for line in strings.split_lines_iterator(&it) {
		fmt.println(line, start)
		command := line[0]
		amount, ok := strconv.parse_int(line[1:])

		if !ok {
			fmt.println("Failed to parse line: ", line[1:])
			os.exit(-1)
		}

		if command == 'L' {
			start -= amount
		} else if command == 'R' {
			start += amount
		}

		start %= 100

		if start < 0 {
			start += 100
		}

		if start == 0 {
			times_0 += 1
		}
	}

	fmt.println(times_0)
}