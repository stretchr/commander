package commander

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestGo(t *testing.T) {
	sharedCommander = new(commander)
	incomingArgs = []string{}

	called := false

	Go(func() {
		Map(DefaultCommand, "", "", func(args map[string]interface{}) {
			called = true
		})
	})

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
		assert.Equal(t, args["description"], "Crazy Brit!")
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
