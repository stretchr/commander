package commander

import (
	"fmt"
	"github.com/stretchr/objx"
	"os"
	"path"
	"strings"
	"sync"
)

// Default is used to register a default command that will be run when no
// arguments are given.
//
// The handler for the default command will be passed a nil map as no arguments
// are present.
const DefaultCommand = ""

// Commander provides methods and functionality to create a command line
// interface quickly and easily.
type commander struct {
	// commands contains all the mapped commands
	commands []*command

	// defaultRegistered stores whether a default has been registered or not
	defaultRegistered bool

	// interactive stores whether or not to use the interactive console
	interactive bool

	// appName stores the name of the currently running application
	appName string
}

// initOnce is used to guarantee that the sharedCommander is initialized only once.
var initOnce sync.Once

// sharedCommander is the shared instance of the Commander type
var sharedCommander *commander

// incomingArgs is the array of arguments to be analyzed. This exists to facilitate
// testing.
var incomingArgs []string

// commandMap builds a map of indentifier,value to be passed to the handler
func commandMap(cmd *command, args []string) map[string]interface{} {
	argMap := make(map[string]interface{})
	for i, a := range cmd.arguments {
		if len(args) <= i {
			break
		}
		if !a.isLiteral() {
			if !a.isVariable() {
				argMap[a.identifier] = args[i]
			} else {
				if len(cmd.arguments) == len(args) {
					argMap[a.identifier] = args[i]
				} else {
					argMap[a.identifier] = args[i:]
				}
			}
		}
	}
	return argMap
}

// printUsage prints the usage of the program
func printUsage(cmd *command) {

	if cmd == nil {
		if !sharedCommander.interactive {
			fmt.Printf("\nusage: %s <command> [arguments]\n\n", sharedCommander.appName)
		}
		for _, cmd := range sharedCommander.commands {
			if !cmd.isDefaultCommand() {
				fmt.Printf("    %s - %s\n", cmd.definition, cmd.summary)
			}
		}
	} else {
		fmt.Printf("\n\"%s\" usage:\n\n", cmd.arguments[0].literal)
		fmt.Printf("    %s - %s\n", cmd.definition, cmd.summary)
		fmt.Printf("    %s\n", cmd.description)
	}
	fmt.Println()

}

// moveHelpToEnd moves the help entry to the end of the array for printing
func moveHelpToEnd() {
	length := len(sharedCommander.commands)

	for i := 0; i < length-1; i++ {
		sharedCommander.commands[i], sharedCommander.commands[i+1] = sharedCommander.commands[i+1], sharedCommander.commands[i]
	}
}

// initialize sets up various internal fields to ready the system. If this is not
// called, Commander will not function.
func initialize() {
	initOnce.Do(func() {
		sharedCommander = new(commander)

		sharedCommander.appName = path.Base(os.Args[0])
		if extension := path.Ext(os.Args[0]); extension != "" {
			sharedCommander.appName = strings.Replace(sharedCommander.appName, extension, "", 1)
		}

		Map("help [arg=(string)]", "Prints help and usage",
			"Prints help and usage for the commands. \"help <command>\" will print additional information about the command.",
			func(args objx.Map) {
				printed := false
				if len(args) == 1 {
					for _, cmd := range sharedCommander.commands {
						if cmd.arguments[0].literal == args["arg"].(string) {
							printUsage(cmd)
							printed = true
						}
					}
				}
				if !printed {
					printUsage(nil)
				}
			})
	})
}

// execute fires up the commander system, either launching the interactive
// console, or executing the command provided by the arguments
func execute() {

	moveHelpToEnd()

	if sharedCommander.interactive && len(os.Args) == 1 {
		launchConsole()
		return
	}

	if incomingArgs == nil {
		incomingArgs = os.Args[1:]
	}

	// handle the arguments passed during program invocation
	handleInvocation(incomingArgs)

}

// handleInvocation analyzes the arguments and executes the
// appropriate command handler function
func handleInvocation(args []string) {

	executed := false
	closestMatchCount := 0
	var closestMatch *command

	executeDefault := len(args) == 0

	if executeDefault {
		for _, cmd := range sharedCommander.commands {
			if cmd.isDefaultCommand() {
				cmd.handler(nil)
				executed = true
			}
		}
	} else {
		for _, cmd := range sharedCommander.commands {
			if represents, matchCount := cmd.represents(args); represents {
				argMap := commandMap(cmd, args)
				cmd.handler(argMap)
				executed = true
			} else {
				if matchCount > closestMatchCount {
					closestMatchCount = matchCount
					closestMatch = cmd
				}
			}
		}
	}
	if !executed {
		printUsage(closestMatch)
	}

}

// Map is used to map a definition string to a handler function. If the arguments
// given on the command line are represented by the definition string, the
// handler function will be called.
func Map(definition, summary, description string, handler Handler) {

	if sharedCommander == nil {
		panic("Initialize must be called before Map")
	}

	if definition == DefaultCommand {
		if sharedCommander.defaultRegistered {
			panic("Only one default command can be registered.")
		} else {
			sharedCommander.defaultRegistered = true
		}
	}

	newCommand := makeCommand(definition, summary, description, handler)

	for _, cmd := range sharedCommander.commands {
		if cmd.isEqualTo(newCommand) {
			panic("Each command must have a unique signature.")
		}
	}

	sharedCommander.commands = append(sharedCommander.commands, newCommand)

}
