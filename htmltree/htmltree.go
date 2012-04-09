package htmltree
/*
#cgo pkg-config: libxml-2.0
#include <libxml/HTMLparser.h>
#include <libxml/HTMLtree.h>

static inline void free_string(char* s) { free(s); }
static inline xmlChar *to_xmlcharptr(const char *s) { return (xmlChar *)s; }
static inline char *to_charptr(const xmlChar *s) { return (char *)s; }

*/
import "C"
import "unsafe"

import "github.com/jbussdieker/golibxml/xmltree"

////////////////////////////////////////////////////////////////////////////////
// TYPES/STRUCTS
////////////////////////////////////////////////////////////////////////////////

type DocumentPtr C.htmlDocPtr
type Document struct {
	*xmltree.Document
	Ptr DocumentPtr
}

////////////////////////////////////////////////////////////////////////////////
// INTERFACE
////////////////////////////////////////////////////////////////////////////////

// htmlNodeDump
func (doc *Document) NodeDump(buf *xmltree.Buffer, cur *xmltree.Node) int {
	node := C.xmlNodePtr(unsafe.Pointer(cur.Ptr))
	bp := C.xmlBufferPtr(unsafe.Pointer(buf.Ptr))
	return int(C.htmlNodeDump(bp, doc.Ptr, node))
}

