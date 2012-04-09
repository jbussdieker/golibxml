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
	buffer.WriteChar("test")
}

func TestBufferEmpty(t *testing.T) {
	buffer := testNewBuffer(t)
	buffer.WriteChar("test")
	buffer.Empty()
	if buffer.Content() != "" {
		t.Fail()
	}
}

func TestBufferCat(t *testing.T) {
	buffer := testNewBuffer(t)
	buffer.Cat("test")
}

func TestBufferContent(t *testing.T) {
	buffer := testNewBuffer(t)
	buffer.WriteChar("test")
	if buffer.Content() != "test" {
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
	println("Result:", buffer.Content())
}

func TestNewNode(t *testing.T) {
	doc := NewDoc("1.0")
	node := NewNode(nil, "div")
	doc.AddChild(node)
	buffer := BufferCreate()
	result := doc.NodeDump(buffer, doc.NodePtr, 0, 0)

	println("Result Size:", result)
	println("Result:", buffer.Content())
}

func TestNewComment(t *testing.T) {
	doc := NewDoc("1.0")
	comment := doc.NewDocComment("this is a comment")
	doc.AddChild(comment)
	buffer := BufferCreate()
	result := doc.NodeDump(buffer, doc.NodePtr, 0, 0)

	println("Result Size:", result)
	println("Result:", buffer.Content())
}

