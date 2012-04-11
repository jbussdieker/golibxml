package golibxml

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

type Dtd struct {
	Ptr C.xmlDtdPtr
}

type Attribute struct {
	Ptr C.xmlAttrPtr
}

type Node struct {
	Ptr C.xmlNodePtr
}

type TextNode struct {
	*Node
}

type Document struct {
	*Node
	Ptr C.xmlDocPtr
}

type Namespace struct {
	Ptr C.xmlNsPtr
}

type Buffer struct {
	Ptr C.xmlBufferPtr
}

////////////////////////////////////////////////////////////////////////////////
// CONSTANTS/ENUM
////////////////////////////////////////////////////////////////////////////////

const (
	XML_BUFFER_ALLOC_DOUBLEIT  AllocationScheme = 1 //: double each time one need to grow
	XML_BUFFER_ALLOC_EXACT                      = 2 //: grow only to the minimal size
	XML_BUFFER_ALLOC_IMMUTABLE                  = 3 //: immutable buffer
	XML_BUFFER_ALLOC_IO                         = 4 //: special allocation scheme used for I/O
)

////////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func makeDoc(doc C.xmlDocPtr) *Document {
	if doc == nil {
		return nil
	}
	return &Document{
		Ptr:  doc,
		Node: &Node{C.xmlNodePtr(unsafe.Pointer(doc))},
	}
}

func makeDtd(dtd C.xmlDtdPtr) *Dtd {
	if dtd == nil {
		return nil
	}
	return &Dtd{dtd}
}

func makeNode(node C.xmlNodePtr) *Node {
	if node == nil {
		return nil
	}
	return &Node{node}
}

func makeAttribute(attr C.xmlAttrPtr) *Attribute {
	if attr == nil {
		return nil
	}
	return &Attribute{attr}
}

////////////////////////////////////////////////////////////////////////////////
// INTERFACE
////////////////////////////////////////////////////////////////////////////////

// xmlAddChild
func (parent *Node) AddChild(cur *Node) *Node {
	return makeNode(C.xmlAddChild(parent.Ptr, cur.Ptr))
}

// xmlAddChildList
func (parent *Node) AddChildList(cur Node) *Node {
	return makeNode(C.xmlAddNextSibling(parent.Ptr, cur.Ptr))
}

// xmlAddNextSibling
func (cur *Node) AddNextSibling(elem Node) *Node {
	return makeNode(C.xmlAddNextSibling(cur.Ptr, elem.Ptr))
}

// xmlAddPrevSibling
func (cur *Node) AddPrevSibling(elem Node) *Node {
	return makeNode(C.xmlAddPrevSibling(cur.Ptr, elem.Ptr))
}

