package automaton

type rule struct {
	totals []int
}

func getValue(array []bool) int {
	total := 0
	if array[2] { // backwards, remember?
		total += 1
	}
	if array[1] {
		total += 2
	}
	if array[0] {
		total += 4
	}
	return total
}

var rules = map[string]rule{
	"30": {[]int{1, 2, 3, 4}},
}
