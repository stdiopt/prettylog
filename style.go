package prettylog

import "fmt"

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
			"Counter":  "\033[37m",
			"Message":  "\033[37m",
			"Prefix":   "\033[33m",
			"Time":     "\033[34m",
			"Duration": "\033[90m",
			"File":     "\033[30m",
		},
	}
}

func (s *Style) color(code string, str interface{}) string {
	return fmt.Sprintf("%s%v\033[0m", code, str)
}

//Get a msg string with style if enabled
func (s *Style) Get(str string, msg interface{}) string {
	if s.Disabled {
		return fmt.Sprint(msg)
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
