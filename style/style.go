package style

import "fmt"

//Styler string wrapper
type Styler interface {
	Len(name string) int
	SetOne(name string, style Option)
	Set(Options)
	Get(name string, msg interface{}) string
	GetX(name string, msg interface{}) string
	Enable(enabled bool)
}

// Options map
type Options map[string]Option

// Option for a style entry
type Option struct {
	Color        string
	UseGlobalPad bool
}

type styleEntry struct {
	Opt       Option
	globalPad int
}

// several default styles
type style struct {
	Disabled bool
	stylmap  map[string]*styleEntry
}

// New styler
func New() Styler {
	return &style{stylmap: map[string]*styleEntry{}}
}

// NewDefault New styler with default values as parameter
func NewDefault(opts Options) Styler {
	s := New()
	s.Set(opts)
	return s
}

//func (s *style) color(code string, str interface{}) string {
//	return fmt.Sprintf("%s%v\033[0m", code, str)
//}

//Set a named style
func (s *style) SetOne(name string, styl Option) {
	s.stylmap[name] = &styleEntry{Opt: styl}
}
func (s *style) Set(opts Options) {
	for k, v := range opts {
		s.SetOne(k, v)
	}
}

//Get a msg string with style if enabled
func (s *style) Get(name string, msg interface{}) string {
	if s.Disabled {
		return fmt.Sprint(msg)
	}

	st, ok := s.stylmap[name]
	if !ok {
		return fmt.Sprint(msg)
	}

	m := fmt.Sprintf("%v", msg)

	if st.Opt.UseGlobalPad {
		if t := len(m); t > st.globalPad {
			st.globalPad = t
		}
		m = fmt.Sprintf("%*s", st.globalPad, m)
	}
	// Calc globalpadding if true
	return fmt.Sprintf("%s%s%s", st.Opt.Color, m, "\033[0m")
}

//GetX gets a msg with style or none if disabled
func (s *style) GetX(str string, msg interface{}) string {
	if s.Disabled {
		return ""
	}
	return s.Get(str, msg)
}

//Len with style (escape chars increases string length)
func (s *style) Len(styl string) int {
	t := s.Get(styl, "")
	return len(t)
}

func (s *style) Enable(val bool) {
	s.Disabled = !val
}
