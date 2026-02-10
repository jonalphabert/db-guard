package logviewer

import (
	"fmt"
)

func PrintLine(line string, noColor bool) {
	level := extractLevel(line)
	if !noColor {
		c, ok := levelColors[level]
		if ok {
			c.Printf("%s: ", level)
			fmt.Println(line[len(level)+3:])
		} else {
			fmt.Printf("%s: ", line)
		}
		
	} else {
		fmt.Printf("%s: %s\n", level, line[len(level)+3:])
	}
}