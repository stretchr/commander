Commander - Control your lines
========= 

Commander is a Go package that makes it easy to build and maintain command-line tools and provides
an attractive alternative to the [flag](http://golang.org/pkg/flag/) package.

    package main

    import (
      "github.com/stretchrcom/commander"
    )

    func main() {

      /*
        Use the commander.Go wrapper to initialise and execute your commands
      */
      commander.Go(func(){

        /*
          Map the create command
        */
        commander.Map("create kind=(string) name=(string)", 
          "Creates something",
          "Creates a thing of the specified kind, with the specified name.",
          func(args map[string]interface{}){
      
            // TODO: create something of type args["kind"] called args["name"]
      
          }
        )

      })

    }

The above code will create a tool that supports a single `create` command, with two string arguments.

### Commander vs Flag

Depending on how you would like users to interact with your command-line tool, you should make a choice between
Commander and the built-in `flag` package.

#### Flag

Flag provides traditional interactions where you set parameters by name.  For example;

    mycommand -action=create -name=Mat -age=29
    mycommand -action=update -id=123 -name=Mat

#### Commander

Commander provides a more modern and easy-to-read-and-write alternative.  For example;

    mycommand create Mat 29
    mycommand update 123 Mat

## Get started

  * Check out the [API Documentation](http://godoc.org/github.com/stretchrcom/commander).

## Features

  * Automatic usage help generation
  * Typed arguments
  * Optional arguments
  * Literal (and list literal) arguments