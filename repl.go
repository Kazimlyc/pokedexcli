package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		input := scanner.Text()

		input = strings.Fields(input)[0]

		args := []string{}

		input = strings.ToLower(input)

		command, exists := getCommands()[input]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(cfg, args...); err != nil {
			fmt.Println(err)
		}
	}

}
