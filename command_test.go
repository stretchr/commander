package commander

import (
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	commandString                       = "create kind=project|account name=(string) [description=(string)...]"
	commandStringOptionalBad            = "create kind=project|account [name=(string)] description=(string)"
	commandStringTwoOptional            = "create kind=project|account name=(string) [description=(string)] [domain=(string)]"
	commandStringTwoOptionalVariable    = "create kind=project|account name=(string) [description=(string)] [domains=(string)...]"
	commandStringTwoOptionalVariableBad = "create kind=project|account name=(string) [description=(string)...] [domains=(string)]"
)

var (
	rawCommandArrayOne   = []string{"create", "project", "stretchr"}
	rawCommandArrayTwo   = []string{"create", "account", "mat"}
	rawCommandArrayThree = []string{"create", "project", "stretchr", "Awesome service!"}
	rawCommandArrayFour  = []string{"create", "account", "mat", "Crazy Brit!"}
	rawCommandArrayFive  = []string{"create", "account", "mat", "Crazy Brit!", "localhost"}
	rawCommandArraySix   = []string{"create", "account", "mat", "Awesome Brit!", "localhost", "127.0.0.1", "google.com"}
	rawCommandArraySeven = []string{"create", "account", "mat"}

	cmdArray = []string{commandString, commandStringTwoOptional}
)

func HandlerFunc(args objx.Map) {

}

func TestCommand_makeCommand(t *testing.T) {

	c := makeCommand(commandString, "", "", HandlerFunc)

	if assert.NotNil(t, c) {
		assert.Equal(t, c.definition, commandString)
		assert.Equal(t, len(c.arguments), 4)
		assert.Equal(t, c.arguments[0].literal, "create")
		assert.Equal(t, c.arguments[1].identifier, "kind")
		assert.Equal(t, c.arguments[1].list[0], "project")
		assert.Equal(t, c.arguments[1].list[1], "account")
		assert.Equal(t, c.arguments[2].identifier, "name")
		assert.Equal(t, c.arguments[2].captureType, "string")
		assert.Equal(t, c.arguments[3].identifier, "description")
		assert.Equal(t, c.arguments[3].captureType, "string")
		assert.True(t, c.arguments[3].isOptional())
		assert.True(t, c.arguments[3].isVariable())
	}

	assert.Panics(t, func() {
		_ = makeCommand(commandStringTwoOptionalVariableBad, "", "", HandlerFunc)
	})

	assert.Panics(t, func() {
		_ = makeCommand(commandStringOptionalBad, "", "", HandlerFunc)
	})

	assert.Panics(t, func() {
		_ = makeCommand(commandString, "", "", nil)
	})

}

func repBool(c *command, def []string) bool {
	represents, _ := c.represents(def)
	return represents
}

func TestCommand_Represents(t *testing.T) {

	c := makeCommand(commandString, "", "", HandlerFunc)

	assert.True(t, repBool(c, rawCommandArrayOne))
	assert.True(t, repBool(c, rawCommandArrayTwo))
	assert.True(t, repBool(c, rawCommandArrayThree))
	assert.True(t, repBool(c, rawCommandArrayFour))

	c = makeCommand(commandStringTwoOptional, "", "", HandlerFunc)

	assert.True(t, repBool(c, rawCommandArrayOne))
	assert.True(t, repBool(c, rawCommandArrayTwo))
	assert.True(t, repBool(c, rawCommandArrayThree))
	assert.True(t, repBool(c, rawCommandArrayFour))
	assert.True(t, repBool(c, rawCommandArrayFive))

	c = makeCommand(commandStringTwoOptionalVariable, "", "", HandlerFunc)

	assert.True(t, repBool(c, rawCommandArrayOne))
	assert.True(t, repBool(c, rawCommandArrayTwo))
	assert.True(t, repBool(c, rawCommandArrayThree))
	assert.True(t, repBool(c, rawCommandArrayFour))
	assert.True(t, repBool(c, rawCommandArrayFive))
	assert.True(t, repBool(c, rawCommandArraySix))

}

func TestCommand_IsEqualTo(t *testing.T) {

	for i := 0; i < len(cmdArray); i++ {
		for j := 0; j < len(cmdArray); j++ {
			a := makeCommand(cmdArray[i], "", "", HandlerFunc)
			a2 := makeCommand(cmdArray[j], "", "", HandlerFunc)
			if cmdArray[i] == cmdArray[j] {
				assert.True(t, a.isEqualTo(a2))
			} else {
				assert.False(t, a.isEqualTo(a2))
			}
		}
	}

}
