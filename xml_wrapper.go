package golibxml

import "unsafe"

func (doc *Document) String() string {
	buf := NewBuffer()
	defer buf.Free()
	doc.NodeDump(buf, doc.Node, 0, 0)
	return buf.Content()
}

func (node *Node) Document() *Document {
	return makeDoc(_Ctype_xmlDocPtr(unsafe.Pointer(node.Ptr.doc)))
}

func (node *Node) String() string {
	buf := NewBuffer()
	defer buf.Free()
	node.Document().NodeDump(buf, node, 0, 0)
	return buf.Content()
}
