package day4

import "core:fmt"
import "core:os"
import "core:strings"
main :: proc() {
	data, ok := os.read_entire_file(os.args[2])

	if !ok {
		fmt.eprintln("Failed to read file ", os.args[2])
		os.exit(1)
	}

	string_data := string(data)

	lines := strings.split_lines(string_data)

	grid := make([][]rune, len(lines))
	defer {
		for row in grid {
			delete(row)
		}
		delete(grid)
	}

	for str, i in lines {
		grid[i] = make([]rune, len(str))
		for r, j in str {
			grid[i][j] = r
		}
	}

	switch os.args[1] {
	case "1":
		part1(grid)
	case "2":
		part2(grid)
	}
}

part1 :: proc(grid: [][]rune) {
	valid := 0

	for line, x in grid {
		for char, y in line {
			if char == '@' && check_perimeter(x, y, grid) {
				valid += 1
			}
		}
	}

	fmt.println(valid)
}

part2 :: proc(grid: [][]rune) {
	valid := 0
	alt_grid := copy_grid(grid)
	defer {
		for row in alt_grid {
			delete(row)
		}
		delete(grid)
	}
	for {
		removed := 0
		for line, x in grid {
			for char, y in line {
				if char == '@' && check_perimeter(x, y, grid) {
					removed += 1
					alt_grid[x][y] = '.'
				}
			}
		}

		for row, i in alt_grid {
			copy(grid[i], row)
		}

		if (removed == 0) {
			break
		}

		valid += removed
	}

	fmt.println(valid)
}

check_perimeter :: proc(x: int, y: int, grid: [][]rune) -> bool {
	x_start := x == 0 ? 0 : x - 1
	x_end := x == len(grid) - 1 ? x : x + 1
	y_start := y == 0 ? 0 : y - 1
	y_end := y == len(grid[0]) - 1 ? y : y + 1

	count := 0

	for i in x_start ..= x_end {
		for j in y_start ..= y_end {
			if i == x && j == y {
				continue
			}

			if grid[i][j] == '@' {
				count += 1
			}
		}
	}

	return count < 4
}

copy_grid :: proc(grid: [][]rune) -> [][]rune {
	dst := make([][]rune, len(grid))

	for row, i in grid {
		dst[i] = make([]rune, len(row))
		copy(dst[i], row)
	}

	return dst
}
