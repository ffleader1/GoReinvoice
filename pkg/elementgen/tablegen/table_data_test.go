package tablegen

import (
	"log"
	"testing"
)

func TestNumberToColumn(t *testing.T) {
	type testCase struct {
		input  int
		output Column
	}

	tests := map[string]testCase{
		"case1": {
			input:  12,
			output: "L",
		},
		"case2": {
			input:  21,
			output: "U",
		},
		"case3": {
			input:  36,
			output: "AJ",
		},
		"case4": {
			input:  64,
			output: "BL",
		},
	}

	for tn, tc := range tests {
		if out := numberToColumn(tc.input); tc.output != out {
			log.Fatalln("Error at testcase: ", tn, " - Got :", out)
		}
	}
}
