package commander

import (
	"github.com/stretchr/objx"
	"strings"
)

// Handler is a func type the defines the function signature of the function
// to be called when a command is matched.
type Handler func(args objx.Map)

// command is a type used to create and manage individual command strings
type command struct {
	// definition is the original string contining the command definition
	definition string

	// summary is a string containing a short summary of this command
	summary string

	// description is a string containing a description of this command
	description string

	// handler is the Handler associated with this command
	handler Handler

	// arguments is an array of all the arguments in the command string
	arguments []*argument

	// numOptional contains the number of optional arguments in this command
	numOptional int

	// defaultCommand holds whether this is the default command or not
	defaultCommand bool
}

// makeCommand makes a new Command object and sets it up appropriately
func makeCommand(definition, summary, description string, handler Handler) *command {

	if handler == nil {
		panic("A handler must be defined for each command registered.")
	}

	c := new(command)
	c.definition = definition
	c.handler = handler
	c.defaultCommand = definition == DefaultCommand
	c.description = description
	c.summary = summary

	// make the arguments

	argumentStrings := strings.Split(definition, delimiterArgumentSeparator)
	c.arguments = make([]*argument, len(argumentStrings))
	optionalFound := false

	for argumentIndex, value := range strings.Split(definition, delimiterArgumentSeparator) {
		c.arguments[argumentIndex] = makeArgument(value)
		if !c.arguments[argumentIndex].isOptional() && optionalFound {
			panic("An optional argument may not precede a required argument")
		} else if c.arguments[argumentIndex].isOptional() {
			c.numOptional++
			optionalFound = true
		}
		if c.arguments[argumentIndex].isVariable() {
			if argumentIndex != len(argumentStrings)-1 {
				panic("A variable argument may only appear at the end of a command string")
			}
		}
	}

	return c

}

// represents determines if this command represents the array of arguments
func (c *command) represents(rawArgs []string) (bool, int) {

	argIndex := 0
	for rawArgIndex, _ := range rawArgs {

		if argIndex == len(c.arguments)-1 && rawArgIndex != len(rawArgs)-1 {
			if !c.arguments[argIndex].isVariable() {
				return false, argIndex
			} else {
				argIndex++
				return true, argIndex
			}
		}

		// this is an optional argument. If we don't get a match, keep trying
		if c.arguments[argIndex].isOptional() {
			if c.arguments[argIndex].represents(rawArgs[rawArgIndex]) {
				argIndex++
				rawArgIndex++
			} else {
				rawArgIndex++
			}
		} else {
			if c.arguments[argIndex].represents(rawArgs[rawArgIndex]) {
				argIndex++
				rawArgIndex++
			} else {
				return false, argIndex
			}
		}
	}

	// We've made it this far without an explicit failure, but that doesn't mean
	// every argument has been satisfied. Check our

	if argIndex == len(c.arguments) || argIndex >= len(c.arguments)-c.numOptional {
		return true, argIndex
	}

	return false, argIndex
}

func (c *command) isEqualTo(cmd *command) bool {

	if len(c.arguments) != len(cmd.arguments) {
		return false
	}

	for i := 0; i < len(c.arguments); i++ {
		if !c.arguments[i].isEqualTo(cmd.arguments[i]) {
			return false
		}
	}

	return true

}

func (c *command) isDefaultCommand() bool {
	return c.defaultCommand
}
