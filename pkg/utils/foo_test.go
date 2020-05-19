package utils

import "testing"

func TestFooWord(t *testing.T) {
	if "foo_word" != FooWord("word") {
		t.FailNow()
	}
}
