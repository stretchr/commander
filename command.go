package commander

import (
	"strings"
)

// command is a type used to create and manage individual command strings
type command struct {
	// definition is the original string contining the command definition
	definition string

	// arguments is an array of all the arguments in the command string
	arguments []*argument
}

// makeCommand makes a new Command object and sets it up appropriately
func makeCommand(definition string) *command {

	c := new(command)
	c.definition = definition

	// make the arguments

	argumentStrings := strings.Split(definition, DelimiterArgumentSeparator)
	c.arguments = make([]*argument, len(argumentStrings))

	for argumentIndex, value := range strings.Split(definition, DelimiterArgumentSeparator) {
		c.arguments[argumentIndex] = makeArgument(value)
		if c.arguments[argumentIndex].isVariable() && argumentIndex != len(argumentStrings)-1 {
			return nil
		}
	}

	return c

}

// represents determines if this command represents the array of arguments
func (c *command) represents(rawArgs []string) bool {

	argIndex := 0
	for rawArgIndex, _ := range rawArgs {

		if argIndex == len(c.arguments)-1 && rawArgIndex != len(rawArgs)-1 {
			if !c.arguments[argIndex].isVariable() {
				return false
			} else {
				return true
			}
		}

		// this is an optional argument. If we don't get a match, keep trying
		if c.arguments[argIndex].isOptional() {
			if c.arguments[argIndex].represents(rawArgs[rawArgIndex]) {
				rawArgIndex++
				argIndex++
			} else {
				rawArgIndex++
			}
		} else {
			if c.arguments[argIndex].represents(rawArgs[rawArgIndex]) {
				rawArgIndex++
				argIndex++
			} else {
				return false
			}
		}
	}
	return true

}
