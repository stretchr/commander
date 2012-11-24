package commander

import (
	"github.com/stretchrcom/testify/assert"
	"strings"
	"testing"
)

const (
	commandString                       = "create kind=project|account name=(string) [description=(string)...]"
	commandStringOptionalBad            = "create kind=project|account [name=(string)] description=(string)"
	commandStringTwoOptional            = "create kind=project|account name=(string) [description=(string)] [domain=(string)]"
	commandStringTwoOptionalVariable    = "create kind=project|account name=(string) [description=(string)] [domains=(string)...]"
	commandStringTwoOptionalVariableBad = "create kind=project|account name=(string) [description=(string)...] [domains=(string)]"
	rawCommandStringOne                 = "create project stretchr"
	rawCommandStringTwo                 = "create account mat"
)

var (
	rawCommandArrayThree = []string{"create", "project", "stretchr", "Awesome service!"}
	rawCommandArrayFour  = []string{"create", "account", "mat", "Crazy Brit!"}
	rawCommandArrayFive  = []string{"create", "account", "mat", "Crazy Brit!", "localhost"}
	rawCommandArraySix   = []string{"create", "account", "mat", "Crazy Brit!", "localhost", "127.0.0.1", "google.com"}
	rawCommandArraySeven = []string{"create", "account", "mat"}

	commandArray = []string{commandString, commandStringTwoOptional}
)

func HandlerFunc(args map[string]interface{}) {

}

func TestCommand_makeCommand(t *testing.T) {

	c := makeCommand(commandString, HandlerFunc)

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
		_ = makeCommand(commandStringTwoOptionalVariableBad, HandlerFunc)
	})

	assert.Panics(t, func() {
		_ = makeCommand(commandStringOptionalBad, HandlerFunc)
	})

	assert.Panics(t, func() {
		_ = makeCommand(commandString, nil)
	})

}

func TestCommand_Represents(t *testing.T) {

	c := makeCommand(commandString, HandlerFunc)

	assert.True(t, c.represents(strings.Split(rawCommandStringOne, " ")))
	assert.True(t, c.represents(strings.Split(rawCommandStringTwo, " ")))
	assert.True(t, c.represents(rawCommandArrayThree))
	assert.True(t, c.represents(rawCommandArrayFour))

	c = makeCommand(commandStringTwoOptional, HandlerFunc)

	assert.True(t, c.represents(strings.Split(rawCommandStringOne, " ")))
	assert.True(t, c.represents(strings.Split(rawCommandStringTwo, " ")))
	assert.True(t, c.represents(rawCommandArrayThree))
	assert.True(t, c.represents(rawCommandArrayFour))
	assert.True(t, c.represents(rawCommandArrayFive))

	c = makeCommand(commandStringTwoOptionalVariable, HandlerFunc)

	assert.True(t, c.represents(strings.Split(rawCommandStringOne, " ")))
	assert.True(t, c.represents(strings.Split(rawCommandStringTwo, " ")))
	assert.True(t, c.represents(rawCommandArrayThree))
	assert.True(t, c.represents(rawCommandArrayFour))
	assert.True(t, c.represents(rawCommandArrayFive))
	assert.True(t, c.represents(rawCommandArraySix))

}

func TestCommand_IsEqualTo(t *testing.T) {

	for i := 0; i < len(commandArray); i++ {
		for j := 0; j < len(commandArray); j++ {
			a := makeCommand(commandArray[i], HandlerFunc)
			a2 := makeCommand(commandArray[j], HandlerFunc)
			if commandArray[i] == commandArray[j] {
				assert.True(t, a.isEqualTo(a2))
			} else {
				assert.False(t, a.isEqualTo(a2))
			}
		}
	}

}
