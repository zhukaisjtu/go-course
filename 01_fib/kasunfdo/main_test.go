package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`1
1
2
3
5
8
13
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestAbs(t *testing.T) {
	type test struct {
		input    int
		expected int
	}

	tests := []test{
		{input: 1, expected: 1},
		{input: 0, expected: 0},
		{input: -1, expected: 1},
	}

	for _, test := range tests {
		actual := abs(test.input)
		if test.expected != actual {
			t.Errorf("Unexpected output in main()\ninput: %d\nexpected: %q\nactual: %q", test.input, test.expected, actual)
		}
	}
}

func TestFib(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	type test struct {
		input    int
		expected string
	}

	tests := []test{
		{input: -7,
			expected: strconv.Quote(`1
-1
2
-3
5
-8
13
`)},
		{input: -1,
			expected: strconv.Quote("1\n")},
		{input: 1,
			expected: strconv.Quote("1\n")},
		{input: 0,
			expected: strconv.Quote("0\n")},
	}

	for _, test := range tests {
		fib(test.input)
		actual := strconv.Quote(buf.String())

		if test.expected != actual {
			t.Errorf("Unexpected output in main()\ninput: %d\nexpected: %q\nactual: %q", test.input, test.expected, actual)
		}

		buf.Reset()
	}
}
