package golibxml

func (doc *HTMLDocument) String() string {
	buf := NewBuffer()
	defer buf.Free()
	doc.NodeDump(buf, doc.HTMLNode)
	return buf.Content()
}
