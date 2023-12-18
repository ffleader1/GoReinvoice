package utils

import (
	"log"
	"testing"
)

func TestIsRelativePath(t *testing.T) {
	type testCase struct {
		input  string
		output bool
	}

	tests := map[string]testCase{
		"case 1": {
			"hello_world",
			true,
		},
		"case 2": {
			"C://Windows//hello_world",
			false,
		},
		"case 3": {
			"https://www.google.com",
			false,
		},
	}
	for tn, tc := range tests {
		if out := IsRelativePath(tc.input); tc.output != out {
			log.Fatalln("Error at testcase: ", tn, "Expect: ", tc.output, " - Got :", out)
		}
	}
}
