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
	previous_times := 0

	for line in strings.split_lines_iterator(&it) {
		command := line[0]
		amount, ok := strconv.parse_int(line[1:])

		if !ok {
			fmt.println("Failed to parse line: ", line[1:])
			os.exit(-1)
		}
		was_0 := start == 0

		previous_n := start

		if command == 'L' {
			start -= amount
		} else if command == 'R' {
			start += amount
		}
		after_command := start

		is_negative := start < 0

		if start == 0 {
			times_0 += 1
			fmt.printfln("%d %s %d %d %d", previous_n, line, after_command, start, times_0)
			previous_times = times_0
			continue
		}

		if start >= 100 {
			times_0 += start / 100
		} else if start < 0 {
			times_0 += -((start / 100) - 1)
		}

		start = abs(start % 100)
		/* for start < 0 || start >= 100 {
			if is_negative {
				start += 100
				if was_0 {
					was_0 = false
				} else {
					times_0 += 1
				}
				fmt.printfln("%d %s %d %d %d", previous_n, line, after_command, start, times_0)
			} else {
				start -= 100
				times_0 += 1
				fmt.printfln("%d %s %d %d %d", previous_n, line, after_command, start, times_0)
			}
		}

		is_changed := previous_times != times_0

		previous_times = times_0 */

			fmt.printfln("%d %s %d %d %d", previous_n, line, after_command, start, times_0)
		

	}

	fmt.println(times_0)
}

part1 :: proc() {
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
