package golibxml

import "testing"

func TestParseDoc(t *testing.T) {
	doc := ParseDoc("<html></html>")
	if doc == nil {
		t.Fail()
	}
}
