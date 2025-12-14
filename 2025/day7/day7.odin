package day7

import "core:fmt"
import "core:os"
import "core:slice"
import "core:sort"
import "core:strings"

main :: proc() {
	data, ok := os.read_entire_file(os.args[2])

	if !ok {
		fmt.eprintln("Failed to read file ", os.args[2])
		os.exit(1)
	}

	string_data := string(data)

	switch os.args[1] {
	case "1":
		part1(&string_data)
	case "2":
		part2(&string_data)
	}
}

part1 :: proc(data: ^string) {
	start_index := strings.index_rune(data^, 'S')

	lines := strings.split_lines(data^)

	ray_indices, i_to_remove, i_to_add :=
		make([dynamic]int), make([dynamic]int), make([dynamic]int)
	append(&ray_indices, start_index)
	defer {
		delete(ray_indices)
		delete(i_to_remove)
		delete(i_to_add)
	}

	sum := 0

	for line in lines[1:] {
		for ray_index, i in ray_indices {
			if rune(line[ray_index]) == '^' {
				sum += 1
				left_i, right_i := ray_index - 1, ray_index + 1

				append(&i_to_remove, i)

				if !slice.contains(i_to_add[:], left_i) &&
				   !slice.contains(ray_indices[:], left_i) {
					append(&i_to_add, left_i)
				}

				if !slice.contains(i_to_add[:], right_i) &&
				   !slice.contains(ray_indices[:], right_i) {
					append(&i_to_add, right_i)
				}
			}
		}

		sort.quick_sort_proc(i_to_remove[:], proc(a: int, b: int) -> int {
			return b - a
		})

		for r in i_to_remove {
			unordered_remove(&ray_indices, r)
		}

		append_elems(&ray_indices, ..i_to_add[:])

		clear(&i_to_add)
		clear(&i_to_remove)
	}

	fmt.println(sum)
}

Memo :: struct {
	value, iter_cost: int
}

part2 :: proc(data: ^string) {
	start_index := strings.index_rune(data^, 'S')

	lines := strings.split_lines(data^)
	walked_paths := make(map[int]Memo)
	defer delete(walked_paths)

	fmt.println(traverse(start_index, 0, &lines, &walked_paths), "Traverse called: ", (t))
}

t := 0

traverse :: proc(
	ray_index: int,
	line_no: int,
	lines: ^[]string,
	walked_paths: ^map[int]Memo,
) -> Memo {
	t+=1
	if line_no >= len(lines) {
		return Memo{1, 1}
	}

	if data, exists := walked_paths[line_no * len(lines) + ray_index]; exists {
		return data
	} 

	for i := line_no; i < len(lines); i += 2 {
		if rune(lines[i][ray_index]) == '^' {
			left_memo, right_memo := traverse(ray_index - 1, i + 2, lines, walked_paths),
				traverse(ray_index + 1, i + 2, lines, walked_paths)
			walked_paths[line_no * len(lines) + ray_index] = Memo{left_memo.value + right_memo.value, left_memo.iter_cost + right_memo.iter_cost +1 }
				
			return walked_paths[line_no * len(lines) + ray_index]
		}
	}

	return Memo{1,1}
}
