package xmltree
/*
#cgo pkg-config: libxml-2.0
#include <libxml/tree.h>

static inline void free_string(char* s) { free(s); }
static inline xmlChar *to_xmlcharptr(const char *s) { return (xmlChar *)s; }

*/
import "C"
//import "unsafe"

type xmlNodePtr struct {
	Ptr C.xmlNodePtr
}

type xmlDocPtr struct {
	Ptr C.xmlDocPtr
}

// xmlAddChild
func (parent xmlNodePtr) xmlAddChild(cur xmlNodePtr) (xmlNodePtr) {
	return xmlNodePtr{C.xmlAddChild(parent.Ptr, cur.Ptr)}
}

// xmlAddChildList
func (parent xmlNodePtr) xmlAddChildList(cur xmlNodePtr) (xmlNodePtr) {
	return xmlNodePtr{C.xmlAddNextSibling(parent.Ptr, cur.Ptr)}
}

// xmlAddNextSibling
func (cur xmlNodePtr) xmlAddNextSibling(elem xmlNodePtr) (xmlNodePtr) {
	return xmlNodePtr{C.xmlAddNextSibling(cur.Ptr, elem.Ptr)}
}

// xmlAddPrevSibling
func (cur xmlNodePtr) xmlAddPrevSibling(elem xmlNodePtr) (xmlNodePtr) {
	return xmlNodePtr{C.xmlAddPrevSibling(cur.Ptr, elem.Ptr)}
}

// xmlAddSibling
func (cur xmlNodePtr) xmlAddSibling(elem xmlNodePtr) (xmlNodePtr) {
	return xmlNodePtr{C.xmlAddSibling(cur.Ptr, elem.Ptr)}
}

// xmlNewComment
func xmlNewComment(content string) (xmlNodePtr) {
	ptr := C.CString(content)
	defer C.free_string(ptr)
	return xmlNodePtr{C.xmlNewComment(C.to_xmlcharptr(ptr))}
}

// xmlNewDoc
func xmlNewDoc(version string) (xmlDocPtr) {
	ptr := C.CString(version)
	defer C.free_string(ptr)
	return xmlDocPtr{C.xmlNewDoc(C.to_xmlcharptr(ptr))}
}

// xmlNewDocComment
func xmlNewDocComment(doc xmlDocPtr, content string) (xmlNodePtr) {
	ptr := C.CString(content)
	defer C.free_string(ptr)
	return xmlNodePtr{C.xmlNewDocComment(doc.Ptr, C.to_xmlcharptr(ptr))}
}

// xmlNewDocFragment
func xmlNewDocFragment(doc xmlDocPtr) (xmlNodePtr) {
	return xmlNodePtr{C.xmlNewDocFragment(doc.Ptr)}
}
