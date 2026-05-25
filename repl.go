package main

import "strings"

func cleanInput(input string) []string {
	result := []string{}
	for _, word := range strings.Fields(input) {
		result = append(result, strings.ToLower(word))
	}
	return result
}
