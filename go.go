package commander

func Go(mappings func()) {

	// ensure commander is initialized
	initialize()

	// calling the mappings function
	mappings()

	// execute commander
	execute()

}