// xmlAddSibling
func (cur *Node) AddSibling(elem Node) *Node {
	return makeNode(C.xmlAddSibling(cur.Ptr, elem.Ptr))
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
func NewBuffer() *Buffer {
	return &Buffer{C.xmlBufferCreate()}
}

// xmlBufferCreateSize
func NewBufferSize(size int) *Buffer {
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
func (node *Node) ChildCount() int {
	return int(C.xmlChildElementCount(node.Ptr))
}

// xmlCopyDoc
func (doc *Document) Copy(recursive int) *Document {
	cdoc := C.xmlCopyDoc(doc.Ptr, C.int(recursive))
	return makeDoc(cdoc)
}

// xmlCopyDtd
func (dtd *Dtd) Copy() *Dtd {
	return makeDtd(C.xmlCopyDtd(dtd.Ptr))
}

// xmlCopyNamespace
func (ns *Namespace) Copy(extended int) *Namespace {
	return &Namespace{C.xmlCopyNamespace(ns.Ptr)}
}

// xmlCopyNamespaceList
func (ns *Namespace) CopyList(extended int) *Namespace {
	return &Namespace{C.xmlCopyNamespaceList(ns.Ptr)}
}

// xmlCopyNode
func (node *Node) Copy(extended int) *Node {
	cnode := C.xmlCopyNode(node.Ptr, C.int(extended))
	return makeNode(cnode)
}

// xmlCopyNodeList
func (node *Node) CopyList() *Node {
	cnode := C.xmlCopyNodeList(node.Ptr)
	return makeNode(cnode)
}

// xmlCopyProp
func (attr *Attribute) Copy(target *Node) *Attribute {
	cattr := C.xmlCopyProp(target.Ptr, attr.Ptr)
	return makeAttribute(cattr)
}

// xmlCopyPropList
func (attr *Attribute) CopyList(target *Node) *Attribute {
	cattr := C.xmlCopyPropList(target.Ptr, attr.Ptr)
	return makeAttribute(cattr)
}

// xmlDocGetRootElement
func (doc *Document) Root() *Node {
	cnode := C.xmlDocGetRootElement(doc.Ptr)
	return makeNode(cnode)
}

// xmlDocSetRootElement
func (doc *Document) SetRoot(root *Node) *Node {
	cnode := C.xmlDocSetRootElement(doc.Ptr, root.Ptr)
	return makeNode(cnode)
}

// xmlFirstElementChild
func (node *Node) FirstChild() *Node {
	cnode := C.xmlFirstElementChild(node.Ptr)
	return makeNode(cnode)
}

// xmlFreeDoc
func (doc *Document) Free() {
	C.xmlFreeDoc(doc.Ptr)
	doc.Ptr = nil
	doc.Node = nil
}

// xmlFreeDtd
func (dtd *Dtd) Free() {
	C.xmlFreeDtd(dtd.Ptr)
	dtd.Ptr = nil
}

// xmlFreeNode
func (node *Node) Free() {
	C.xmlFreeNode(node.Ptr)
	node.Ptr = nil
}

// xmlFreeNodeList
func (node *Node) FreeList() {
	C.xmlFreeNodeList(node.Ptr)
	node.Ptr = nil
}

// xmlFreeNs
func (ns *Namespace) Free() {
	C.xmlFreeNs(ns.Ptr)
	ns.Ptr = nil
}

// xmlFreeNsList
func (ns *Namespace) FreeList() {
	C.xmlFreeNsList(ns.Ptr)
	ns.Ptr = nil
}

// xmlFreeProp
func (attr *Attribute) Free() {
	C.xmlFreeProp(attr.Ptr)
	attr.Ptr = nil
}

// xmlFreePropList
func (attr *Attribute) FreeList() {
	C.xmlFreePropList(attr.Ptr)
	attr.Ptr = nil
}

// xmlGetLastChild
func (node *Node) LastChild() *Node {
	return makeNode(C.xmlGetLastChild(node.Ptr))
}

// xmlGetNodePath
func (node *Node) Path() string {
	cstr := C.xmlGetNodePath(node.Ptr)
	return C.GoString(C.to_charptr(cstr))
}

// xmlLastElementChild
func (node *Node) LastElementChild() *Node {
	return makeNode(C.xmlLastElementChild(node.Ptr))
}

// xmlNewChild
func (node *Node) NewChild(ns *Namespace, name string, content string) *Node {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	ptrc := C.CString(content)
	defer C.free_string(ptrc)
	if ns != nil {
		return &Node{C.xmlNewChild(node.Ptr, ns.Ptr, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrc))}
	}
	return &Node{C.xmlNewChild(node.Ptr, nil, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrc))}
}

// xmlNewComment
func NewComment(content string) *Node {
	ptr := C.CString(content)
	defer C.free_string(ptr)
	cnode := C.xmlNewComment(C.to_xmlcharptr(ptr))
	return makeNode(cnode)
}

// xmlNewDoc
func NewDoc(version string) *Document {
	ptr := C.CString(version)
	defer C.free_string(ptr)
	cdoc := C.xmlNewDoc(C.to_xmlcharptr(ptr))
	return makeDoc(cdoc)
}

// xmlNewDocComment
func (doc *Document) NewComment(content string) *Node {
	ptr := C.CString(content)
	defer C.free_string(ptr)
	return makeNode(C.xmlNewDocComment(doc.Ptr, C.to_xmlcharptr(ptr)))
}

// xmlNewDocFragment
func (doc *Document) NewFragment() *Node {
	return makeNode(C.xmlNewDocFragment(doc.Ptr))
}

// xmlNewDocNode
func (doc *Document) NewNode(ns *Namespace, name string, content string) *Node {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	ptrc := C.CString(content)
	defer C.free_string(ptrc)
	if ns != nil {
		return &Node{C.xmlNewDocNode(doc.Ptr, ns.Ptr, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrc))}
	}
	return &Node{C.xmlNewDocNode(doc.Ptr, nil, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrc))}
}

// xmlNewDocProp
func (doc *Document) NewProp(name string, value string) *Attribute {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	ptrv := C.CString(value)
	defer C.free_string(ptrv)
	cattr := C.xmlNewDocProp(doc.Ptr, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrv))
	return makeAttribute(cattr)
}

// xmlNewDocRawNode
func (doc *Document) NewRawNode(ns *Namespace, name string, content string) *Node {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	ptrc := C.CString(content)
	defer C.free_string(ptrc)
	if ns != nil {
		return &Node{C.xmlNewDocRawNode(doc.Ptr, ns.Ptr, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrc))}
	}
	return &Node{C.xmlNewDocRawNode(doc.Ptr, nil, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrc))}
}

// xmlNewDocText
func (doc *Document) NewText(content string) *TextNode {
	ptr := C.CString(content)
	defer C.free_string(ptr)
	return &TextNode{makeNode(C.xmlNewDocText(doc.Ptr, C.to_xmlcharptr(ptr)))}
}

