package golibxml

import "testing"

func TestIsBooleanAttr(t *testing.T) {
	if IsBooleanAttr("width") == true {
		t.Fail()
	}
	if IsBooleanAttr("readonly") == false {
		t.Fail()
	}
}

func TestNewHTMLDoc(t *testing.T) {
	doc := NewHTMLDoc("", "")
	if doc == nil {
		t.Fail()
	}
}
