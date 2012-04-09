package xmltree
/*
#cgo pkg-config: libxml-2.0
#include <libxml/tree.h>

static inline void free_string(char* s) { free(s); }
static inline xmlChar *to_xmlcharptr(const char *s) { return (xmlChar *)s; }
static inline char *to_charptr(const xmlChar *s) { return (char *)s; }

*/
import "C"
import "unsafe"

type NodePtr struct {
	Ptr C.xmlNodePtr
}

type DocPtr struct {
	NodePtr
	Ptr C.xmlDocPtr
}

type NsPtr struct {
	Ptr C.xmlNsPtr
}

type Buffer struct {
	Ptr C.xmlBufferPtr
}

// xmlAddChild
func (parent NodePtr) AddChild(cur NodePtr) (NodePtr) {
	return NodePtr{C.xmlAddChild(parent.Ptr, cur.Ptr)}
}

// xmlAddChildList
func (parent NodePtr) AddChildList(cur NodePtr) (NodePtr) {
	return NodePtr{C.xmlAddNextSibling(parent.Ptr, cur.Ptr)}
}

// xmlAddNextSibling
func (cur NodePtr) AddNextSibling(elem NodePtr) (NodePtr) {
	return NodePtr{C.xmlAddNextSibling(cur.Ptr, elem.Ptr)}
}

// xmlAddPrevSibling
func (cur NodePtr) AddPrevSibling(elem NodePtr) (NodePtr) {
	return NodePtr{C.xmlAddPrevSibling(cur.Ptr, elem.Ptr)}
}

// xmlAddSibling
func (cur NodePtr) AddSibling(elem NodePtr) (NodePtr) {
	return NodePtr{C.xmlAddSibling(cur.Ptr, elem.Ptr)}
}

// xmlBufferCat/xmlBufferCCat
func (buffer Buffer) Cat(str string) int {
	ptr := C.CString(str)
	defer C.free_string(ptr)
	return int(C.xmlBufferCCat(buffer.Ptr, ptr))
}

// xmlBufferContent
func (buffer Buffer) Content() string {
	return C.GoString(C.to_charptr(C.xmlBufferContent(buffer.Ptr)))
}

// xmlBufferCreate
func BufferCreate() Buffer {
	return Buffer{C.xmlBufferCreate()}
}

// xmlBufferEmpty
func (buffer *Buffer) Empty() {
	C.xmlBufferEmpty(buffer.Ptr)
}

// xmlBufferFree
func (buffer *Buffer) Free() {
	C.xmlBufferFree(buffer.Ptr)
	buffer.Ptr = nil
}

// xmlBufferGrow
func (buffer *Buffer) Grow(length int) int {
	return int(C.xmlBufferGrow(buffer.Ptr, C.uint(length)))
}

// xmlBufferWriteChar/xmlBufferWriteCHAR
func (buffer *Buffer) WriteChar(str string) {
	ptr := C.CString(str)
	defer C.free_string(ptr)
	C.xmlBufferWriteChar(buffer.Ptr, ptr)
}

// xmlNewComment
func NewComment(content string) (NodePtr) {
	ptr := C.CString(content)
	defer C.free_string(ptr)
	return NodePtr{C.xmlNewComment(C.to_xmlcharptr(ptr))}
}

// xmlNewDoc
func NewDoc(version string) (DocPtr) {
	ptr := C.CString(version)
	defer C.free_string(ptr)
	doc := C.xmlNewDoc(C.to_xmlcharptr(ptr))
	return DocPtr{
		Ptr: doc, 
		NodePtr: NodePtr{C.xmlNodePtr(unsafe.Pointer(doc))},
	}
}

// xmlNewDocComment
func (doc *DocPtr) NewDocComment(content string) (NodePtr) {
	ptr := C.CString(content)
	defer C.free_string(ptr)
	return NodePtr{C.xmlNewDocComment(doc.Ptr, C.to_xmlcharptr(ptr))}
}

// xmlNewDocFragment
func (doc *DocPtr) NewDocFragment() (NodePtr) {
	return NodePtr{C.xmlNewDocFragment(doc.Ptr)}
}

// xmlNewNode
func NewNode(ns *NsPtr, name string) (NodePtr) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	if ns != nil {
		return NodePtr{C.xmlNewNode(ns.Ptr, C.to_xmlcharptr(ptr))}
	}
	return NodePtr{C.xmlNewNode(nil, C.to_xmlcharptr(ptr))}
}

// xmlNewNs
func (node *NodePtr) NewNs(href string, prefix string) NsPtr {
	ptrh := C.CString(href)
	defer C.free_string(ptrh)
	ptrp := C.CString(prefix)
	defer C.free_string(ptrp)
	return NsPtr{C.xmlNewNs(node.Ptr, C.to_xmlcharptr(ptrh), C.to_xmlcharptr(ptrp))}
}

// xmlNodeDump
func (doc *DocPtr) NodeDump(buf Buffer, cur NodePtr, level int, format int) int {
	return int(C.xmlNodeDump(buf.Ptr, doc.Ptr, cur.Ptr, C.int(level), C.int(format)))
}

// xmlNodeGetContent
func (node *NodePtr) NodeGetContent() string {
	return C.GoString(C.to_charptr(C.xmlNodeGetContent(node.Ptr)))
}

