package colorizer

type Colorizer struct {
	Colors     map[Colors]string
	Commands   map[Commands]string
	FontStyles map[Styles]string
}

func NewColorizer() Colorizer {
	return Colorizer{
		Colors: map[Colors]string{
			Red:        "\033[31m",
			Green:      "\033[32m",
			ThickGreen: "\033[1;32m",
			Yellow:     "\033[33m",
			Blue:       "\033[34m",
			Purple:     "\033[35m",
			Cyan:       "\033[36m",
			White:      "\033[37m",
			Pink:       "\033[1;31m",
		},
		Commands: map[Commands]string{
			Reset: "\033[0m",
		},
		FontStyles: map[Styles]string{
			Bold:      "\033[1m",
			Underline: "\033[4m",
			Italic:    "\033[5m",
		},
	}
}
