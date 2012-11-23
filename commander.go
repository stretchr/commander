package commander

import (
	"sync"
)

// Commander provides methods and functionality to create a command line
// interface quickly and easily.
type Commander struct {
	// commands contains all the mapped commands
	commands []*command
}

// initOnce is used to guarantee that the sharedCommander is initialized only once.
var initOnce sync.Once

// sharedCommander is the shared instance of the Commander type
var sharedCommander *Commander

func Map(definition string, handler Handler) {

	initOnce.Do(func() {
		sharedCommander = new(Commander)
	})

	newCommand := makeCommand(definition, handler)

	sharedCommander.commands = append(sharedCommander.commands, newCommand)

}
