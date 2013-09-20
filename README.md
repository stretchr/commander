Commander - Drop and Give Me 20
=========

Commander is a Go package that makes it easy to build and maintain command-line tools and provides
an attractive alternative to the [flag](http://golang.org/pkg/flag/) package.

    package main

    import (
      "github.com/stretchr/commander"
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
          func(args objx.Map){

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

  * Check out the [API Documentation](http://godoc.org/github.com/stretchr/commander).

## Features

  * Automatic usage help generation
  * Typed arguments
  * Optional arguments
  * Literal (and list literal) arguments



------

Contributing
============

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include steps to reproduce the issue so we can see it on our end also!


Licence
=======
Copyright (c) 2012 Mat Ryer and Tyler Bunnell

Please consider promoting this project if you find it useful.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
