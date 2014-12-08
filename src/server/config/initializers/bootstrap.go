package initializers

// Put all Initialization Functions Here

func Init(debug bool) {
	// Connect to & Migrate the database
	InitDB(debug)

	go UpdateScoreLoop()
}
