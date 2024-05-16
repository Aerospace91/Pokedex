package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Command represents a command in the REPL
type Command struct {
	Name        string
	Description string
	Handler     func(args []string) (string, error)
}

// Commands maps command names to their corresponding evaluation functions
var Commands = map[string]Command{}

func init() {
	Commands["echo"] = Command{
		Name:        "echo",
		Description: "Echoes back the provided text",
		Handler: func(args []string) (string, error) {
			return strings.Join(args, " "), nil
		},
	}

	Commands["help"] = Command{
		Name:        "help",
		Description: "Displays available commands and their descriptions",
		Handler: func(args []string) (string, error) {
			var output strings.Builder
			for _, cmd := range Commands {
				fmt.Fprintf(&output, "%s: %s\n", cmd.Name, cmd.Description)
			}
			return output.String(), nil
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanned := scanner.Scan()
		if !scanned {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}

		input := scanner.Text()

		// Check if the input is the exit command
		if strings.TrimSpace(input) == "exit" {
			fmt.Println("Exiting REPL...")
			return
		}

		// Split input into command and arguments
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}
		cmdName := parts[0]
		cmdArgs := parts[1:]

		// Look up the command and execute its handler
		cmd, ok := Commands[cmdName]
		if !ok {
			fmt.Println("Unknown command. Type 'help' to see available commands.")
			continue
		}

		result, err := cmd.Handler(cmdArgs)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println(result)
	}
}
