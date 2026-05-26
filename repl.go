package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a command: ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		output := cleaned[0]
		if output == "" {
			continue
		}
		fmt.Println("Your command was:", output)
		if output == "exit" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
		switch output {
		case "":
		}
	}
}
