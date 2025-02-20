package main

import (
	"fmt"
)

func main() {
	commands := []string{
		"BrowserHistory", "visit", "visit", "back", "visit", "visit", "back", "forward", "visit", "visit", "visit", "visit", "visit", "visit", "forward", "forward", "visit", "back", "visit", "visit", "visit", "visit", "forward", "visit", "visit", "visit",
	}
	args := [][]string{
		{"momn.com"}, {"bx.com"}, {"bjyfmln.com"}, {"3"}, {"ijtrqk.com"}, {"dft.com"}, {"10"}, {"10"}, {"yc.com"}, {"yhl.com"}, {"xynxvix.com"}, {"izfscdv.com"}, {"cdenhm.com"}, {"ocgcjz.com"}, {"5"}, {"5"}, {"gtd.com"}, {"9"}, {"hfeour.com"}, {"ghmh.com"}, {"nnm.com"}, {"knm.com"}, {"4"}, {"cbtg.com"}, {"acyvwod.com"}, {"mydr.com"},
	}

	var bh BrowserHistory
	var result []string

	for i, cmd := range commands {
		switch cmd {
		case "BrowserHistory":
			bh = Constructor1472(args[i][0])
			result = append(result, "null")
		case "visit":
			bh.Visit(args[i][0])
			result = append(result, "null")
		case "back":
			steps := toInt(args[i][0])
			result = append(result, bh.Back(steps))
		case "forward":
			steps := toInt(args[i][0])
			result = append(result, bh.Forward(steps))
		}
	}

	fmt.Println(result)
}

func toInt(s string) int {
	var res int
	fmt.Sscanf(s, "%d", &res)
	return res
}
