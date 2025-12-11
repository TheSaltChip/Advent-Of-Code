package day6

import "core:fmt"
import "core:os"
import "core:strconv"
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

part1 :: proc(string_data: ^string) {
	mat := make([dynamic][dynamic]string)
	defer {
		for col in mat {
			delete(col)
		}
		delete(mat)
	}

	for line in strings.split_lines_iterator(string_data) {
		elements, _ := strings.split(line, " ")
		row := make([dynamic]string)

		for elem in elements {
			if elem != "" {
				append(&row, elem)
			}
		}
		// fmt.println(row)
		append(&mat, row)
	}

	sum := 0
	col := 0

	for {
		com := mat[len(mat) - 1][col]
		mult := com == "*"
		line_sum := mult ? 1 : 0
		for y in 0 ..< len(mat) - 1 {
			num, _ := strconv.parse_int(mat[y][col])
			if mult {
				line_sum *= num
			} else {
				line_sum += num
			}
		}
		sum += line_sum
		col += 1
		if len(mat[0]) == col {
			break
		}
	}

	fmt.println(sum)
}

part2 :: proc(string_data: ^string) {
	lines := strings.split_lines(string_data^)
	op_line := lines[len(lines) - 1]

	sum := 0
	col_length := 0

	for col_start_index := 0;
	    col_start_index + 1 < len(op_line);
	    col_start_index += col_length + 1 {
		found := false
		col_length = 0
		for r in op_line[col_start_index + 1:] {
			if r == '*' || r == '+' {
				found = true
				break
			}
			col_length += 1
		}

		mult := rune(op_line[col_start_index]) == '*'

		section_num := mult ? 1 : 0

		col_end := col_length + col_start_index - 1

		if !found {
			col_end += 1
		}

		for j := col_end; j >= col_start_index; j -= 1 {
			col_num := 0
			digit := 1
			for k := 0; k < len(lines) - 1; k += 1 {
				if (lines[k][j] != ' ') {
					num := int(rune(lines[k][j]) - '0')
					col_num = col_num * digit + num
					digit = 10
				}
			}
			section_num = mult ? section_num * col_num : section_num + col_num
		}
		sum += section_num
	}

	fmt.println(sum)
}
