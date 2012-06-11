package golibxml

import "testing"
//import "syscall"

type ElementTypeTestCase struct {
	got ElementType
	expected string
}

var element_type_tests[] ElementTypeTestCase = []ElementTypeTestCase{
	{ XML_ELEMENT_NODE, "Node" },
	{ XML_ATTRIBUTE_NODE, "Attribute" },
	{ XML_TEXT_NODE, "Text" },
}

func getRSS() uint64 {
/*
	rusage := &syscall.Rusage{}
	ret := syscall.Getrusage(0, rusage)
	if ret == nil && rusage.Maxrss > 0 {
		return uint64(rusage.Maxrss)
	}
*/
	return 0
}

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

func TestNewBufferLeak(t *testing.T) {
	var buffer *Buffer
	for i := 0; i < 1000000; i++ {
		buffer = testNewBuffer(t)
		buffer.Free()
	}
	if getRSS() > 4000 {
		t.Fatal("Memory leak")
	}
}

func TestNewBufferSize(t *testing.T) {
	buffer := NewBufferSize(10)
	if buffer.Ptr == nil {
		t.Fail()
	}
	return
}

func TestNewBufferSizeLeak(t *testing.T) {
	var buffer *Buffer
	for i := 0; i < 1000000; i++ {
		buffer = NewBufferSize(1024)
		buffer.Free()
	}
	if getRSS() > 4000 {
		t.Fatal("Memory leak")
	}
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

func TestNewDoc(t *testing.T) {
	doc := NewDoc("1.0")
	defer doc.Free()
	result := doc.String()
	if result != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" {
		t.Fail()
	}
}

func TestNewNode(t *testing.T) {
	doc := NewDoc("1.0")
	defer doc.Free()
	node := NewNode(nil, "div")
	doc.AddChild(node)
	result := node.String()
	if result != "<div/>" {
		t.Fail()
	}
}

func TestNewNodePath(t *testing.T) {
	doc := NewDoc("1.0")
	defer doc.Free()
	node := NewNode(nil, "div")
	doc.AddChild(node)
	result := node.Path()
	if result != "/div" {
		t.Fatal("Expected: /div Got:", node.Path())
	}
}

func TestNewComment(t *testing.T) {
	doc := NewDoc("1.0")
	defer doc.Free()
	comment := doc.NewComment("this is a comment")
	doc.AddChild(comment)
	result := comment.String()
	if result != "<!--this is a comment-->" {
		t.Fail()
	}
}

func TestElementTypeString(t *testing.T) {
	for _, test := range element_type_tests {
		if test.got.String() != test.expected {
			t.Fatal("Testing node type:", test.got, "got:", test.got.String(), "expected:", test.expected)
		}
	}
}
