package main

import "strings"

func cleanInput(text string) []string {
	result := []string{}
	arr := strings.Fields(text)
	for _, i := range arr {
		a := strings.TrimSpace(i)
		a = strings.ToLower(a)
		result = append(result, a)
	}

	return result
}
