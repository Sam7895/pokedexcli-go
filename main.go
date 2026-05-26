package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
