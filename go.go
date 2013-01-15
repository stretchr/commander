package commander

func Go(mappings func()) {

	// ensure commander is initialized
	Initialize()

	// calling the mappings function
	mappings()

	// execute commander
	Execute()

}
