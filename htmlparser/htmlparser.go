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
// INTERFACE
////////////////////////////////////////////////////////////////////////////////

func ParseDoc(cur string, encoding string) *htmltree.Document {
	ptrc := C.CString(cur)
	defer C.free_string(ptrc)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	doc := C.htmlParseDoc(C.to_xmlcharptr(ptrc), ptre)
	dp := xmltree.DocumentPtr(unsafe.Pointer(doc))
	np := xmltree.NodePtr(unsafe.Pointer(doc))
	return &htmltree.Document{
		&xmltree.Document{
			Ptr: dp,
			Node: &xmltree.Node{np},
		},
		htmltree.DocumentPtr(unsafe.Pointer(doc)),
	}
}

