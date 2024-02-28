package board

import (
	"strconv"
	"strings"
)

func GenBoard() string {
	example :=
		"ooooo\n" +
			"xxoxx\n" +
			"oxxxo\n" +
			"ooooo"
	output := "<table hx-get=\"/board/piece.html\" hx-trigger=\"load\" hx-swap=\"innerHTML\" hx-target=\"#x-0y-1\">\n <tbody>\n"
	lines := strings.Split(example, "\n")
	for i := 0; i < len(lines); i++ {
		output += "  <tr>\n"
		line := strings.Split(lines[i], "")
		for j := 0; j < len(line); j++ {
			if line[j] == "o" {
				output += "    <td class=\"p-0\">" + "<div " + genID(j, i) + genClass("white") + "><div>" + "</td>\n"
			} else {
				output += "    <td class=\"p-0\">" + "<div " + genID(j, i) + genClass("black") + "><div>" + "</td>\n"
			}
		}
		output += "  </tr>\n"
		// println("\n")
	}
	output += " </tbody>\n</table>"
	// println(output)
	return output
}

func genID(x int, y int) string {
	return " id=\"x-" + strconv.Itoa(x) + "y-" + strconv.Itoa(y) + "\""
}

func genClass(color string) string {
	return " class=\"flex items-center justify-center bg-" + color + " w-32 h-32\""
}
