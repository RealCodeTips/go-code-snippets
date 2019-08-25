package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		desc           string
		input          int
		expectedOutput string
	}{
		{
			desc:           "Input of 0 should return empty string",
			input:          0,
			expectedOutput: "",
		},
		{
			desc:           "Input of 1 should return '1'",
			input:          1,
			expectedOutput: "1",
		},
		{
			desc:           "Input of 15 should return '12fizz4buzzfizz78fizzbuzz11fizz1314fizzbuzz'",
			input:          15,
			expectedOutput: "12fizz4buzzfizz78fizzbuzz11fizz1314fizzbuzz",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := fizzBuzz(tC.input)

			if r != tC.expectedOutput {
				t.Errorf("Expected '%[1]s' to equal '%[2]s' but got '%[1]s'", r, tC.expectedOutput)
			}
		})
	}
}
