package style_test

import (
	"fmt"
	"testing"

	"github.com/gohxs/clistyle"
)

type MyStyle struct {
	Keyword clistyle.Style
	Pad     clistyle.Style
	Test    clistyle.Style
}

var (
	Style = MyStyle{
		Keyword: clistyle.Style{Prefix: "pre"},
		Pad:     clistyle.Style{IncrementPad: true},
		Test:    clistyle.Style{IncrementPad: true},
	}
)

func TestUse(t *testing.T) {
	res := fmt.Sprintf(Style.Keyword.Use("test"))
	if res != "pretest" {
		t.Fatal("output does not match")
	}

}

func TestPad(t *testing.T) {

	Style.Pad.Use("msg1")
	Style.Pad.Use("msg1oooo")
	Style.Pad.Use("msg1te")
	res := Style.Pad.Use("msg1")
	if res != "    msg1" {
		t.Fatal("Style padding fail")
	}
}

func TestPad2(t *testing.T) {

	Style.Pad.Use("msg1")
	Style.Pad.Use("msg1oooo")
	Style.Pad.Use("msg1te")
	res := Style.Pad.Use("msg1")
	if res != "    msg1" {
		t.Fatal("Style padding fail")
	}

	res = Style.Test.Use("m")

	if res != "m" {
		t.Fatal("result failed")
	}

}
