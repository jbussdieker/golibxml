package htmlparser
/*
#cgo pkg-config: libxml-2.0
#include <libxml/HTMLparser.h>

static inline void free_string(char* s) { free(s); }
static inline xmlChar *to_xmlcharptr(const char *s) { return (xmlChar *)s; }
static inline char *to_charptr(const xmlChar *s) { return (char *)s; }

*/
import "C"
import "unsafe"

import "github.com/jbussdieker/golibxml/xmltree"
import "github.com/jbussdieker/golibxml/htmltree"

////////////////////////////////////////////////////////////////////////////////
// TYPES/STRUCTS
////////////////////////////////////////////////////////////////////////////////

type ElemDesc struct {
	Ptr C.htmlElemDescPtr
}

type Document struct {
	*htmltree.Document
}

type Node struct {
	*htmltree.Node
}

////////////////////////////////////////////////////////////////////////////////
// INTERFACE
////////////////////////////////////////////////////////////////////////////////

// htmlParseDoc
func ParseDoc(cur string, encoding string) *Document {
	ptrc := C.CString(cur)
	defer C.free_string(ptrc)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	doc := C.htmlParseDoc(C.to_xmlcharptr(ptrc), ptre)
	dp := xmltree.DocumentPtr(unsafe.Pointer(doc))
	np := xmltree.NodePtr(unsafe.Pointer(doc))
	return &Document{
		&htmltree.Document{
			&xmltree.Document{
				Ptr: dp,
				Node: &xmltree.Node{np},
			},
			htmltree.DocumentPtr(unsafe.Pointer(doc)),
		},
	}
}

// htmlAutoCloseTag
func (doc *Document) AutoCloseTag(name string, node *Node) int {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	d := doc.Document.Ptr
	n := node.Node.Ptr
	return int(C.htmlAutoCloseTag(C.xmlDocPtr(unsafe.Pointer(d)), C.to_xmlcharptr(ptr), C.xmlNodePtr(unsafe.Pointer(n))))
}

// htmlTagLookup
func TagLookup(tag string) *ElemDesc {
	ptr := C.CString(tag)
	defer C.free_string(ptr)
	return &ElemDesc{C.htmlTagLookup(C.to_xmlcharptr(ptr))}
}

