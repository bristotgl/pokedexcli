package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		if scanner.Err() != nil {
			os.Exit(1)
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		fmt.Printf("Your command was: %s\n", commandName)
	}
}

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	words := strings.Fields(lowercase)
	return words
}
