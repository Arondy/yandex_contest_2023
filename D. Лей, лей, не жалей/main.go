package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Order struct {
	start int
	end   int
	cost  int
}

func NewOrder(data []string) *Order {
	return &Order{
		start: mustAtoi(data[0]),
		end:   mustAtoi(data[1]),
		cost:  mustAtoi(data[2]),
	}
}

type Request struct {
	start int
	end   int
	type_ int
}

func NewRequest(data []string) *Request {
	return &Request{
		start: mustAtoi(data[0]),
		end:   mustAtoi(data[1]),
		type_: mustAtoi(data[2]),
	}
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func buildPrefixsums(scan *bufio.Scanner, N int) (map[int]int, []int, map[int]int, []int) {
	start2cost := make(map[int]int)
	end2time := make(map[int]int)
	start2cost[0] = 0
	end2time[0] = 0

	for range N {
		scan.Scan()
		lineElements := strings.Split(scan.Text(), " ")
		order := NewOrder(lineElements)

		start2cost[order.start] += order.cost
		end2time[order.end] += order.end - order.start
	}

	s2cKeys := slices.Sorted(maps.Keys(start2cost))
	e2tKeys := slices.Sorted(maps.Keys(end2time))

	for i := 1; i < len(s2cKeys); i++ {
		start2cost[s2cKeys[i]] += start2cost[s2cKeys[i-1]]
	}
	for i := 1; i < len(e2tKeys); i++ {
		end2time[e2tKeys[i]] += end2time[e2tKeys[i-1]]
	}

	return start2cost, s2cKeys, end2time, e2tKeys
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Scan()
	N, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}

	start2cost, s2cKeys, end2time, e2tKeys := buildPrefixsums(scan, N)

	// Пропускаем число запросов
	scan.Scan()

	for scan.Scan() {
		request := NewRequest(strings.Split(scan.Text(), " "))

		if request.type_ == 1 {
			leftIndex := sort.Search(len(s2cKeys), func(i int) bool { return s2cKeys[i] >= request.start })
			rightIndex := sort.Search(len(s2cKeys), func(i int) bool { return s2cKeys[i] > request.end })
			fmt.Printf("%d ", start2cost[s2cKeys[rightIndex-1]]-start2cost[s2cKeys[leftIndex-1]])
		} else {
			leftIndex := sort.Search(len(e2tKeys), func(i int) bool { return e2tKeys[i] >= request.start })
			rightIndex := sort.Search(len(e2tKeys), func(i int) bool { return e2tKeys[i] > request.end })
			fmt.Printf("%d ", end2time[e2tKeys[rightIndex-1]]-end2time[e2tKeys[leftIndex-1]])
		}
	}
}
