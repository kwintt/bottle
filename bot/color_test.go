package bot

import (
	"testing"
)

func TestColor(t *testing.T) {
	r, g, b := Color("#FFFFFF")
	if r != 255 {
		t.Error(r)
	}
	if g != 255 {
		t.Error(g)
	}
	if b != 255 {
		t.Error(b)
	}
}

func TestColor_green(t *testing.T) {
	r, g, b := Color("#5B9B7B")
	if r != 91 {
		t.Error(r)
	}
	if g != 155 {
		t.Error(g)
	}
	if b != 123 {
		t.Error(b)
	}
}