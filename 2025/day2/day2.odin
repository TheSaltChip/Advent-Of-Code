package day2

import "core:fmt"
import "core:os"
import "core:strconv"
import "core:strings"
import "core:unicode/utf8"

main :: proc() {
	part := os.args[1]
	data, ok_read := os.read_entire_file(os.args[2], context.allocator)

	if !ok_read {
		fmt.println("Failed to read file ", os.args[1])
		os.exit(1)
	}

	switch part {
	case "1":
		part1(string(data))
	case "2":
		part2(string(data))
	}
}

part2 :: proc(input: string) {
	ranges, ok_split := strings.split(input, ",")

	if ok_split != nil {
		fmt.println(ok_split)
		os.exit(2)
	}
	invalid_ids := 0

	for range in ranges {
		invalid_ids += check_range_part2(range)
	}

	fmt.println(invalid_ids)
}

check_range_part2 :: proc(range: string) -> int {
	arr := strings.split(range, "-")

	start, _ := strconv.parse_int(arr[0])
	end, _ := strconv.parse_int(arr[1])

	invalid_ids_sum := 0

	outer: for i in start ..= end {
		num_string := utf8.string_to_runes(fmt.tprintf("%d", i))
		defer delete(num_string)

		max_seq_len := len(num_string) / 2

		sequence: for curr_seq_len in 1 ..= max_seq_len {
			seq := make([dynamic]rune)
			defer delete(seq)
			for k in 0 ..< curr_seq_len {
				if len(seq) > 0 && len(num_string) < k + 1 && seq[k] != num_string[k + 1] {
					continue sequence
				}

				append(&seq, num_string[k])
			}

			if (len(num_string) % curr_seq_len != 0) {
				continue
			}

			for k := curr_seq_len; k < len(num_string); k += 1 {
				if seq[k % curr_seq_len] != num_string[k] {
					continue sequence
				}
			}

			invalid_ids_sum += i
			continue outer
		}
	}

	return invalid_ids_sum
}

part1 :: proc(input: string) {
	ranges, ok_split := strings.split(input, ",")

	if ok_split != nil {
		fmt.println(ok_split)
		os.exit(2)
	}

	invalid_ids := 0

	for range in ranges {
		invalid_ids += check_range_part1(range)
	}

	fmt.println(invalid_ids)
}

check_range_part1 :: proc(range: string) -> int {
	arr := strings.split(range, "-")

	start, _ := strconv.parse_int(arr[0])
	end, _ := strconv.parse_int(arr[1])

	invalid_ids_sum := 0

	for i in start ..= end {
		num_string := fmt.tprintf("%d", i)
		max_seq_len := len(num_string) / 2
		if num_string[0:max_seq_len] == num_string[max_seq_len:] {
			invalid_ids_sum += i
		}
	}

	return invalid_ids_sum
}
