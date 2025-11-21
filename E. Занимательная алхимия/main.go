package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const CAP_LIMIT = int(1e9)

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func buildRecipes(lines [][]int) (map[int][]int, map[int]bool) {
	unreachable := make(map[int]bool)
	saved := map[int][]int{
		1: {1, 0},
		2: {0, 1},
	}

	for ind, line := range lines {
		potion := ind + 3
		if _, exists := unreachable[potion]; exists {
			continue
		}
		if _, exists := saved[potion]; exists {
			continue
		}

		stack := []int{potion}
		temp := map[int][]int{
			potion: {len(line) - 1, 0, 0},
		}

		for len(stack) != 0 {
			curr_potion := stack[len(stack)-1]
			curr_index := temp[curr_potion][0]

			if curr_index < 0 {
				reqA := temp[curr_potion][1]
				reqB := temp[curr_potion][2]
				saved[curr_potion] = []int{reqA, reqB}
				stack = stack[0 : len(stack)-1]
				delete(temp, curr_potion)

				if len(stack) != 0 {
					parent_potion := stack[len(stack)-1]

					temp[parent_potion][1] += reqA
					temp[parent_potion][2] += reqB

					if temp[parent_potion][1] > CAP_LIMIT || temp[parent_potion][2] > CAP_LIMIT {
						for _, p := range stack {
							unreachable[p] = true
						}
						break
					}
				}
				continue
			}

			recipe_line := lines[curr_potion-3]
			x := recipe_line[curr_index]
			temp[curr_potion][0] -= 1

			if _, exists := saved[x]; exists {
				temp[curr_potion][1] += saved[x][0]
				temp[curr_potion][2] += saved[x][1]

				if temp[curr_potion][1] > CAP_LIMIT || temp[curr_potion][2] > CAP_LIMIT {
					for _, p := range stack {
						unreachable[p] = true
					}
					break
				}
			} else if _, exists := unreachable[x]; exists {
				for _, p := range stack {
					unreachable[p] = true
				}
				break
			} else if _, exists := temp[x]; exists {
				for _, p := range stack {
					unreachable[p] = true
				}
				break
			} else {
				stack = append(stack, x)
				new_line := lines[x-3]
				temp[x] = []int{len(new_line) - 1, 0, 0}
			}
		}
	}

	return saved, unreachable
}

func turnToIntLine(line []string) []int {
	intLine := make([]int, 0, len(line))

	for _, x := range line {
		intLine = append(intLine, mustAtoi(x))
	}

	return intLine
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	buf := make([]byte, 0, bufio.MaxScanTokenSize)
	scan.Buffer(buf, 2*1024*1024)
	scan.Scan()
	N, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}

	lines := make([][]int, 0, N-2)

	for range N - 2 {
		scan.Scan()
		line := strings.Split(scan.Text(), " ")[1:]
		intLine := turnToIntLine(line)

		lines = append(lines, intLine)
	}

	recipes, unreachable := buildRecipes(lines)

	// Пропускаем кол-во вопросов
	scan.Scan()

	for scan.Scan() {
		line := strings.Split(scan.Text(), " ")
		intLine := turnToIntLine(line)
		a := intLine[0]
		b := intLine[1]
		potion := intLine[2]

		if _, exists := unreachable[potion]; exists {
			fmt.Print("0")
		} else if a >= recipes[potion][0] && b >= recipes[potion][1] {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
}
