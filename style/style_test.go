package style_test

import (
	"fmt"
	"testing"

	"github.com/gohxs/prettylog/style"
)

func TestNewDefault(t *testing.T) {
	s := style.NewDefault(style.Options{
		"red": {Color: "\033[31m"},
	})

	res := fmt.Sprintf(s.Get("red", "msg"))

	expected := "\033[31mmsg\033[0m"
	if res != expected {
		t.Fatalf("Did not match results %q != %q", res, expected)
	}

}

func TestSetOne(t *testing.T) {
	s := style.New()

	s.SetOne("Hello", style.Option{Color: "\033[01m"})

	res := fmt.Sprintf(s.Get("Hello", "msg"))

	expected := "\033[01mmsg\033[0m"
	if res != expected {
		t.Fatalf("Did not match results %q != %q", res, expected)
	}
}

func TestDisable(t *testing.T) {
	s := style.New()

	s.SetOne("Hello", style.Option{Color: "\033[01m"})

	{
		res := fmt.Sprintf(s.Get("Hello", "msg"))

		expected := "\033[01mmsg\033[0m"
		if res != expected {
			t.Fatalf("Did not match results %q != %q", res, expected)
		}
	}
	s.Enable(false)
	{
		res := fmt.Sprintf(s.Get("Hello", "msg"))
		expected := "msg"
		if res != expected {
			t.Fatalf("Did not match results %q != %q", res, expected)
		}

	}
}
