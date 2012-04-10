package golibxml

import "testing"

func TestParseDoc(t *testing.T) {
	doc := ParseDoc("<catalog/>")
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	buffer := NewBuffer()
	defer buffer.Free()
	doc.NodeDump(buffer, doc.Node, 0, 0)
	if buffer.Content() != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<catalog/>\n" {
		println(buffer.Content())
		t.Fail()
	}
}

func TestReadDoc(t *testing.T) {
	doc := ReadDoc("<catalog/>", "", "", 0)
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	buffer := NewBuffer()
	defer buffer.Free()
	doc.NodeDump(buffer, doc.Node, 0, 0)
	if buffer.Content() != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<catalog/>\n" {
		println(buffer.Content())
		t.Fail()
	}
}

func TestGetFeaturesList(t *testing.T) {
	features := GetFeaturesList()
	if len(features) != 42 {
		t.Fail()
	}
}
