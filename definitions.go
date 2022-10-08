package main

type GridDefinitions struct {
	Width          int
	Height         int
	SleepIteration int
	AnsiEscapeSeq  string
	RedBlocks      string
	BlueBlocks     string
}

func Setup() GridDefinitions {
	return GridDefinitions{
		Width:          40,
		Height:         20,
		SleepIteration: 200,
		AnsiEscapeSeq:  "\033c\x0c",
		RedBlocks:      "\xF0\x9F\x9F\xA5",
		BlueBlocks:     "\xF0\x9F\x9F\xA6",
	}
}
