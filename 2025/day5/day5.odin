package day5

import "core:math"
import "core:fmt"
import "core:math/rand"
import "core:mem"
import "core:os"
import "core:slice"
import "core:sort"
import "core:strconv"
import "core:strings"

Range :: struct {
	start: int,
	end:   int,
}

main :: proc() {
	data, ok := os.read_entire_file(os.args[2])

	if !ok {
		fmt.eprintln("Failed to read file ", os.args[2])
		os.exit(1)
	}

	string_data := string(data)

	valid_ranges := make([dynamic]Range)
	ids := make([dynamic]int)
	defer {
		delete(valid_ranges)
		delete(ids)
	}

	at_ranges := true

	for line in strings.split_lines_iterator(&string_data) {
		if line == "" {
			at_ranges = false
			continue
		}

		if at_ranges {
			ranges_str := strings.split(line, "-")
			start, _ := strconv.parse_int(ranges_str[0])
			end, _ := strconv.parse_int(ranges_str[1])

			append(&valid_ranges, Range{start, end})
		} else {
			id, _ := strconv.parse_int(line)
			append(&ids, id)
		}
	}

	consolidated_valid_ranges := consolidate_ranges(valid_ranges)
	defer delete(consolidated_valid_ranges)

	switch os.args[1] {
	case "1":
		part1(consolidated_valid_ranges, ids)
	case "2":
		part2(consolidated_valid_ranges)
	}
}

part1 :: proc(valid_ranges: [dynamic]Range, ids: [dynamic]int) {
	fresh_ids := 0
	for id in ids {
		for range in valid_ranges {
			if id < range.start {
				break
			}

			if range.start <= id && range.end >= id {
				fresh_ids += 1
				break
			}
		}
	}
	fmt.println(fresh_ids)
}

part2 :: proc(valid_ranges: [dynamic]Range) {
	fresh_ids := 0
	for range in valid_ranges {
		fresh_ids += range.end - range.start + 1
	}
	fmt.println(fresh_ids)
}

consolidate_ranges :: proc(ranges: [dynamic]Range) -> [dynamic]Range {
	ranges_to_use := make([]Range, len(ranges))
	defer delete(ranges_to_use)
	copy(ranges_to_use[:], ranges[:])

	sort.quick_sort_proc(ranges_to_use[:], proc(r1, r2: Range) -> int {
		return r1.start - r2.start
	})

	consolidated_ranges := make([dynamic]Range)
	append(&consolidated_ranges, ranges_to_use[0])

	for i in 1 ..< len(ranges_to_use) {
		con_range := &consolidated_ranges[len(consolidated_ranges)-1]
		range := ranges_to_use[i]

		if con_range.end >= range.start - 1 {
			con_range.end = math.max(range.end, con_range.end)
		} else {
			append(&consolidated_ranges, range)
		}
	}

	return consolidated_ranges
}
