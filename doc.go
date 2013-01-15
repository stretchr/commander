/*
Commander - Control your lines

Commander is a Go package that makes it easy to build and maintain command-line tools.

Usage

Commander works by mapping handler funcs to command signatures, much like matching URL routes
to controllers for websites.

In the `main` func (in the `main` package) you call the `commander.Go` func like this:

    package main

    import (
      "github.com/stretchrcom/commander"
    ) 

    func main() {

      // wrap all commander.Map calls in the Go call...
      commander.Go(func(){

        // map something
        commander.Map({definition}, {summary}, {description}, {handler})

      })

    }

{definition} - The definition is a string that describes the mapping of the command.

{summary} - The summary is a tiny overview of what the command does.

{description} - The description is a long description of what the command does.

{handler} - Handler is the func that will be called when the user initiates this command.

Handler Func 

The Handler func is a normal func that takes a `map[string]interface{}` as its only argument, and
returns nothing.

The argument will contain a map of the arguments described in the definition.

Definitions

A definition is a string that describes the command, including arguments, so that Commander knows when to
called the associated handler func.

Literal

A literal is any string that is not contained inside ( ) and is not followed by =

A literal will be literally matched when parsing the command.

Identifier

An identifier is a string that is followed by = equals character.

An identifier becomes the key for this argument in the map passed to your handler function.

List

A list is a group of literals separated by | pipe character.

Capture Type

A capture type is a string contained in ( ) parenthesis.

The string inside the ( ) defines what type is required. If the argument cannot be represented by this type, an error will occur.

Optional Argument

An optional argument is surrounded by [ ] square brackets.

Variable Arguments

A variable argument is defined by placing "..." (three period characters) after a capture type
It is only valid as the last argument in the command string.

Examples

If we wanted to provide a command-line tool that allowed you to create two types of objects, we
could use the following mapping definition:

    create kind=project|account name=(string) [description=(string)]

If we compile this into a Go command called 'please', then in our Terminal all of these lines
would hit this mapping:

    please create project ProjectName ProjectDesc
    please create project ProjectName
    please create account MyAccount

The associated handler func would be called, and the `args` map would contain the appropriate 
values.  For example, for:

    please create project ProjectName ProjectDesc

The `args` map would contain:

    args["kind"]        == "project"
    args["name"]        == "ProjectName"
    args["description"] == "ProjectDesc"

NOTE: `please` is the actual command name, `create` is what tells Commander to use the specified
handler func, and anything following that are the arguments relevant to that command.

Because we used `kind=project|account`, any other value will NOT match that command.  So these
calls would NOT hit the same handler func:

    please create logs mylogname

In order to provide that functionality, another Map call would have to be made.

*/
package commander
