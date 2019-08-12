package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(atoi(""))
	fmt.Println(atoi("1"))
	fmt.Println(atoi("01"))
	fmt.Println(atoi("3345"))
}

func atoi(s string) (total int, err error) {
	// var base = 10
	if s == "" {
		err = errors.New("empty string")
		return
	}

	s = strings.TrimSpace(s)
	var neg bool
	switch s[0] {
	case '0':
		err = errors.New("not 10 jinzhi")
		return

	case '+':
		s = s[1:]

	case '-':
		neg = true
	}

	for _, c := range s {
		if !unicode.IsDigit(c) {
			err = errors.New("not digit for base 10")
			return
		}
		total *= 10
		total += int(c - '0')
	}

	if neg {
		total = -total
	}
	return
}
