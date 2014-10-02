package seestack

import (
	"runtime/debug"
	"strings"
	"regexp"
)

// Full
//
// gives the full call stack
//
func Full() string {
	return string(debug.Stack())
}

// ShortExclude
//
// gives a short version of the call stack excluding the top n calls.
// ShortExclude(0) will give all calls excluding this package
//
func ShortExclude(exclude int) string {

	lines := strings.Split(string(debug.Stack()), "\n")

	num_words := len(lines) - 1
	num_words = num_words / 2
	num_words -= 3

	var ret string

	cnt := 1
	for i, l := range(lines) {
		// exclude line 0 and odd lines
		if (i == 0 || i%2 != 0) {
			continue
		}
		// skip excluded calls
		if (cnt <= exclude) {
			cnt++
			continue
		}
		// make sure we don't get the lower level calls
		if (cnt > num_words) {
			break
		}
		cnt++

		// remove extra stuff in line to get only the package name
		// maybe include the line number? that is pretty nice
		r, _ := regexp.Compile(".*/")
		l = r.ReplaceAllString(l, "")
		r, _ = regexp.Compile("\\.go.*$")
		l = r.ReplaceAllString(l, "")

		if (ret == "") {
			ret = l
		} else {
			ret = l + "::" + ret
		}
	}
	return ret
}

// Short
//
// gives a short version of the call stack
//
func Short() string {
	return ShortExclude(1)
}
