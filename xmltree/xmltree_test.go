package xmltree

import "testing"

func TestSomething(t *testing.T) {
	doc := NewDoc("1.0")
	buffer := BufferCreate()
	result := doc.NodeDump(buffer, doc.NodePtr, 0, 0)
	println("Result Size:", result)
	println("Result:", buffer.BufferContent())
}

