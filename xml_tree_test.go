package golibxml

import "testing"

//
// Buffer tests
func testNewBuffer(t *testing.T) (buffer *Buffer) {
	buffer = NewBuffer()
	if buffer.Ptr == nil {
		t.Fail()
	}
	return
}

func testBufferFree(t *testing.T, buffer *Buffer) {
	buffer.Free()
	if buffer.Ptr != nil {
		t.Fail()
	}
}

func TestNewBuffer(t *testing.T) {
	testNewBuffer(t)
}

func TestNewBufferSize(t *testing.T) {
	buffer := NewBufferSize(10)
	if buffer.Ptr == nil {
		t.Fail()
	}
	return
}

func TestBufferFree(t *testing.T) {
	buffer := testNewBuffer(t)
	defer testBufferFree(t, buffer)
}

func TestBufferWriteChar(t *testing.T) {
	buffer := testNewBuffer(t)
	defer testBufferFree(t, buffer)
	buffer.WriteChar("test")
}

func TestBufferEmpty(t *testing.T) {
	buffer := testNewBuffer(t)
	defer testBufferFree(t, buffer)
	buffer.WriteChar("test")
	buffer.Empty()
	if buffer.Content() != "" {
		t.Fail()
	}
}

func TestBufferResize(t *testing.T) {
	buffer := testNewBuffer(t)
	defer testBufferFree(t, buffer)
	if buffer.Resize(10) != 1 {
		t.Fail()
	}
}

func TestBufferLength(t *testing.T) {
	buffer := testNewBuffer(t)
	defer testBufferFree(t, buffer)
	if buffer.Length() != 0 {
		t.Fail()
	}
	buffer.WriteChar("test")
	if buffer.Length() != 4 {
		t.Fail()
	}
}
/*
func TestBufferGrow(t *testing.T) {
	buffer := testNewBuffer(t)
	defer testBufferFree(t, buffer)
	buffer.Grow(128)
	if buffer.Length() != 128 {
		t.Fail()
	}
}

func TestBufferShrink(t *testing.T) {
	buffer := testNewBuffer(t)
	defer testBufferFree(t, buffer)
	buffer.Shrink(128)
	if buffer.Length() != 128 {
		t.Fail()
	}
}
*/
func TestBufferCat(t *testing.T) {
	buffer := testNewBuffer(t)
	defer testBufferFree(t, buffer)
	buffer.Cat("test")
}

func TestBufferContent(t *testing.T) {
	buffer := testNewBuffer(t)
	defer testBufferFree(t, buffer)
	buffer.WriteChar("test")
	if buffer.Content() != "test" {
		t.Fail()
	}
}

//
// Advanced/Combo tests
func TestNewDocAdv(t *testing.T) {
	doc := NewDoc("1.0")
	defer doc.Free()
	buffer := NewBuffer()
	result := doc.NodeDump(buffer, doc.Node, 0, 0)

	if result != 39 {
		println("Result Size:", result)
		println("Result:", buffer.Content())
		t.Fail()
	}
}

func TestNewNode(t *testing.T) {
	doc := NewDoc("1.0")
	defer doc.Free()
	node := NewNode(nil, "div")
	doc.AddChild(node)
	buffer := NewBuffer()
	result := doc.NodeDump(buffer, doc.Node, 0, 0)

	if result != 46 {
		println("Result Size:", result)
		println("Result:", buffer.Content())
		t.Fail()
	}
}

func TestNewComment(t *testing.T) {
	doc := NewDoc("1.0")
	defer doc.Free()
	comment := doc.NewComment("this is a comment")
	doc.AddChild(comment)
	buffer := NewBuffer()
	result := doc.NodeDump(buffer, doc.Node, 0, 0)

	if result != 64 {
		println("Result Size:", result)
		println("Result:", buffer.Content())
		t.Fail()
	}
}
