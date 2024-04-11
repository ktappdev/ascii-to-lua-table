package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readStdIn() (string, error) {
	// NOTE:
	var inputLines []string
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		inputLines = append(inputLines, input)
		//had to add this check to know
		if strings.TrimSpace(input) == "" {
			break
		}
	}
	return strings.Join(inputLines, ""), nil
}

func findLongestLine(input string) string {
	//NOTE: So we can know how wide the art is and pad the shorter rows
	lines := strings.Split(input, "\n")
	longestLine := ""
	for _, line := range lines {
		if len(line) > len(longestLine) {
			longestLine = line
		}
	}
	return longestLine
}

func toLuaTable(s string) string {
	var result strings.Builder
	first := true
	longestLine := findLongestLine(s)
	longestLineLen := len(longestLine)

	lines := strings.Split(s, "\n")
	for _, line := range lines {
		if !first {
			result.WriteString(", \n")
		}
		first = false

		lineLen := len(line)
		diff := longestLineLen - lineLen
		newline := strings.ReplaceAll(strings.ReplaceAll(line, "\\", "\\\\"), "\"", "\\\"")
		// newline := line
		var toPush string
		if lineLen < longestLineLen {
			padding := strings.Repeat(" ", diff)
			toPush = fmt.Sprintf("\"%s%s\"", newline, padding)
		} else {
			toPush = fmt.Sprintf("\"%s\"", newline)
		}

		result.WriteString(toPush)
	}

	return fmt.Sprintf("{%s}", result.String())
}

func main() {
	input, err := readStdIn()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	luaTable := toLuaTable(input)
	fmt.Printf("Lua table:\n%s\n", luaTable)
}
