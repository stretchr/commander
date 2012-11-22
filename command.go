package commander

import (
	"strings"
)

// command is a type used to create and manage individual command strings
type command struct {
	// definition is the string contining the command definition
	definition string

	// arguments is an array of all the arguments in the command string
	arguments []*argument
}

// makeCommand makes a new Command object and sets it up appropriately
func makeCommand(definition string) *command {

	c := new(command)
	c.definition = definition

	// make the arguments

	argumentStrings := strings.Split(definition, " ")
	c.arguments = make([]*argument, len(argumentStrings))

	for argumentIndex, value := range strings.Split(definition, " ") {
		c.arguments[argumentIndex] = MakeArgument(value)
	}

	return c

}
