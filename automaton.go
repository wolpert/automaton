package main

import (
	"flag"
	"fmt"
	"os"
)

func initArgs() (int, int, string) {
	widthPtr := flag.Int("width", 80, "Line Width")
	rowsPtr := flag.Int("rows", 80, "Rows to execute")
	rulePtr := flag.String("rule", "30", "Which rule to use")
	printRules := flag.Bool("rules", false, "Print rules")
	flag.Parse()
	if *printRules {
		fmt.Print("Rules: ")
		for key := range rules {
			fmt.Print(key, " ")
		}
		fmt.Println()
		os.Exit(0)
	}
	_, ok := rules[*rulePtr]
	if !ok {
		panic("Rule not found: " + *rulePtr)
	}
	return *widthPtr, *rowsPtr, *rulePtr
}

func main() {
	width, rows, rule := initArgs()
	world := make([]bool, width)
	for i := range world {
		world[i] = false
	}
	world[int(len(world)/2)] = true
	nextWorld := make([]bool, len(world))
	printSlice(world)
	for i := 0; i < rows; i++ {
		processSlice(world, nextWorld, rule)
		printSlice(nextWorld)
		world, nextWorld = nextWorld, world // no need to copy...
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

func processSlice(src []bool, dest []bool, rule string) {
	for i := 1; i < len(dest)-1; i++ { // all but first & last
		dest[i] = rules[rule].processRule(src[i-1 : i+2]) // second is exclusive
	}
}
