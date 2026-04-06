package clientloop

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func GetInput() []string {
	fmt.Printf("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {return nil}
	input := scanner.Text()
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	return strings.Fields(input)
}

func PrintCommands() {
	fmt.Println("List of available commands:")
	fmt.Println("	*register:\n	     Register a new password\n	     Usage: register [password_name] [password]")
	fmt.Println("	*help:\n	     Shows available commands")
	fmt.Println("	*quit:\n	     Stop program")
}