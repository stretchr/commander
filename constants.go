package commander

const (

	// delimiterArgumentSeparator is the string that separates arguments
	// in a command line.
	delimiterArgumentSeparator string = " "

	// delimiterEquality is the string that indicates an identifier is associated
	// with a capture group or literal
	delimiterEquality string = "="

	// delimiterListItems is the string that separates a group of literals,
	// indicating that it is a list
	delimiterListItems string = "|"
)

const (
	submatchKeyKind     string = "kind"
	submatchKeyType     string = "type"
	submatchKeyOpen     string = "open"
	submatchKeyClose    string = "close"
	submatchKeyVariable string = "variable"
)
