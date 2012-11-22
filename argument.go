package commander

import (
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
	// TODO: Is there an easy way to use the constants in this string too?
	captureRegex = regexp.MustCompile(`^(?P<open>[\[])?(?P<kind>[^=|()\[\]]+)=\((?P<type>[^=|()\[\]]+)\)(?P<variable>\.\.\.)?(?P<close>[\]])?$`)

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

// canCastToType determines if the cmdArg can be cast to a given type
func canCastToType(cmdArg string, castType string) bool {

	switch castType {
	case "string":
		return true
	case "int":
		if _, err := strconv.ParseInt(cmdArg, 10, 0); err != nil {
			return false
		}
		return true
	case "int64":
		if _, err := strconv.ParseInt(cmdArg, 10, 64); err != nil {
			return false
		}
		return true
	case "uint":
		if _, err := strconv.ParseUint(cmdArg, 10, 0); err != nil {
			return false
		}
		return true
	case "uint64":
		if _, err := strconv.ParseUint(cmdArg, 10, 64); err != nil {
			return false
		}
		return true
	case "bool":
		if _, err := strconv.ParseBool(cmdArg); err != nil {
			return false
		}
		return true
	case "time":
		if _, err := time.Parse(time.ANSIC, cmdArg); err == nil {
			return true
		}
		if _, err := time.Parse(time.UnixDate, cmdArg); err == nil {
			return true
		}
		if _, err := time.Parse(time.RubyDate, cmdArg); err == nil {
			return true
		}
		if _, err := time.Parse(time.RFC822, cmdArg); err == nil {
			return true
		}
		if _, err := time.Parse(time.RFC822Z, cmdArg); err == nil {
			return true
		}
		if _, err := time.Parse(time.RFC850, cmdArg); err == nil {
			return true
		}
		if _, err := time.Parse(time.RFC1123, cmdArg); err == nil {
			return true
		}
		if _, err := time.Parse(time.RFC1123Z, cmdArg); err == nil {
			return true
		}
		if _, err := time.Parse(time.RFC3339, cmdArg); err == nil {
			return true
		}
		if _, err := time.Parse(time.RFC3339Nano, cmdArg); err == nil {
			return true
		}
		if _, err := time.Parse(time.Kitchen, cmdArg); err == nil {
			return true
		}
		return false
	}

	return false

}

// mapSubmatchNames creates a map of named matches and their contents
func mapSubmatchNames(submatchNames []string, submatches []string) map[string]string {

	var mappedSubexpNames = make(map[string]string)

	for i := 1; i < len(submatchNames); i++ {
		mappedSubexpNames[submatchNames[i]] = submatches[i]
	}

	return mappedSubexpNames

}

func (a *argument) parseArgument() {

	// TODO: replace string literals with string constants

	switch {
	case literalRegex.MatchString(a.rawArg):
		a.literal = a.rawArg
	case listRegex.MatchString(a.rawArg):
		parts := strings.Split(a.rawArg, DelimiterEquality)
		a.identifier = parts[0]
		a.list = strings.Split(parts[1], DelimiterListItems)
	case captureRegex.MatchString(a.rawArg):
		submatches := captureRegex.FindStringSubmatch(a.rawArg)
		submatchMap := mapSubmatchNames(captureSubmatchNames, submatches)

		a.identifier = submatchMap[SubmatchKeyKind]
		a.captureType = submatchMap[SubmatchKeyType]

		if containsKey(submatchMap, "open") && containsKey(submatchMap, "close") {
			a.optional = true
		}
		if containsKey(submatchMap, "variable") {
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

	return len(a.literal) != 0

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
