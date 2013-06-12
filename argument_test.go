package commander

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	argLiteral                     = "create"
	argList                        = "kind=project|account"
	argListLong                    = "kind=project|account|user|admin"
	argCaptureType                 = "name=(string)"
	argCaptureTypeInt              = "num=(int)"
	argCaptureTypeInt64            = "num=(int64)"
	argCaptureTypeUint             = "num=(uint)"
	argCaptureTypeUint64           = "num=(uint64)"
	argCaptureTypeBool             = "enabled=(bool)"
	argCaptureTypeTime             = "date=(time)"
	argOptionalCaptureType         = "[description=(string)]"
	argVariableCaptureType         = "initialUsers=(string)..."
	argOptionalVariableCaptureType = "[initialUsers=(string)...]"
)

var argArray = []string{argLiteral, argList, argListLong, argCaptureType,
	argCaptureTypeInt, argCaptureTypeInt64, argCaptureTypeUint,
	argCaptureTypeUint64, argCaptureTypeBool, argCaptureTypeTime}

func TestArgument_MakeArgument(t *testing.T) {

	a := makeArgument("rawArg")

	if assert.NotNil(t, a) {
		assert.Equal(t, a.rawArg, "rawArg")
	}

}

func TestArgument_Represents(t *testing.T) {

	a := makeArgument(argLiteral)

	assert.True(t, a.represents("create"))
	assert.False(t, a.represents("delete"))

	a = makeArgument(argList)

	assert.True(t, a.represents("project"))
	assert.True(t, a.represents("account"))
	assert.False(t, a.represents("thanksgiving"))

	a = makeArgument(argListLong)
	assert.True(t, a.represents("project"))
	assert.True(t, a.represents("account"))
	assert.True(t, a.represents("user"))
	assert.True(t, a.represents("admin"))
	assert.False(t, a.represents("thanksgiving"))

	a = makeArgument(argCaptureType)
	assert.True(t, a.represents("Anything at all!"))

	a = makeArgument(argCaptureTypeInt)
	assert.True(t, a.represents("-123"))
	assert.True(t, a.represents("123"))
	assert.False(t, a.represents("Not an integer"))

	a = makeArgument(argCaptureTypeInt64)
	assert.True(t, a.represents("-123"))
	assert.True(t, a.represents("123"))
	assert.False(t, a.represents("Not an integer"))

	a = makeArgument(argCaptureTypeUint)
	assert.True(t, a.represents("123"))
	assert.False(t, a.represents("-123"))
	assert.False(t, a.represents("Not an integer"))

	a = makeArgument(argCaptureTypeUint64)
	assert.True(t, a.represents("123"))
	assert.False(t, a.represents("-123"))
	assert.False(t, a.represents("Not an integer"))

	a = makeArgument(argCaptureTypeBool)
	assert.True(t, a.represents("true"))
	assert.False(t, a.represents("unicorn"))

	a = makeArgument(argCaptureTypeTime)
	assert.True(t, a.represents("3:04PM"))
	assert.False(t, a.represents("three thirty pm"))

}

func TestArgument_ParseArgument(t *testing.T) {

	a := makeArgument(argLiteral)

	assert.Equal(t, a.literal, argLiteral)

	a = makeArgument(argList)

	assert.Equal(t, a.identifier, "kind")
	if assert.Equal(t, len(a.list), 2) {
		assert.Equal(t, a.list[0], "project")
		assert.Equal(t, a.list[1], "account")
	}

	a = makeArgument(argListLong)
	assert.Equal(t, a.identifier, "kind")
	if assert.Equal(t, len(a.list), 4) {
		assert.Equal(t, a.list[0], "project")
		assert.Equal(t, a.list[1], "account")
		assert.Equal(t, a.list[2], "user")
		assert.Equal(t, a.list[3], "admin")
	}

	a = makeArgument(argCaptureType)

	assert.Equal(t, a.identifier, "name")
	assert.Equal(t, a.captureType, "string")

	a = makeArgument(argOptionalCaptureType)

	assert.Equal(t, a.identifier, "description")
	assert.Equal(t, a.captureType, "string")

	a = makeArgument(argVariableCaptureType)

	assert.Equal(t, a.identifier, "initialUsers")
	assert.Equal(t, a.captureType, "string")

	a = makeArgument(argOptionalVariableCaptureType)

	assert.Equal(t, a.identifier, "initialUsers")
	assert.Equal(t, a.captureType, "string")

}

func TestArgument_isLiteral(t *testing.T) {

	a := makeArgument(argLiteral)

	assert.True(t, a.isLiteral())

}

func TestArgument_isList(t *testing.T) {

	a := makeArgument(argList)

	assert.True(t, a.isList())

}

func TestArgument_isCapture(t *testing.T) {

	a := makeArgument(argCaptureType)

	assert.True(t, a.isCapture())

	a = makeArgument(argOptionalCaptureType)

	assert.True(t, a.isCapture())

	a = makeArgument(argVariableCaptureType)

	assert.True(t, a.isCapture())

	a = makeArgument(argOptionalVariableCaptureType)

	assert.True(t, a.isCapture())

}

func TestArgument_isOptional(t *testing.T) {

	a := makeArgument(argOptionalCaptureType)

	assert.True(t, a.isOptional())

	a = makeArgument(argOptionalVariableCaptureType)

	assert.True(t, a.isOptional())
}

func TestArgument_isVariable(t *testing.T) {

	a := makeArgument(argVariableCaptureType)

	assert.True(t, a.isVariable())

	a = makeArgument(argOptionalVariableCaptureType)

	assert.True(t, a.isVariable())

}

func TestArgument_isEqualTo(t *testing.T) {

	for i := 0; i < len(argArray); i++ {
		for j := 0; j < len(argArray); j++ {
			a := makeArgument(argArray[i])
			a2 := makeArgument(argArray[j])
			if argArray[i] == argArray[j] {
				assert.True(t, a.isEqualTo(a2))
			} else {
				assert.False(t, a.isEqualTo(a2))
			}
		}
	}
}
