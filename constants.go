package commander

const (

	// DelimiterArgumentSeparator is the string that separates arguments
	// in a command line.
	DelimiterArgumentSeparator string = " "

	DelimiterEquality string = "="

	DelimiterListItems string = "|"
)

// TODO: once the SubmatchKey* constants are used in the regex, simplify them
// by making them one character?
const (
	SubmatchKeyKind     string = "kind"
	SubmatchKeyType     string = "type"
	SubmatchKeyOpen     string = "open"
	SubmatchKeyClose    string = "close"
	SubmatchKeyVariable string = "variable"
)
