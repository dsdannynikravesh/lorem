package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	clip := flag.Bool("c", false, "copy output to clipboard")
	flag.Parse()

	numberOfParagraphs := 1
	if len(flag.Args()) > 0 {
		n, err := strconv.Atoi(flag.Args()[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "invalid number:", err)
			os.Exit(1)
		}
		numberOfParagraphs = n

	}

	lorem := getLorem(numberOfParagraphs)

	if *clip {
		err := copyToClipboard(lorem)
		if err != nil {
			fmt.Fprintln(os.Stderr, "could not copy to clipboard:", err)
		}

		fmt.Println("Copied to clipboard")
	} else {
		fmt.Println(lorem)
	}
}

func getLorem(n int) string {
	var b strings.Builder
	lorem := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	switch n {
	case 1:
		return lorem
	default:
		for i := range n {
			b.WriteString(lorem)
			if i != n-1 {
				b.WriteString("\n")
				b.WriteString("\n")
			}
		}
	}

	return b.String()
}

func copyToClipboard(text string) error {
	return clipboard.WriteAll(text)
}
