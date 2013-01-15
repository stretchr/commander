package commander

// Go wraps calls to `commander.Map` (which should be placed in the func argument) and
// initializes Commander and executes the commands.
//
// Usage
//
//     commander.Go(func(){
//     
//	     // make calls to commander.Map here
//     
//     })
func Go(mappings func()) {

	// ensure commander is initialized
	initialize()

	// calling the mappings function
	mappings()

	// execute commander
	execute()

}
