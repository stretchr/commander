package commander

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

const (
	argLiteral                     = "create"
	argList                        = "kind=project|account"
	argListLong                    = "kind=project|account|user|admin"
	argCaptureType                 = "name=(string)"
	argOptionalCaptureType         = "[description=(string)]"
	argVariableCaptureType         = "initialUsers=(string)..."
	argOptionalVariableCaptureType = "[initialUsers=(string)...]"
)

func TestArgument_MakeArgument(t *testing.T) {

	a := MakeArgument("rawArg")

	if assert.NotNil(t, a) {
		assert.Equal(t, a.rawArg, "rawArg")
	}

}

func TestArgument_ParseArgument(t *testing.T) {

	a := MakeArgument(argLiteral)

	assert.Equal(t, a.literal, argLiteral)

	a = MakeArgument(argList)

	assert.Equal(t, a.identifier, "kind")
	if assert.Equal(t, len(a.list), 2) {
		assert.Equal(t, a.list[0], "project")
		assert.Equal(t, a.list[1], "account")
	}

	a = MakeArgument(argListLong)
	assert.Equal(t, a.identifier, "kind")
	if assert.Equal(t, len(a.list), 4) {
		assert.Equal(t, a.list[0], "project")
		assert.Equal(t, a.list[1], "account")
		assert.Equal(t, a.list[2], "user")
		assert.Equal(t, a.list[3], "admin")
	}

	a = MakeArgument(argCaptureType)

	assert.Equal(t, a.identifier, "name")
	assert.Equal(t, a.captureType, "string")

	a = MakeArgument(argOptionalCaptureType)

	assert.Equal(t, a.identifier, "description")
	assert.Equal(t, a.captureType, "string")

	a = MakeArgument(argVariableCaptureType)

	assert.Equal(t, a.identifier, "initialUsers")
	assert.Equal(t, a.captureType, "string")

	a = MakeArgument(argOptionalVariableCaptureType)

	assert.Equal(t, a.identifier, "initialUsers")
	assert.Equal(t, a.captureType, "string")

}

func TestArgument_IsLiteral(t *testing.T) {

	a := MakeArgument(argLiteral)

	assert.True(t, a.IsLiteral())

}

func TestArgument_IsList(t *testing.T) {

	a := MakeArgument(argList)

	assert.True(t, a.IsList())

}

func TestArgument_IsCapture(t *testing.T) {

	a := MakeArgument(argCaptureType)

	assert.True(t, a.IsCapture())

	a = MakeArgument(argOptionalCaptureType)

	assert.True(t, a.IsCapture())

	a = MakeArgument(argVariableCaptureType)

	assert.True(t, a.IsCapture())

	a = MakeArgument(argOptionalVariableCaptureType)

	assert.True(t, a.IsCapture())

}

func TestArgument_IsOptional(t *testing.T) {

	a := MakeArgument(argOptionalCaptureType)

	assert.True(t, a.IsOptional())

	a = MakeArgument(argOptionalVariableCaptureType)

	assert.True(t, a.IsOptional())
}

func TestArgument_IsVariable(t *testing.T) {

	a := MakeArgument(argVariableCaptureType)

	assert.True(t, a.IsVariable())

	a = MakeArgument(argOptionalVariableCaptureType)

	assert.True(t, a.IsVariable())

}
