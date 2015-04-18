package automaton

import "testing"

type testpair struct {
	array    []bool
	expected byte
}

var tests = []testpair{
	{[]bool{false, false, false}, 0},
	{[]bool{false, false, true}, 1},
	{[]bool{false, true, false}, 2},
	{[]bool{false, true, true}, 3},
	{[]bool{true, false, false}, 4},
	{[]bool{true, false, true}, 5},
	{[]bool{true, true, false}, 6},
	{[]bool{true, true, true}, 7},
}

func TestGetValue(t *testing.T) {
	for _, pair := range tests {
		result := getValue(pair.array)
		if result != pair.expected {
			t.Error("For", pair.array, "expected", pair.expected, "got", result)
		}
	}
}
