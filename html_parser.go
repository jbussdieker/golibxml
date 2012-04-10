package golibxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/HTMLparser.h>

static inline void free_string(char* s) { free(s); }
static inline xmlChar *to_xmlcharptr(const char *s) { return (xmlChar *)s; }
static inline char *to_charptr(const xmlChar *s) { return (char *)s; }

*/
import "C"
import "unsafe"

////////////////////////////////////////////////////////////////////////////////
// TYPES/STRUCTS
////////////////////////////////////////////////////////////////////////////////

type ElemDesc struct {
	Ptr C.htmlElemDescPtr
}

type HTMLDocument struct {
	*Document
	*HTMLNode
	Ptr C.htmlDocPtr
}

////////////////////////////////////////////////////////////////////////////////
// INTERFACE
////////////////////////////////////////////////////////////////////////////////

// htmlParseDoc
func ParseHTMLDoc(cur string, encoding string) *HTMLDocument {
	ptrc := C.CString(cur)
	defer C.free_string(ptrc)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	doc := C.htmlParseDoc(C.to_xmlcharptr(ptrc), ptre)
	return &HTMLDocument{
		Ptr: doc,
		Document: &Document{
			Ptr:  C.xmlDocPtr(doc),
			Node: &Node{C.xmlNodePtr(unsafe.Pointer(doc))},
		},
		HTMLNode: &HTMLNode{
			Ptr:  C.htmlNodePtr(unsafe.Pointer(doc)),
			Node: &Node{C.xmlNodePtr(unsafe.Pointer(doc))},
		},
	}
}

// htmlAutoCloseTag
func (doc *HTMLDocument) AutoCloseTag(name string, node *Node) int {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	return int(C.htmlAutoCloseTag(doc.Document.Ptr, C.to_xmlcharptr(ptr), node.Ptr))
}

// htmlTagLookup
func TagLookup(tag string) *ElemDesc {
	ptr := C.CString(tag)
	defer C.free_string(ptr)
	return &ElemDesc{C.htmlTagLookup(C.to_xmlcharptr(ptr))}
}

