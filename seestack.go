package seestack

import (
	//"fmt"
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
// Gives a short version of the call stack excluding the top n calls and
// function calls. ShortExclude(0) will give all package names excluding
// this package.
//
func ShortExclude(exclude int) string {

	lines := strings.Split(string(debug.Stack()), "\n")

	//for i, l := range(lines) {
	//	fmt.Println(i, l)
	//}

	num_words := len(lines) - 1
	num_words = num_words / 2
	num_words -= 2

	//fmt.Println("num_words", num_words)

	var ret string
	var current_pkg string
	//var func_stack string
	cnt := 1
	for i, l := range(lines) {
		// exclude line 0 and odd lines
		if (i == 0 || i%2 != 0) {
			continue
		}
		// exclude lines that don't contain a go file name
		m, _ := regexp.Match("\\.go", []byte(l))
		if (!m) {
			continue
		}
		// skip exclude number of calls
		if (cnt <= exclude) {
			cnt++
			continue
		}

		// remove extra stuff in line to get only the package name
		// TODO: have options for line number/
		r, _ := regexp.Compile(".*/")
		l = r.ReplaceAllString(l, "")
		r, _ = regexp.Compile("\\s+\\(.*$")
		l = r.ReplaceAllString(l, "")

		// when called from a function there will be a package name for the
		// the function and the package

		// TODO: this is the option to show the functions along with the packages.
		// if show_funcs == 0 then skip packages calling their own functions.
		// may make this into a global config option(do this but keep it out of
		// this package). Or could also make it paramter. make generic function
		// for modifying stack will multiple options. have ShortExclude() call 
		// ShortStack(exclude int, showfuncs bool)
		show_funcs := false
		if (show_funcs) {
			// get the function name from the line below
			r, _ = regexp.Compile("^\\W*")
			func_name := r.ReplaceAllString(lines[i+1], "")
			r, _ = regexp.Compile(":.*")
			func_name = r.ReplaceAllString(func_name, "")

			l += "." + func_name
		} else {

			// if the package is the same as last time then we were called from a
			// function within that package
			if (current_pkg == l) {
				continue
			} else {
				current_pkg = l
			}
		}

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
