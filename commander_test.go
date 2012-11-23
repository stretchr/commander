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
	args = []string{}

	defaultRun := false

	Map(DefaultCommand, func(map[string]interface{}) {
		defaultRun = true
	})

	Execute()

	assert.True(t, defaultRun)

}
