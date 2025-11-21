package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Time struct {
	day    int
	hour   int
	minute int
}

func (t *Time) add(day, hour, minute int) {
	t.day += day
	t.hour += hour
	t.minute += minute
}

func (t *Time) sub(day, hour, minute int) {
	t.day -= day
	t.hour -= hour
	t.minute -= minute
}

func (t *Time) countMinutes() int {
	return (t.day*24+t.hour)*60 + t.minute
}

type LogEntry struct {
	day    int
	hour   int
	minute int
	id     int
	letter string
}

func NewLogEntry(data []string) *LogEntry {
	return &LogEntry{
		day:    mustAtoi(data[0]),
		hour:   mustAtoi(data[1]),
		minute: mustAtoi(data[2]),
		id:     mustAtoi(data[3]),
		letter: data[4],
	}
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Scan()
	rocketsTime := make(map[int]*Time, 0)

	for scan.Scan() {
		line := scan.Text()
		log := NewLogEntry(strings.Fields(line))

		if log.letter == "B" {
			continue
		}

		r, exists := rocketsTime[log.id]

		if !exists {
			rocketsTime[log.id] = new(Time)
			r = rocketsTime[log.id]
		}

		if log.letter == "A" {
			r.sub(log.day, log.hour, log.minute)
		} else {
			r.add(log.day, log.hour, log.minute)
		}
	}

	sortedKeys := slices.Sorted(maps.Keys(rocketsTime))
	for _, key := range sortedKeys {
		fmt.Printf("%d ", rocketsTime[key].countMinutes())
	}
}
