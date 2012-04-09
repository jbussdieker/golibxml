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

////////////////////////////////////////////////////////////////////////////////
// TYPES/STRUCTS
////////////////////////////////////////////////////////////////////////////////

type AllocationScheme int

type AttrPtr struct {
	Ptr C.xmlAttrPtr
}

type NodePtr struct {
	Ptr C.xmlNodePtr
}

type DocPtr struct {
	*NodePtr
	Ptr C.xmlDocPtr
}

type NsPtr struct {
	Ptr C.xmlNsPtr
}

type Buffer struct {
	Ptr C.xmlBufferPtr
}

////////////////////////////////////////////////////////////////////////////////
// CONSTANTS/ENUM
////////////////////////////////////////////////////////////////////////////////

const (
    XML_BUFFER_ALLOC_DOUBLEIT AllocationScheme = 1 //: double each time one need to grow
    XML_BUFFER_ALLOC_EXACT = 2 //: grow only to the minimal size
    XML_BUFFER_ALLOC_IMMUTABLE = 3 //: immutable buffer
    XML_BUFFER_ALLOC_IO = 4 //: special allocation scheme used for I/O
)

////////////////////////////////////////////////////////////////////////////////
// INTERFACE
////////////////////////////////////////////////////////////////////////////////

// xmlAddChild
func (parent *NodePtr) AddChild(cur *NodePtr) (*NodePtr) {
	return &NodePtr{C.xmlAddChild(parent.Ptr, cur.Ptr)}
}

// xmlAddChildList
func (parent *NodePtr) AddChildList(cur NodePtr) (*NodePtr) {
	return &NodePtr{C.xmlAddNextSibling(parent.Ptr, cur.Ptr)}
}

// xmlAddNextSibling
func (cur *NodePtr) AddNextSibling(elem NodePtr) (*NodePtr) {
	return &NodePtr{C.xmlAddNextSibling(cur.Ptr, elem.Ptr)}
}

// xmlAddPrevSibling
func (cur *NodePtr) AddPrevSibling(elem NodePtr) (*NodePtr) {
	return &NodePtr{C.xmlAddPrevSibling(cur.Ptr, elem.Ptr)}
}

// xmlAddSibling
func (cur *NodePtr) AddSibling(elem NodePtr) (*NodePtr) {
	return &NodePtr{C.xmlAddSibling(cur.Ptr, elem.Ptr)}
}

// xmlBufferCat/xmlBufferCCat
func (buffer *Buffer) Cat(str string) int {
	ptr := C.CString(str)
	defer C.free_string(ptr)
	return int(C.xmlBufferCCat(buffer.Ptr, ptr))
}

// xmlBufferContent
func (buffer *Buffer) Content() string {
	return C.GoString(C.to_charptr(C.xmlBufferContent(buffer.Ptr)))
}

// xmlBufferCreate
func NewBuffer() (*Buffer) {
	return &Buffer{C.xmlBufferCreate()}
}

