package xmlparser
/*
#cgo pkg-config: libxml-2.0
#include <libxml/parser.h>

static inline void free_string(char* s) { free(s); }
static inline xmlChar *to_xmlcharptr(const char *s) { return (xmlChar *)s; }
static inline char *to_charptr(const xmlChar *s) { return (char *)s; }

*/
import "C"
import "unsafe"

import "github.com/jbussdieker/golibxml/xmltree"

////////////////////////////////////////////////////////////////////////////////
// INTERFACE
////////////////////////////////////////////////////////////////////////////////

func ParseDoc(cur string) *xmltree.Document {
	ptr := C.CString(cur)
	defer C.free_string(ptr)
	doc := C.xmlParseDoc(C.to_xmlcharptr(ptr))
	dp := xmltree.DocumentPtr(unsafe.Pointer(doc))
	np := xmltree.NodePtr(unsafe.Pointer(doc))
	return &xmltree.Document{
		Ptr: dp,
		Node: &xmltree.Node{np},
	}
}
