package commander

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommander_Map(t *testing.T) {

	sharedCommander = new(commander)

	Map(commandString, "", "", func(map[string]interface{}) {
	})

	assert.Equal(t, len(sharedCommander.commands), 1)

	Map(DefaultCommand, "", "", func(map[string]interface{}) {
	})

	assert.Equal(t, len(sharedCommander.commands), 2)

	assert.Panics(t, func() {
		Map(DefaultCommand, "", "", func(map[string]interface{}) {
		})
	})

	assert.Panics(t, func() {
		Map(commandString, "", "", func(map[string]interface{}) {
		})
	})

}

func TestCommander_execute(t *testing.T) {

	sharedCommander = new(commander)
	incomingArgs = []string{}

	called := false

	Map(DefaultCommand, "", "", func(args map[string]interface{}) {
		called = true
	})

	execute()
	assert.True(t, called)

	called = false
	sharedCommander = new(commander)

	Map(commandString, "", "", func(args map[string]interface{}) {
		called = true
		assert.Equal(t, len(args), 3)
		assert.Equal(t, args["kind"], "account")
		assert.Equal(t, args["name"], "mat")
		assert.Equal(t, args["description"], "Crazy Brit!")
	})

	incomingArgs = rawCommandArrayFour

	execute()
	assert.True(t, called)

	called = false
	sharedCommander = new(commander)

	Map(commandStringTwoOptionalVariable, "", "", func(args map[string]interface{}) {
		called = true

		assert.Equal(t, len(args), 4)
		assert.Equal(t, args["kind"], "account")
		assert.Equal(t, args["name"], "mat")
		assert.Equal(t, args["description"], "Awesome Brit!")
		if assert.Equal(t, len(args["domains"].([]string)), 3) {
			domains := args["domains"].([]string)
			assert.Equal(t, domains[0], "localhost")
			assert.Equal(t, domains[1], "127.0.0.1")
			assert.Equal(t, domains[2], "google.com")
		}
	})

	incomingArgs = rawCommandArraySix

	execute()
	assert.True(t, called)

	called = false
}

func TestCommander_NoOptional(t *testing.T) {

	sharedCommander = new(commander)

	Map(commandStringTwoOptionalVariable, "", "", func(args map[string]interface{}) {
	})

	incomingArgs = rawCommandArraySeven

	assert.NotPanics(t, func() {
		execute()
	})

}

func TestCommander_Real(t *testing.T) {

	initialize()

	Map(DefaultCommand, "", "", func(args map[string]interface{}) {
	})

	Map("test [name=(string)]", "", "",
		func(args map[string]interface{}) {
		})

	Map("install [name=(string)]", "", "",
		func(args map[string]interface{}) {
		})

	Map("vet [name=(string)]", "", "",
		func(args map[string]interface{}) {
		})

	Map("exclude name=(string)", "", "",
		func(args map[string]interface{}) {
		})

	Map("include name=(string)", "", "",
		func(args map[string]interface{}) {
		})

	Map("exclusions", "", "",
		func(args map[string]interface{}) {
		})

	incomingArgs = []string{"test"}
	execute()

}

/*func TestCommander_PrintUsage(t *testing.T) {

	Initialize()

	Map(commandString, "", "", func(args map[string]interface{}) {
	})

	Map(commandStringTwoOptionalVariable, "", "", func(args map[string]interface{}) {
	})

	incomingArgs = []string{"help"}

	assert.NotPanics(t, func() {
		execute()
	})

}*/
/*
func TestCommander_ClosestMatch(t *testing.T) {

	sharedCommander = new(Commander)

	Map(commandString, "", "", func(args map[string]interface{}) {
		t.Error("Shouldn't get here!")
	})

	incomingArgs = []string{"create"}

	execute()

}*/