// xmlBufferCreateSize
func NewBufferSize(size int) (*Buffer) {
	return &Buffer{C.xmlBufferCreateSize(C.size_t(size))}
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

// xmlBufferLength
func (buffer *Buffer) Length() int {
	return int(C.xmlBufferLength(buffer.Ptr))
}

// xmlBufferResize
func (buffer *Buffer) Resize(size int) int {
	return int(C.xmlBufferResize(buffer.Ptr, C.uint(size)))
}

// xmlBufferSetAllocationScheme
func (buffer *Buffer) SetAllocationScheme(scheme AllocationScheme) {
	C.xmlBufferSetAllocationScheme(buffer.Ptr, C.xmlBufferAllocationScheme(scheme))
}

// xmlBufferShrink
func (buffer *Buffer) Shrink(length int) int {
	return int(C.xmlBufferShrink(buffer.Ptr, C.uint(length)))
}

// xmlBufferWriteChar/xmlBufferWriteCHAR
func (buffer *Buffer) WriteChar(str string) {
	ptr := C.CString(str)
	defer C.free_string(ptr)
	C.xmlBufferWriteChar(buffer.Ptr, ptr)
}

// xmlChildElementCount
func (node *NodePtr) ChildElementCount() int {
	return int(C.xmlChildElementCount(node.Ptr))
}

// xmlCopyDoc
func (doc *DocPtr) Copy(recursive int) (*DocPtr) {
	docptr := C.xmlCopyDoc(doc.Ptr, C.int(recursive))
	return &DocPtr{
		Ptr: docptr, 
		NodePtr: &NodePtr{C.xmlNodePtr(unsafe.Pointer(docptr))},
	}
}

// xmlCopyNamespace
func (ns *NsPtr) Copy(extended int) (*NsPtr) {
	return &NsPtr{C.xmlCopyNamespace(ns.Ptr)}
}

// xmlCopyNode
func (node *NodePtr) Copy(extended int) (*NodePtr) {
	return &NodePtr{C.xmlCopyNode(node.Ptr, C.int(extended))}
}

// xmlDocGetRootElement
func (doc *DocPtr) GetRootElement() (*NodePtr) {
	return &NodePtr{C.xmlDocGetRootElement(doc.Ptr)}
}

// xmlDocSetRootElement
func (doc *DocPtr) SetRootElement(root *NodePtr) (*NodePtr) {
	return &NodePtr{C.xmlDocSetRootElement(doc.Ptr, root.Ptr)}
}

// xmlFirstElementChild
func (node *NodePtr) FirstElementChild() (*NodePtr) {
	return &NodePtr{C.xmlFirstElementChild(node.Ptr)}
}

// xmlFreeDoc
func (doc *DocPtr) Free() {
	C.xmlFreeDoc(doc.Ptr)
	doc.Ptr = nil
	doc.NodePtr = nil
}

// xmlFreeNode
func (node *NodePtr) Free() {
	C.xmlFreeNode(node.Ptr)
	node.Ptr = nil
}

// xmlFreeNs
func (ns *NsPtr) Free() {
	C.xmlFreeNs(ns.Ptr)
	ns.Ptr = nil
}

// xmlNewComment
func NewComment(content string) (*NodePtr) {
	ptr := C.CString(content)
	defer C.free_string(ptr)
	return &NodePtr{C.xmlNewComment(C.to_xmlcharptr(ptr))}
}

// xmlNewDoc
func NewDoc(version string) (*DocPtr) {
	ptr := C.CString(version)
	defer C.free_string(ptr)
	doc := C.xmlNewDoc(C.to_xmlcharptr(ptr))
	return &DocPtr{
		Ptr: doc, 
		NodePtr: &NodePtr{C.xmlNodePtr(unsafe.Pointer(doc))},
	}
}

// xmlNewDocComment
func (doc *DocPtr) NewComment(content string) (*NodePtr) {
	ptr := C.CString(content)
	defer C.free_string(ptr)
	return &NodePtr{C.xmlNewDocComment(doc.Ptr, C.to_xmlcharptr(ptr))}
}

// xmlNewDocFragment
func (doc *DocPtr) NewFragment() (*NodePtr) {
	return &NodePtr{C.xmlNewDocFragment(doc.Ptr)}
}

// xmlNewDocNode
func (doc *DocPtr) NewNode(ns *NsPtr, name string, content string) (*NodePtr) {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	ptrc := C.CString(content)
	defer C.free_string(ptrc)
	if ns != nil {
		return &NodePtr{C.xmlNewDocNode(doc.Ptr, ns.Ptr, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrc))}
	}
	return &NodePtr{C.xmlNewDocNode(doc.Ptr, nil, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrc))}	
}

// xmlNewDocProp
func (doc *DocPtr) NewProp(name string, value string) (*AttrPtr) {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	ptrv := C.CString(value)
	defer C.free_string(ptrv)
	return &AttrPtr{C.xmlNewDocProp(doc.Ptr, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrv))}
}

// xmlNewDocRawNode
func (doc *DocPtr) NewRawNode(ns *NsPtr, name string, content string) (*NodePtr) {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	ptrc := C.CString(content)
	defer C.free_string(ptrc)
	if ns != nil {
		return &NodePtr{C.xmlNewDocRawNode(doc.Ptr, ns.Ptr, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrc))}
	}
	return &NodePtr{C.xmlNewDocRawNode(doc.Ptr, nil, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrc))}	
}

// xmlNewNode
func NewNode(ns *NsPtr, name string) (*NodePtr) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	if ns != nil {
		return &NodePtr{C.xmlNewNode(ns.Ptr, C.to_xmlcharptr(ptr))}
	}
	return &NodePtr{C.xmlNewNode(nil, C.to_xmlcharptr(ptr))}
}

// xmlNewNs
func (node *NodePtr) NewNs(href string, prefix string) *NsPtr {
	ptrh := C.CString(href)
	defer C.free_string(ptrh)
	ptrp := C.CString(prefix)
	defer C.free_string(ptrp)
	return &NsPtr{C.xmlNewNs(node.Ptr, C.to_xmlcharptr(ptrh), C.to_xmlcharptr(ptrp))}
}

// xmlNodeDump
func (doc *DocPtr) NodeDump(buf *Buffer, cur *NodePtr, level int, format int) int {
	return int(C.xmlNodeDump(buf.Ptr, doc.Ptr, cur.Ptr, C.int(level), C.int(format)))
}

// xmlNodeGetContent
func (node *NodePtr) GetContent() string {
	return C.GoString(C.to_charptr(C.xmlNodeGetContent(node.Ptr)))
}

