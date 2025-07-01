package core

import (
	"fmt"
	"os"
	"strings"
	"memory/src/types"
)

var l []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}


func Tokenise(t string) []types.Item {
	var line int32 = 0
	var items []types.Item = []types.Item{}
	var current string = ""
	var depth int8 = 0
	var pointed_depth int8 = 0
	for i, char := range t {

		switch string(char) {

		case "]":
			depth--
			if depth < 0 {
				fmt.Printf("[73402] - At line %d - Unexpected closing bracket ']' (depth mismatch).\n", line)
				os.Exit(1)
			}
			if depth > 0 {
				current += "]"
				continue
			}
			t := &items[len(items)-1]
			for range pointed_depth {
				if t.SubFunction == nil {
					t.SubFunction = &types.Item{}
				}
				t = t.SubFunction
			}
			t.Closed = true
			t.Listed_args = current
			current = ""
			pointed_depth = 0

		case "[":
			depth++
			if depth > 1 {
				current += "["
				continue
			}
			item := types.Item{
				Name:        current,
				Listed_args: "",
				Opened:      true,
				Closed:      false,
				Line:        line,
				SubFunction: nil,
			}
			if pointed_depth > 0 {
				t := &items[len(items)-1]
				for range pointed_depth {
					if t.SubFunction == nil {
						t.SubFunction = &types.Item{}
					}
					t = t.SubFunction
				}
				if t.SubFunction == nil {
					t.SubFunction = &item
				} else {
					*t.SubFunction = item
				}
			} else {
				items = append(items, item)
			}
			current = ""

		case "$":
			if depth == 0 && current != "" {
				if strings.TrimSpace(current) != "" {
					fmt.Printf("[73402] - At line %d - %s is useless and not complete.\n", line, current)
					os.Exit(1)
				}
			}
			current += "$"

		case ">":
			if i == 0 || string(t[i-1]) != "-" {
				fmt.Printf("[73402] - At line %d - The superiority sign can only be used if preceded by an -.\n", line)
				os.Exit(1)
			}
			if len(items) == 0 {
				fmt.Printf("[73402] - At line %d - Can't match the next function with a previous one.\n", line)
				os.Exit(1)
			}
			pointed_depth++
			current = ""

		case "\n":
			line++
			if depth > 0 {
				current += "\n"
			} else {
				if current != "" && len(items) > 0 && !items[len(items)-1].Closed {
					items[len(items)-1].Listed_args += current + "\n"
					current = ""
				} else if strings.TrimSpace(current) != "" {
					fmt.Printf("[73402] - At line %d - Unhandled content at line end: '%s'\n", line, current)
					os.Exit(1)
				} else {
					current = ""
				}
			}

		default:
			if depth == 0 && current != "" && !strings.HasPrefix(current, "$") && !Contains(strings.ToLower(string(char)), l) {
				if len(items) > 0 && !items[len(items)-1].Closed {
					items[len(items)-1].Listed_args += current + string(char)
					current = ""
					continue
				} else {
					if strings.TrimSpace(string(char)) != "" {
						fmt.Printf("[73402] - At line %d - Unable to execute a function using %s.\n", line, string(char))
						os.Exit(1)
					}
				}
			}
			current += string(char)

		}
	}

	if current != "" && depth == 0 {
		if len(items) > 0 && !items[len(items)-1].Closed {
			items[len(items)-1].Listed_args += current
		} else {
			if strings.TrimSpace(current) != "" {
				fmt.Printf("[73402] - At line %d - %s is useless and not complete.\n", line, current)
				os.Exit(1)
			}
		}
	}

	return items
}

func Contains(t string, l []string) bool {
	for i := range l {
		if l[i] == t {
			return true
		}
	}
	return false
}