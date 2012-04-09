package xmlparser

import "testing"

import . "github.com/jbussdieker/golibxml/xmltree"

func TestParseDoc(t *testing.T) {
	doc := ParseDoc("<html></html>")
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	buffer := NewBuffer()
	defer buffer.Free()
	doc.NodeDump(buffer, doc.Node, 0, 0)
	println(buffer.Content())
}
