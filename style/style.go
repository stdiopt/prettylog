package style

import "fmt"

// Style options for specific style
type Style struct {
	Prefix       string
	Suffix       string
	IncrementPad bool
	Pad          int
}

//Get styled message
func (s *Style) Get(msg interface{}) string {
	m := fmt.Sprintf("%v", msg)

	if s.IncrementPad {
		if t := len(m); t > s.Pad {
			s.Pad = t
		}
		m = fmt.Sprintf("%*s", s.Pad, m)
	}

	return fmt.Sprintf("%s%s%s", s.Prefix, m, s.Suffix)
}
