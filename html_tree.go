package golibxml
/*
#cgo pkg-config: libxml-2.0
#include <libxml/HTMLparser.h>
#include <libxml/HTMLtree.h>

static inline void free_string(char* s) { free(s); }
static inline xmlChar *to_xmlcharptr(const char *s) { return (xmlChar *)s; }
static inline char *to_charptr(const xmlChar *s) { return (char *)s; }

*/
import "C"

////////////////////////////////////////////////////////////////////////////////
// TYPES/STRUCTS
////////////////////////////////////////////////////////////////////////////////

type HTMLNode struct {
	*Node
	Ptr C.htmlNodePtr
}

////////////////////////////////////////////////////////////////////////////////
// INTERFACE
////////////////////////////////////////////////////////////////////////////////

// htmlNodeDump
func (doc *HTMLDocument) NodeDump(buf *Buffer, cur *HTMLNode) int {
	return int(C.htmlNodeDump(buf.Ptr, doc.Ptr, cur.Ptr))
}

