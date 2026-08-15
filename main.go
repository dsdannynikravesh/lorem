package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	numberOfParagraphs := 1
	if len(os.Args) > 2 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "invalid number:", err)
			os.Exit(1)
		}
		numberOfParagraphs = n

	}

	fmt.Println(numberOfParagraphs)

	lorem := getLorem(numberOfParagraphs)

	err := copyToClipboard(lorem)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not copy to clipboard:", err)
	}

	fmt.Println(getLorem(numberOfParagraphs))
}

func getLorem(n int) string {
	var b strings.Builder
	lorem := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	switch n {
	case 1:
		return lorem
	default:
		for range n {
			b.WriteString(lorem)
			b.WriteString("\n")
		}
	}

	return b.String()
}

func copyToClipboard(text string) error {
	cmd := exec.Command("wl-copy")
	cmd.Stdin = strings.NewReader(text)

	return cmd.Run()
}
