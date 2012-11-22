package commander

import (
	"regexp"
	"strings"
)

/*
const PathSegmentRegexString = `(?:[A-Za-z0-9\-._~!$&'()*+,;=:@]|%[0-9A-Fa-f]{2})*`

var PathSegmentRegex = regexp.MustCompile(PathSegmentRegexString)
*/

var (
	literalRegex         = regexp.MustCompile(`^[^=|()\[\]]+$`)
	listRegex            = regexp.MustCompile(`^[^=|()\[\]]+=[^=|()\[\]]+(?:\|[^=|()\[\]]+)+$`)
	captureRegex         = regexp.MustCompile(`^(?P<open>[\[])?(?P<kind>[^=|()\[\]]+)=\((?P<type>[^=|()\[\]]+)\)(?P<variable>\.\.\.)?(?P<close>[\]])?$`)
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
	isOptional bool

	// isVariable is a bool used to determine if this argument is variable
	isVariable bool
}

// MapSubmatchNames creates a map of named matches and their contents
func MapSubmatchNames(submatchNames []string, submatches []string) map[string]string {

	var mappedSubexpNames = make(map[string]string)

	for i := 1; i < len(submatchNames); i++ {
		mappedSubexpNames[submatchNames[i]] = submatches[i]
	}

	return mappedSubexpNames

}

// ContainsString determines if a map[string]string contains string key
func containsKey(stringMap map[string]string, key string) bool {

	if _, ok := stringMap[key]; ok {
		return true
	}
	return false

}

func MakeArgument(rawArg string) *argument {

	a := new(argument)
	a.rawArg = rawArg

	a.parseArgument()

	return a

}

func (a *argument) parseArgument() {

	// TODO: replace string literals with string constants

	switch {
	case literalRegex.MatchString(a.rawArg):
		a.literal = a.rawArg
	case listRegex.MatchString(a.rawArg):
		parts := strings.Split(a.rawArg, "=")
		a.identifier = parts[0]
		a.list = strings.Split(parts[1], "|")
	case captureRegex.MatchString(a.rawArg):
		submatches := captureRegex.FindStringSubmatch(a.rawArg)
		submatchMap := MapSubmatchNames(captureSubmatchNames, submatches)

		a.identifier = submatchMap["kind"]
		a.captureType = submatchMap["type"]

		if containsKey(submatchMap, "open") && containsKey(submatchMap, "close") {
			a.isOptional = true
		}
		if containsKey(submatchMap, "variable") {
			a.isVariable = true
		}
	}

}

func (a *argument) IsLiteral() bool {

	return len(a.literal) != 0

}

func (a *argument) IsList() bool {

	return a.identifier != "" && len(a.list) > 0

}

func (a *argument) IsCapture() bool {

	return a.identifier != "" && a.captureType != ""

}

func (a *argument) IsOptional() bool {

	return a.isOptional

}

func (a *argument) IsVariable() bool {

	return a.isVariable

}
