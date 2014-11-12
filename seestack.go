package seestack

import (
	"regexp"
	"runtime/debug"
	"strings"
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
	//   fmt.Println(i, l)
	//}

	var ret string
	var current_pkg string
	//var func_stack string
	cnt := 1
	for i, l := range lines {
		// exclude line 0 and odd lines
		if i == 0 || i%2 != 0 {
			continue
		}
		// exclude lines that don't contain a go file name
		m, _ := regexp.Match("\\.go", []byte(l))
		if !m {
			continue
		}

		// remove extra stuff in line to get only the package name
		// TODO: have options for line number/
		r, _ := regexp.Compile(".*/")
		l = r.ReplaceAllString(l, "")
		r, _ = regexp.Compile("\\s+\\(.*$")
		l = r.ReplaceAllString(l, "")

		// get the package name. ex: seestack
		r, _ = regexp.Compile("\\.go:\\d+")
		pacname := r.ReplaceAllString(l, "")

		// get the line number. ex: 42
		r, _ = regexp.Compile(".*\\.go:")
		linenum := r.ReplaceAllString(l, "")

		// when called from a function there will be a package name for the
		// the function and the package

		// TODO: this is the option to show the functions along with the packages.
		// if show_funcs == true then skip packages calling their own functions.
		// may make this into a global config option(do this but keep it out of
		// this package). Or could also make it paramter. make generic function
		// for modifying stack will multiple options. have ShortExclude() call
		// ShortStack(exclude int, showfuncs bool)
		//show_funcs := false
		//if show_funcs {
		//   // get the function name from the line below
		//   r, _ = regexp.Compile("^\\W*")
		//   func_name := r.ReplaceAllString(lines[i+1], "")
		//   r, _ = regexp.Compile(":.*")
		//   func_name = r.ReplaceAllString(func_name, "")

		//   l = pacname + "." + func_name
		//} else {

		// if the package is the same as last time then we were called from a
		// function within that package
		if current_pkg == pacname {
			continue
		} else {
			current_pkg = pacname
		}
		//}

		// skip exclude number of packages
		if cnt <= exclude {
			cnt++
			continue
		}

		if ret == "" {
			ret = pacname + ":" + linenum
		} else {
			ret = pacname + ":" + linenum + "::" + ret
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

// LastFile
//
// gives the last .go file in the stack
//
func LastFile() string {
	s := ShortExclude(1)

	r, _ := regexp.Compile(":.*$")
	s = r.ReplaceAllString(s, "")

	return s
}
