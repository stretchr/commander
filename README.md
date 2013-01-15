Commander - Control your lines
=========

Commander is a Go package that makes it easy to build and maintain command-line tools and provides
an attractive alternative to the [flag](http://golang.org/pkg/flag/) package.

    package main

    import (
      "github.com/stretchrcom/commander"
    )

    func main() {

      commander.Go(func(){

        commander.Map("create kind=(string) name=(string)", 
                      "Creates something",
                      "Creates a thing of the specified kind, with the specified name.",
                      func(args map[string]interface{}){
                 
                        // TODO: create something of type args["kind"] called args["name"]
              
                      }
        )

      })

    }

  * Check out the [API Documentation](http://godoc.org/github.com/stretchrcom/commander).
