package commander

const (

	// DelimiterArgumentSeparator is the string that separates arguments
	// in a command line.
	DelimiterArgumentSeparator string = " "

	// DelimiterEquality is the string that indicates an identifier is associated
	// with a capture group or literal
	DelimiterEquality string = "="

	// DelimiterListItems is the string that separates a group of literals,
	// indicating that it is a list
	DelimiterListItems string = "|"
)

const (
	SubmatchKeyKind     string = "kind"
	SubmatchKeyType     string = "type"
	SubmatchKeyOpen     string = "open"
	SubmatchKeyClose    string = "close"
	SubmatchKeyVariable string = "variable"
)
