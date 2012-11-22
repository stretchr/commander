package commander

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestCommand_makeCommand(t *testing.T) {

	c := makeCommand("create kind=project|account name=(string) [description=(string)...]")

	if assert.NotNil(t, c) {
		assert.Equal(t, c.definition, "create kind=project|account name=(string) [description=(string)...]")
	}

	assert.Equal(t, len(c.arguments), 4)
	assert.Equal(t, c.arguments[0].literal, "create")
	assert.Equal(t, c.arguments[1].identifier, "kind")
	assert.Equal(t, c.arguments[1].list[0], "project")
	assert.Equal(t, c.arguments[1].list[1], "account")
	assert.Equal(t, c.arguments[2].identifier, "name")
	assert.Equal(t, c.arguments[2].captureType, "string")
	assert.Equal(t, c.arguments[3].identifier, "description")
	assert.Equal(t, c.arguments[3].captureType, "string")
	assert.True(t, c.arguments[3].IsOptional())
	assert.True(t, c.arguments[3].IsVariable())
}
