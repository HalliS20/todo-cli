package models

type Colorizer struct {
	Colors     map[string]string
	Commands   map[string]string
	FontStyles map[string]string
}

func NewColorizer() Colorizer {
	return Colorizer{
		Colors: map[string]string{
			"red":        "\033[31m",
			"green":      "\033[32m",
			"thickGreen": "\033[1;32m",
			"yellow":     "\033[33m",
			"blue":       "\033[34m",
			"purple":     "\033[35m",
			"cyan":       "\033[36m",
			"white":      "\033[37m",
			"pink":       "\033[1;31m",
		},
		Commands: map[string]string{
			"reset": "\033[0m",
		},
		FontStyles: map[string]string{
			"bold":      "\033[1m",
			"underline": "\033[4m",
		},
	}
}
