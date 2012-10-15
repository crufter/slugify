package slugify

import (
	"testing"
)

// Test Slugify
func TestSlugify(t *testing.T) {
	if S("Hello World") != "hello-world" {
		t.Fail()
	}
	if S("Hello World69") != "hello-world69" {
		t.Fail()
	}
	if S("Hello           World????????") != "hello-world-" {
		t.Fail()
	}
	if S("aáäâeéëeiíiîoóöőôuúüűunç·/_,:;") != "aaaaeeeeiiiiooooouuuuunc-" {
		t.Fail()
	}
}
