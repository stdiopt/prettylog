package prettylog

// Style structure with disable option
type Style struct {
	Disabled bool
	stylmap  map[string]string
}

// NewStyle fetches new Style with default values
func NewStyle() *Style {
	return &Style{
		Disabled: false,
		stylmap: map[string]string{
			"Message":  "\033[37m",
			"Prefix":   "\033[33m",
			"Time":     "\033[34m",
			"Duration": "\033[90m",
			"File":     "\033[30m",
		},
	}
}

func (s *Style) color(code string, str string) string {
	return code + str + "\033[0m"
}

//Get a msg string with style if enabled
func (s *Style) Get(str, msg string) string {
	if s.Disabled {
		return msg
	}
	return s.color(s.stylmap[str], msg)
}

//GetX gets a msg with style or none if disabled
func (s *Style) GetX(str, msg string) string {
	if s.Disabled {
		return ""
	}
	return s.Get(str, msg)
}
