package seestack

import (
	"fmt"
	"runtime/debug"
	"strings"
)

func Full() string {
	return string(debug.Stack())
}

func Short() string {
	lines := strings.Split(string(debug.Stack()), "\n")
	fmt.Println("lines len:", len(lines))

	num_words := len(lines) - 1
	num_words = num_words / 2
	num_words -= 3

	for i, l := range(lines) {
		fmt.Println(i, l)
	}

	fmt.Println("lines excluding some")
	fmt.Println("num_words:", num_words)

	var ret string

	cnt := 1
	for i, l := range(lines) {
		// exclude line 0 and odd lines
		if (i == 0 || i%2 != 0) {
			continue
		}
		// remove extra stuff in line to get only the package name
		// maybe include the line number? that is pretty nice
		fmt.Println(i, l)
		if (cnt == num_words) {
			break
		}
		cnt++
	}
	return "test"
}
