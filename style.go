/* Right now I only added necessary colors for logger */
package prettylog

type Style struct {
	Disabled bool
	stylmap  map[string]string
}

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

func (s *Style) Get(str, msg string) string {
	if s.Disabled {
		return msg
	}
	return s.color(s.stylmap[str], msg)
}

func (s *Style) GetX(str, msg string) string {
	if s.Disabled {
		return ""
	}
	return s.Get(str, msg)
}
