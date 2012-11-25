package commander

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestCommander_Map(t *testing.T) {

	sharedCommander = new(Commander)

	Map(commandString, func(map[string]interface{}) {
	})

	assert.Equal(t, len(sharedCommander.commands), 1)

	Map(DefaultCommand, func(map[string]interface{}) {
	})

	assert.Equal(t, len(sharedCommander.commands), 2)

	assert.Panics(t, func() {
		Map(DefaultCommand, func(map[string]interface{}) {
		})
	})

	assert.Panics(t, func() {
		Map(commandString, func(map[string]interface{}) {
		})
	})

}

func TestCommander_Execute(t *testing.T) {

	sharedCommander = new(Commander)
	incomingArgs = []string{}

	called := false

	Map(DefaultCommand, func(args map[string]interface{}) {
		called = true
	})

	Execute()
	assert.True(t, called)

	called = false
	sharedCommander = new(Commander)

	Map(commandString, func(args map[string]interface{}) {
		called = true
		assert.Equal(t, len(args), 3)
		assert.Equal(t, args["kind"], "account")
		assert.Equal(t, args["name"], "mat")
		assert.Equal(t, args["description"], "Crazy Brit!")
	})

	incomingArgs = rawCommandArrayFour

	Execute()
	assert.True(t, called)

	called = false
	sharedCommander = new(Commander)

	Map(commandStringTwoOptionalVariable, func(args map[string]interface{}) {
		called = true

		assert.Equal(t, len(args), 4)
		assert.Equal(t, args["kind"], "account")
		assert.Equal(t, args["name"], "mat")
		assert.Equal(t, args["description"], "Crazy Brit!")
		if assert.Equal(t, len(args["domains"].([]string)), 3) {
			domains := args["domains"].([]string)
			assert.Equal(t, domains[0], "localhost")
			assert.Equal(t, domains[1], "127.0.0.1")
			assert.Equal(t, domains[2], "google.com")
		}
	})

	incomingArgs = rawCommandArraySix

	Execute()
	assert.True(t, called)

	called = false
}

func TestCommander_NoOptional(t *testing.T) {

	sharedCommander = new(Commander)

	Map(commandStringTwoOptionalVariable, func(args map[string]interface{}) {
	})

	incomingArgs = rawCommandArraySeven

	assert.NotPanics(t, func() {
		Execute()
	})

}

/*func TestCommander_PrintUsage(t *testing.T) {

	Initialize()

	Map(commandString, func(args map[string]interface{}) {
	})

	Map(commandStringTwoOptionalVariable, func(args map[string]interface{}) {
	})

	incomingArgs = []string{"help"}

	assert.NotPanics(t, func() {
		Execute()
	})

}*/
/*
func TestCommander_ClosestMatch(t *testing.T) {

	sharedCommander = new(Commander)

	Map(commandString, func(args map[string]interface{}) {
		t.Error("Shouldn't get here!")
	})

	incomingArgs = []string{"create"}

	Execute()

}*/
