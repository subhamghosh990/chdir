package chdir

import (
	"fmt"
	"testing"
)

func TestOutPutPath(t *testing.T) {
	fmt.Println("running")
	var tests = []struct { //Array of structure intput data expected data the initialized member
		i1, i2, expectedRes string
		testID              int
	}{
		{"/", "abc", "/abc", 1},
		{"/abc/def", "ghi", "/abc/def/ghi", 2},
		{"/abc/def", "..", "/abc", 3},
		{"/abc/def", "/abc", "/abc", 4},
		{"/abc/def", "/abc/klm", "/abc/klm", 5},
		{"/abc/def", "../..", "/", 6},
		{"/abc/def", "../../..", "/", 7},
		{"/abc/def", ".", "/abc/def", 8},
		{"/abc/def", "..klm", "..klm: No such file or directory", 9},
		{"/abc/def", "//////", "/", 10},
		{"/abc/def", "......", "......: No such file or directory", 11},
		{"/abc/def", "../gh///../klm/", "/abc/klm", 12},
	}

	pass := 0
	for _, test := range tests {
		op := OutPutPath(test.i1, test.i2)

		if op == test.expectedRes {
			t.Log("pass testcase ", test, " res : ", op)
			pass++
		} else {
			t.Log("fail testcase ", test, " res : ", op)
			t.Fatal()
		}
	}
}
