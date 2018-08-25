package main

import (
	"bufio"
	"fmt"
	"os"

	"../lua"
)

func main() {
	state := lua.NewState()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")

	for scanner.Scan() {
		line := scanner.Text()

		if err := state.LoadBuffer(line, "line"); err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			if err := state.PCall(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}

		fmt.Print("> ")
	}

	state.Close()
}
