package Helper

//Read constant values from config file in future when implemented

// Constants holds the values instead of hard-coding them
type Constants struct {
	MESSAGE_ID_LENGTH int
}

func SetConstants() *Constants {
	return &Constants{
		MESSAGE_ID_LENGTH: 8,
	}
}
