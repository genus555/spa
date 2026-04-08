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
	fmt.Println("	*get:\n	     Shows the associated password\n	     Usage: get [password_name]")
	fmt.Println("	*delete:\n	     Deletes the associated password\n	     Usage: delete [password_name]")
	fmt.Println("	*list:\n	     Lists all saved password names\n	     Usage: list")
	fmt.Println("	*transfer:\n	     For transferring information to new device (More info on github.com/genus555/spa\n	     Usage: transfer [in/out]")
	fmt.Println("	*help:\n	     Shows available commands")
	fmt.Println("	*quit:\n	     Stop program")
}