package main

import "fmt"

func main() {
	world := make([]bool, 80)
	for i := range world {
		world[i] = false
	}
	world[int(len(world)/2)] = true
	nextWorld := make([]bool, len(world))
	printSlice(world)
	for i := 0; i < 80; i++ {
		processSlice(world, nextWorld)
		printSlice(nextWorld)
		copy(world, nextWorld)
	}

}

func printSlice(slice []bool) {
	for _, value := range slice {
		if value {
			fmt.Print("1")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("\n")
}

func processSlice(src []bool, dest []bool) {
	for i := 1; i < len(dest)-1; i++ { // all but first & last
		value := processRule(src[i-1 : i+2]) // second is exclusive
		dest[i] = value
	}
}

func processRule(array []bool) bool {
	total := getValue(array)
	switch {
	case total == 1, total == 2, total == 3, total == 4:
		return true
	default:
		return false
	}
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