// xmlNewDtd
func (doc *Document) NewDtd(name string, ExternalID string, SystemID string) *Dtd {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	ptre := C.CString(ExternalID)
	defer C.free_string(ptre)
	ptrs := C.CString(SystemID)
	defer C.free_string(ptrs)
	cdtd := C.xmlNewDtd(doc.Ptr, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptre), C.to_xmlcharptr(ptrs))
	return makeDtd(cdtd)
}

// xmlNewNode
func NewNode(ns *Namespace, name string) *Node {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	if ns != nil {
		return &Node{C.xmlNewNode(ns.Ptr, C.to_xmlcharptr(ptr))}
	}
	return &Node{C.xmlNewNode(nil, C.to_xmlcharptr(ptr))}
}

// xmlNewNs
func (node *Node) NewNs(href string, prefix string) *Namespace {
	ptrh := C.CString(href)
	defer C.free_string(ptrh)
	ptrp := C.CString(prefix)
	defer C.free_string(ptrp)
	return &Namespace{C.xmlNewNs(node.Ptr, C.to_xmlcharptr(ptrh), C.to_xmlcharptr(ptrp))}
}

// xmlNewProp
func (node *Node) NewAttribute(name string, value string) *Attribute {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	ptrv := C.CString(value)
	defer C.free_string(ptrv)
	cattr := C.xmlNewProp(node.Ptr, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrv))
	return makeAttribute(cattr)
}

// xmlNewText
func NewText(content string) *TextNode {
	ptr := C.CString(content)
	defer C.free_string(ptr)
	return &TextNode{makeNode(C.xmlNewText(C.to_xmlcharptr(ptr)))}
}

// xmlNewTextChild
func (node *Node) NewTextChild(ns *Namespace, name string, content string) *TextNode {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	ptrc := C.CString(content)
	defer C.free_string(ptrc)
	if ns == nil {
		return &TextNode{&Node{C.xmlNewTextChild(node.Ptr, nil, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrc))}}
	}
	return &TextNode{&Node{C.xmlNewTextChild(node.Ptr, ns.Ptr, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrc))}}
}

// xmlNextElementSibling
func (node *Node) NextSibling() *Node {
	return makeNode(C.xmlNextElementSibling(node.Ptr))
}

// xmlNodeAddContent
func (node *Node) AddContent(content string) {
	ptr := C.CString(content)
	defer C.free_string(ptr)
	C.xmlNodeAddContent(node.Ptr, C.to_xmlcharptr(ptr))
}

// xmlNodeDump
func (doc *Document) NodeDump(buf *Buffer, cur *Node, level int, format int) int {
	return int(C.xmlNodeDump(buf.Ptr, doc.Ptr, cur.Ptr, C.int(level), C.int(format)))
}

// xmlNodeGetContent
func (node *Node) GetContent() string {
	return C.GoString(C.to_charptr(C.xmlNodeGetContent(node.Ptr)))
}

// xmlNodeSetContent
func (node *Node) SetContent(content string) {
	ptr := C.CString(content)
	defer C.free_string(ptr)
	C.xmlNodeSetContent(node.Ptr, C.to_xmlcharptr(ptr))
}

// xmlNodeSetName
func (node *Node) SetName(name string) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.xmlNodeSetName(node.Ptr, C.to_xmlcharptr(ptr))
}

// xmlPreviousElementSibling
func (node *Node) PreviousSibling() *Node {
	return makeNode(C.xmlPreviousElementSibling(node.Ptr))
}

// xmlSetCompressMode
func SetCompressionLevel(level int) {
	C.xmlSetCompressMode(C.int(level))
}

// xmlSetDocCompressMode
func (doc *Document) SetCompressionLevel(level int) {
	C.xmlSetDocCompressMode(doc.Ptr, C.int(level))
}

// xmlSetProp
func (node *Node) SetAttribute(name string, value string) *Attribute {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	ptrv := C.CString(value)
	defer C.free_string(ptrv)
	cattr := C.xmlSetProp(node.Ptr, C.to_xmlcharptr(ptrn), C.to_xmlcharptr(ptrv))
	return makeAttribute(cattr)
}

// xmlTextConcat
func (node *TextNode) Concat(content string) int {
	ptr := C.CString(content)
	defer C.free_string(ptr)
	return int(C.xmlTextConcat(node.Ptr, C.to_xmlcharptr(ptr), C.int(len(content))))
}

// xmlTextMerge
func (first *TextNode) Merge(second *Node) *Node {
	return makeNode(C.xmlTextMerge(first.Ptr, second.Ptr))
}

// xmlUnlinkNode
func (node *Node) Unlink() {
	C.xmlUnlinkNode(node.Ptr)
}

// xmlUnsetProp
func (node *Node) UnsetProp(name string) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.xmlUnsetProp(node.Ptr, C.to_xmlcharptr(ptr))
}
