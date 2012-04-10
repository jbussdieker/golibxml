package golibxml

func (doc *Document) String() string {
	buf := NewBuffer()
	defer buf.Free()
	doc.NodeDump(buf, doc.Node, 0, 0)
	return buf.Content()
}
