package seestack

import (
	"fmt"
	"testing"
)

func recurse(depth int, t *testing.T) {
	tse := ShortExclude(0)
	lf := LastFile()
	fmt.Println(depth, tse, lf)

	if tse != "testing:447::seestack_test:009::seestack:024" {
		fmt.Println("TestAll.recurse(", depth, "): ShortExclude did not return "+
			"expected value")
		t.Fail()
		return
	}
	if lf != "testing" {
		fmt.Println("TestAll.recurse(", depth, "): LastFile did not return "+
			"expected value")
		t.Fail()
		return
	}

	if depth <= 0 {
		return
	}
	depth--
	recurse(depth, t)
}

func TestAll(t *testing.T) {
	recurse(3, t)
}
