package main

type GridDefinitions struct {
	Width          int
	Height         int
	SleepIteration int
	AnsiEscapeSeq  string
	BrownBlocks    string
	GreenBlocks    string
}

func Setup() GridDefinitions {
	return GridDefinitions{
		Width:          40,
		Height:         20,
		SleepIteration: 200,
		AnsiEscapeSeq:  "\033c\x0c",
		BrownBlocks:    "\xF0\x9F\x9F\xAB",
		GreenBlocks:    "\xF0\x9F\x9F\xA9",
	}
}
