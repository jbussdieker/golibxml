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

func TestNewXPathSearch(t *testing.T) {
	doc := ParseDoc("<catalog>asdf</catalog>")
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	ctx := NewXPathContext(doc)
	xpath := ctx.Eval("//*")
	if xpath == nil {
		t.Fail()
	}
	//t.Log(xpath.GetNodeSet().Ptr)
	t.Log(xpath.ConvertString().String())
}

