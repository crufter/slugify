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

	if S("Hello           World????????") != "hello-world" {
		t.Fail()
	}

	if S("aáäâeéëeiíiîoóöőôuúüűunç·/_,:;") != "aaaaeeeeiiiiooooouuuuunc" {
		t.Fail()
	}

	if S("adam's song") != "adams-song" {
		t.Fail()
	}

	if S("what's my name again?") != "whats-my-name-again" {
		t.Fail()
	}
}
