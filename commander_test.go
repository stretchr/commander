package commander

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestCommander_Map(t *testing.T) {

	Map(commandString, func(map[string]interface{}) {
	})

	assert.Equal(t, len(sharedCommander.commands), 1)

}
