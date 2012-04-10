package golibxml

import "testing"

func TestParseHTMLDoc(t *testing.T) {
	doc := ParseHTMLDoc("<html></html>", "UTF-8")
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	buffer := NewBuffer()
	defer buffer.Free()
	doc.NodeDump(buffer, doc.HTMLNode)
	//println(buffer.Content())
}

/*
func TestDocAutoCloseTag(t *testing.T) {
	doc := ParseHTMLDoc("<html></html>", "UTF-8")
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()

	n := doc.NewNode(nil, "head", "test")
	println("BR:", doc.AutoCloseTag("br", n))
	println("A:", doc.AutoCloseTag("link", n))
	doc.Root().AddChild(n)

	buffer := NewBuffer()
	defer buffer.Free()
	doc.NodeDump(buffer, doc.HTMLNode)

	println(buffer.Content())
}*/
