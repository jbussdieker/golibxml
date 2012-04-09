package golibxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/parser.h>

static inline void free_string(char* s) { free(s); }
static inline xmlChar *to_xmlcharptr(const char *s) { return (xmlChar *)s; }
static inline char *to_charptr(const xmlChar *s) { return (char *)s; }

*/
import "C"
import "unsafe"

////////////////////////////////////////////////////////////////////////////////
// TYPES/STRUCTS
////////////////////////////////////////////////////////////////////////////////

func ParseDoc(cur string) *Document {
	ptr := C.CString(cur)
	defer C.free_string(ptr)
	doc := C.xmlParseDoc(C.to_xmlcharptr(ptr))
	return &Document{
		Ptr: C.xmlDocPtr(doc), 
		Node: &Node{C.xmlNodePtr(unsafe.Pointer(doc))},
	}
}
