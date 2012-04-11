package golibxml

import "testing"

func TestXPathCompile(t *testing.T) {
	xpath := CompileXPath("*")
	if xpath == nil {
		t.Fail()
	}
}

func TestXPathCompileNeg(t *testing.T) {
	xpath := CompileXPath("")
	if xpath != nil {
		t.Fail()
	}
}

func TestNewXPathContext(t *testing.T) {
	doc := ParseDoc("<catalog/>")
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	ctx := NewXPathContext(doc)
	xpath := ctx.Compile("*")
	if xpath == nil {
		t.Fail()
	}

}

