package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Candidate struct {
	surname    string
	name       string
	patronymic string
	birthDay   string
	birthMonth string
	birthYear  string
}

func NewCandidate(data []string) *Candidate {
	return &Candidate{
		surname:    data[0],
		name:       data[1],
		patronymic: data[2],
		birthDay:   data[3],
		birthMonth: data[4],
		birthYear:  data[5],
	}
}

func (c *Candidate) CreateCode() string {
	unique := make(map[rune]bool)
	for _, r := range c.surname + c.name + c.patronymic {
		unique[r] = true
	}
	numberSum := len(unique)

	for _, let := range c.birthDay + c.birthMonth {
		if d, err := strconv.Atoi(string(let)); err == nil {
			numberSum += d * 64
		}
	}

	numberSum += int(strings.ToLower(c.surname)[0]-'a'+1) * 256

	code := fmt.Sprintf("%X", numberSum)
	code = code[len(code)-3:]

	if len(code) < 3 {
		code = strings.Repeat("0", 3-len(code)) + code
	}
	return code
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		candidate := NewCandidate(strings.Split(line, ","))
		fmt.Print(candidate.CreateCode(), " ")
	}
}
