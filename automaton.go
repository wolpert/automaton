package main

import (
	"flag"
	"fmt"
)

func main() {
	width, rows := initArgs()
	world := make([]bool, width)
	for i := range world {
		world[i] = false
	}
	world[int(len(world)/2)] = true
	nextWorld := make([]bool, len(world))
	printSlice(world)
	for i := 0; i < rows; i++ {
		processSlice(world, nextWorld)
		printSlice(nextWorld)
		world, nextWorld = nextWorld, world // no need to copy...
	}
}

func initArgs() (int, int) {
	widthPtr := flag.Int("width", 80, "Line Width")
	rowsPtr := flag.Int("rows", 80, "Rows to execute")
	flag.Parse()
	fmt.Println("width=", *widthPtr, "rows=", *rowsPtr)
	return *widthPtr, *rowsPtr
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
		dest[i] = rules["30"].processRule(src[i-1 : i+2]) // second is exclusive
	}
}
