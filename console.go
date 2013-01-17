package commander

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// SetInteractive instructs commander to use the interactive console
func SetInteractive(interactive bool) {
	sharedCommander.interactive = true
}

// launchConsole launches the interactive console. The console accepts
// commands defined by Map(), the same as if you passed them directly
// on the command line. Each command will run the appropriate handler.
func launchConsole() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\nWelcome to the %s console! Type quit or exit when done.\n\n", sharedCommander.appName)

	for {
		fmt.Printf("\n> ")
		if line, err := reader.ReadString('\n'); err != nil {
			fmt.Println("An error occured while reading your input:", err)
		} else {

			line = strings.Replace(line, "\n", "", -1)

			if line == "quit" || line == "exit" {
				os.Exit(0)
			}

			fmt.Println("")

			args := strings.Split(line, " ")
			handleInvocation(args)
		}
	}

}
