package commander

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
const PathSegmentRegexString = `(?:[A-Za-z0-9\-._~!$&'()*+,;=:@]|%[0-9A-Fa-f]{2})*`

var PathSegmentRegex = regexp.MustCompile(PathSegmentRegexString)
*/

var (
	// literalRegex represents the regexp object for literals.
	literalRegex = regexp.MustCompile(`^[^=|()\[\]]+$`)

	// listRegex represents the regexp for lists
	listRegex = regexp.MustCompile(`^[^=|()\[\]]+=[^=|()\[\]]+(?:\|[^=|()\[\]]+)+$`)

	// captureRegex represents the regexp for captures.
	captureRegex = regexp.MustCompile(fmt.Sprintf(`^(?P<%s>[\[])?(?P<%s>[^=|()\[\]]+)=\((?P<%s>[^=|()\[\]]+)\)(?P<%s>\.\.\.)?(?P<%s>[\]])?$`,
		submatchKeyOpen, submatchKeyKind, submatchKeyType, submatchKeyVariable, submatchKeyClose))
	// captureSubmatchNames represents the regexp for capture sub matches.
	captureSubmatchNames = captureRegex.SubexpNames()
)

type argument struct {
	// rawArg is a string containing the argument in its raw form
	rawArg string

	// literal is a string containing the text of the literal argument
	literal string

	// list is an array containing each of the arguments in a list
	list []string

	// identifier is a string containing the identifier of the argument
	identifier string

	// captureType is a string containing the capture type of the argument
	captureType string

	// isOptional is a bool used to determine if this argument is optional
	optional bool

	// isVariable is a bool used to determine if this argument is variable
	variable bool
}

// containsKey determines if a map[string]string contains string key
func containsKey(stringMap map[string]string, key string) bool {

	if _, ok := stringMap[key]; ok {
		return true
	}
	return false

}

// constainsString determines if a []string contains string
func containsString(stringSlice []string, contains string) bool {

	for _, value := range stringSlice {
		if value == contains {
			return true
		}
	}
	return false

}

// slicesAreEqual determines if two string slices are equal
func slicesAreEqual(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

// canCastToType determines if the cmdArg can be cast to a given type
func canCastToType(cmdArg, castType string) bool {
	if castToType(cmdArg, castType) == nil {
		return false
	}
	return true
}

func castToType(cmdArg, castType string) interface{} {

	switch castType {
	case "string":
		return cmdArg
	case "int":
		if value, err := strconv.ParseInt(cmdArg, 10, 0); err != nil {
			return nil
		} else {
			return value
		}
	case "int64":
		if value, err := strconv.ParseInt(cmdArg, 10, 64); err != nil {
			return nil
		} else {
			return value
		}
	case "uint":
		if value, err := strconv.ParseUint(cmdArg, 10, 0); err != nil {
			return nil
		} else {
			return value
		}
	case "uint64":
		if value, err := strconv.ParseUint(cmdArg, 10, 64); err != nil {
			return nil
		} else {
			return value
		}
	case "bool":
		if value, err := strconv.ParseBool(cmdArg); err != nil {
			return nil
		} else {
			return value
		}
	case "time":
		if value, err := time.Parse(time.ANSIC, cmdArg); err == nil {
			return value
		}
		if value, err := time.Parse(time.UnixDate, cmdArg); err == nil {
			return value
		}
		if value, err := time.Parse(time.RubyDate, cmdArg); err == nil {
			return value
		}
		if value, err := time.Parse(time.RFC822, cmdArg); err == nil {
			return value
		}
		if value, err := time.Parse(time.RFC822Z, cmdArg); err == nil {
			return value
		}
		if value, err := time.Parse(time.RFC850, cmdArg); err == nil {
			return value
		}
		if value, err := time.Parse(time.RFC1123, cmdArg); err == nil {
			return value
		}
		if value, err := time.Parse(time.RFC1123Z, cmdArg); err == nil {
			return value
		}
		if value, err := time.Parse(time.RFC3339, cmdArg); err == nil {
			return value
		}
		if value, err := time.Parse(time.RFC3339Nano, cmdArg); err == nil {
			return value
		}
		if value, err := time.Parse(time.Kitchen, cmdArg); err == nil {
			return value
		}
		return nil
	}

	return nil

}

// mapSubmatchNames creates a map of named matches and their contents
func mapSubmatchNames(submatchNames []string, submatches []string) map[string]string {

	var mappedSubexpNames = make(map[string]string)

	for i := 1; i < len(submatchNames); i++ {
		if submatches[i] != "" {
			mappedSubexpNames[submatchNames[i]] = submatches[i]
		}
	}

	return mappedSubexpNames

}

func (a *argument) parseArgument() {

	switch {
	case literalRegex.MatchString(a.rawArg):
		a.literal = a.rawArg
	case listRegex.MatchString(a.rawArg):
		parts := strings.Split(a.rawArg, delimiterEquality)
		a.identifier = parts[0]
		a.list = strings.Split(parts[1], delimiterListItems)
	case captureRegex.MatchString(a.rawArg):
		submatches := captureRegex.FindStringSubmatch(a.rawArg)
		submatchMap := mapSubmatchNames(captureSubmatchNames, submatches)

		a.identifier = submatchMap[submatchKeyKind]
		a.captureType = submatchMap[submatchKeyType]

		if containsKey(submatchMap, submatchKeyOpen) && containsKey(submatchMap, submatchKeyClose) {
			a.optional = true
		}
		if containsKey(submatchMap, submatchKeyVariable) {
			a.variable = true
		}
	}

}

func makeArgument(rawArg string) *argument {

	a := new(argument)
	a.rawArg = rawArg

	a.parseArgument()

	return a

}

// represents determines if this argument represents the cmdArg string
func (a *argument) represents(cmdArg string) bool {

	switch {
	case a.isLiteral() && a.literal == cmdArg:
		return true
	case a.isList() && containsString(a.list, cmdArg):
		return true
	case a.isCapture() && canCastToType(cmdArg, a.captureType):
		return true
	}

	return false

}

func (a *argument) isLiteral() bool {

	return a.literal != ""

}

func (a *argument) isList() bool {

	return a.identifier != "" && len(a.list) > 0

}

func (a *argument) isCapture() bool {

	return a.identifier != "" && a.captureType != ""

}

func (a *argument) isOptional() bool {

	return a.optional

}

func (a *argument) isVariable() bool {

	return a.variable

}

func (a *argument) isEqualTo(arg *argument) bool {

	switch {
	case a.isLiteral() && arg.isLiteral():
		return a.literal == arg.literal
	case a.isList() && arg.isList():
		return slicesAreEqual(a.list, arg.list)
	case a.isCapture() && arg.isCapture():
		return a.captureType == arg.captureType &&
			a.identifier == arg.identifier
	}
	return false

}
