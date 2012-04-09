package htmlparser

import "testing"

import "github.com/jbussdieker/golibxml/xmltree"

func TestParseDoc(t *testing.T) {
	doc := ParseDoc("<html></html>", "UTF-8")
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	buffer := xmltree.NewBuffer()
	defer buffer.Free()
	doc.NodeDump(buffer, doc.Node)
	println(buffer.Content())
}
