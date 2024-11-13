package main

import (
	"fmt"
	"strings"
)

const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Magenta = "\033[35m"
const Cyan = "\033[36m"
const Gray = "\033[37m"
const White = "\033[97m"

func inRed(what string) string {
	return Red + what + Reset
}

func inGreen(what string) string {
	return Green + what + Reset
}

func display(data [][]string) {
	maxColSizes := make([]int, len(data[0]))
	for _, row := range data {
		for j, col := range row {
			if len(col) > maxColSizes[j] {
				maxColSizes[j] = len(col)
			}
		}
	}

	for _, row := range data {
		for j, col := range row {
			row[j] = padr(col, maxColSizes[j])
		}
	}

	for i, row := range data {
		fmt.Println(strings.ReplaceAll(strings.Join(row, inGreen("|")), "(null)", inRed("(null)")))
		if i == 0 {
			for j, size := range maxColSizes {
				if j > 0 {
					fmt.Print(inGreen("+"))
				}
				fmt.Print(inGreen(strings.Repeat("-", size)))
			}
			fmt.Println()
		}
	}
}

func padr(what string, newLength int) string {
	if len(what) >= newLength {
		return what
	}
	return what + strings.Repeat(" ", newLength-len(what))
}
