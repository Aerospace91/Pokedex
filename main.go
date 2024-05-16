package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Create a Scanner to read input from the Console
	version := "0.01"

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Hello Trainer, Welcome to Pokedex v%v\n", version)

	for {
		//Start For Loop and Continiously prompt '>'
		fmt.Print("> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		input := scanner.Text()

		// Check if the input is the exit command
		if input == "exit" {
			fmt.Println("Exiting Pokedex...")
			return
		}

		// Evaluate the Input
		result, err := evaluate(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println(result)
	}
}

func evaluate(input string) (string, error) {
	//Here you can implement evaluation logic
	// For now - Echoes Input

	return input, nil
}
