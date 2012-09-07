package slugify_test

import (
	"../slugify"
	"testing"
)

// Test Slugify
func TestSlugify(t *testing.T) {
	if slugify.S("Hello World") != "hello-world" {
		t.Fail()
	}
	if slugify.S("Hello World69") != "hello-world69" {
		t.Fail()
	}
	if slugify.S("Hello           World????????") != "hello-world-" {
		t.Fail()
	}
	if slugify.S("aáäâeéëeiíiîoóöőôuúüűunç·/_,:;") != "aaaaeeeeiiiiooooouuuuunc-" {
		t.Fail()
	}
}
