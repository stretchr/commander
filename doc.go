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

*/
package commander
