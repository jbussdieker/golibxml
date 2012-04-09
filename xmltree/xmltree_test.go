package xmltree

import "testing"

//
// Buffer tests
func testNewBuffer(t *testing.T) (buffer Buffer) {
	buffer = BufferCreate()
	if buffer.Ptr == nil {
		t.Fail()
	}
	return
}

func TestBufferCreate(t *testing.T) {
	testNewBuffer(t)
}

func TestBufferWriteChar(t *testing.T) {
	buffer := testNewBuffer(t)
	buffer.BufferWriteChar("test")
}

func TestBufferEmpty(t *testing.T) {
	buffer := testNewBuffer(t)
	buffer.BufferWriteChar("test")
	buffer.BufferEmpty()
	if buffer.BufferContent() != "" {
		t.Fail()
	}
}

func TestBufferCat(t *testing.T) {
	buffer := testNewBuffer(t)
	buffer.BufferCat("test")
}

func TestBufferContent(t *testing.T) {
	buffer := testNewBuffer(t)
	buffer.BufferWriteChar("test")
	if buffer.BufferContent() != "test" {
		t.Fail()
	}
}

//
// Advanced/Combo tests
func TestNewDocAdv(t *testing.T) {
	doc := NewDoc("1.0")
	buffer := BufferCreate()
	result := doc.NodeDump(buffer, doc.NodePtr, 0, 0)

	println("Result Size:", result)
	println("Result:", buffer.BufferContent())
}

func TestNewNode(t *testing.T) {
	doc := NewDoc("1.0")
	node := NewNode(nil, "div")
	doc.AddChild(node)
	buffer := BufferCreate()
	result := doc.NodeDump(buffer, doc.NodePtr, 0, 0)

	println("Result Size:", result)
	println("Result:", buffer.BufferContent())
}

func TestNewComment(t *testing.T) {
	doc := NewDoc("1.0")
	comment := doc.NewDocComment("this is a comment")
	doc.AddChild(comment)
	buffer := BufferCreate()
	result := doc.NodeDump(buffer, doc.NodePtr, 0, 0)

	println("Result Size:", result)
	println("Result:", buffer.BufferContent())
}

