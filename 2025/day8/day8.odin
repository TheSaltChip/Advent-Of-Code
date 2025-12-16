package day8

import "core:fmt"
import "core:math"
import "core:os"
import "core:slice"
import "core:strconv"
import "core:strings"

JunctionBox :: struct {
	x, y, z: int,
}

Tuple :: struct($T: typeid) {
	left, right: T,
}

Entry :: struct($K, $V: typeid) {
	key:   K,
	value: V,
}

main :: proc() {
	data, ok := os.read_entire_file(os.args[2])

	if !ok {
		fmt.eprintln("Failed to read file ", os.args[2])
		os.exit(1)
	}

	string_data := string(data)

	lines, err := strings.split_lines(string_data)
	juntion_boxes := make([]JunctionBox, len(lines))
	defer delete(juntion_boxes)

	for line, i in lines {
		coords_str, _ := strings.split(line, ",")
		x, okx := strconv.parse_int(coords_str[0])
		y, oky := strconv.parse_int(coords_str[1])
		z, okz := strconv.parse_int(coords_str[2])

		if okx && oky && okz {
			juntion_boxes[i] = JunctionBox{x, y, z}
		}
	}

	distances_map := make(map[Tuple(JunctionBox)]f32)
	sort_list := make([dynamic]Entry(Tuple(JunctionBox), f32))

	defer {
		delete(distances_map)
		delete(sort_list)
	}

	for b1, i in juntion_boxes {
		for b2, j in juntion_boxes {
			if i == j {
				continue
			}

			tuple1 := Tuple(JunctionBox){b1, b2}
			tuple2 := Tuple(JunctionBox){b2, b1}

			if data, exists := distances_map[tuple1]; !exists {
				if data, exists := distances_map[tuple2]; !exists {
					dist := euclidean_dist(tuple1)
					append(&sort_list, Entry(Tuple(JunctionBox), f32){tuple1, dist})
					distances_map[tuple1] = dist
				}
			}
		}
	}

	slice.sort_by(sort_list[:], proc(a, b: Entry(Tuple(JunctionBox), f32)) -> bool {
		return a.value < b.value
	})

	circuits := make([dynamic][dynamic]JunctionBox)
	defer {
		for arr in circuits {
			delete(arr)
		}
		delete(circuits)
	}

	list := make([dynamic]JunctionBox)
	append_elems(&list, sort_list[0].key.left, sort_list[0].key.right)
	append(&circuits, list)

	switch os.args[1] {
	case "1":
		part1(sort_list, &circuits)
	case "2":
		part2(len(juntion_boxes), sort_list, &circuits)
	}
}

part1 :: proc(
	sort_list: [dynamic]Entry(Tuple(JunctionBox), f32),
	circuits: ^[dynamic][dynamic]JunctionBox,
) {
	outer: for entry, i in sort_list[1:] {
		if i == 999 {
			break
		}

		from, to := entry.key.left, entry.key.right

		for j := 0; j < len(circuits); j += 1 {
			if slice.contains(circuits[j][:], from) && !slice.contains(circuits[j][:], to) {
				if !merge_circuits(circuits, to, j) {
					append(&circuits[j], to)
				} else {
					j -= 1
				}
				continue outer
			} else if slice.contains(circuits[j][:], to) && !slice.contains(circuits[j][:], from) {
				if !merge_circuits(circuits, from, j) {
					append(&circuits[j], from)
				} else {
					j -= 1
				}
				continue outer
			} else if slice.contains(circuits[j][:], to) && slice.contains(circuits[j][:], from) {
				continue outer
			}
		}

		list := make([dynamic]JunctionBox)
		append_elems(&list, from, to)
		append(circuits, list)
	}

	lengths := make([]int, len(circuits))
	defer delete(lengths)

	for c, i in circuits {
		lengths[i] = len(c)
	}

	slice.sort_by(lengths, proc(a, b: int) -> bool {
		return b < a
	})

	sum := 1
	for i := 0; i < 3; i += 1 {
		sum *= lengths[i]
	}

	fmt.println(sum)
}

part2 :: proc(
	num_junction_boxes: int,
	sort_list: [dynamic]Entry(Tuple(JunctionBox), f32),
	circuits: ^[dynamic][dynamic]JunctionBox,
) {
	last_appended: Tuple(JunctionBox) = sort_list[0].key
	outer: for entry in sort_list[1:] {
		from, to := entry.key.left, entry.key.right

		for j := 0; j < len(circuits); j += 1 {
			appended := false
			if slice.contains(circuits[j][:], from) && !slice.contains(circuits[j][:], to) {
				if !merge_circuits(circuits, to, j) {
					append(&circuits[j], to)
				} else {
					j -= 1
				}
				appended = true
			} else if slice.contains(circuits[j][:], to) && !slice.contains(circuits[j][:], from) {
				if !merge_circuits(circuits, from, j) {
					append(&circuits[j], from)
				} else {
					j -= 1
				}
				appended = true
			} else if slice.contains(circuits[j][:], to) && slice.contains(circuits[j][:], from) {
				continue outer
			}

			if appended {
				if len(circuits) == 1 && len(circuits[0]) == num_junction_boxes {
					last_appended = entry.key
					break outer
				}
				continue outer
			}
		}

		list := make([dynamic]JunctionBox)
		append_elems(&list, from, to)
		append(circuits, list)
	}

	fmt.println(last_appended.left.x * last_appended.right.x)
}


euclidean_dist :: proc(t: Tuple(JunctionBox)) -> f32 {
	left, right := t.left, t.right
	return math.sqrt_f32(
		math.pow_f32(f32(left.x - right.x), 2) +
		math.pow_f32(f32(left.y - right.y), 2) +
		math.pow_f32(f32(left.z - right.z), 2),
	)
}

merge_circuits :: proc(
	circuits: ^[dynamic][dynamic]JunctionBox,
	box: JunctionBox,
	index: int,
) -> bool {
	for &circuit, i in circuits {
		if slice.contains(circuit[:], box) && index != i {
			append_elems(&circuits[index], ..circuit[:])
			delete(circuit)
			unordered_remove(circuits, i)
			return true
		}
	}
	return false
}
