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

// htmlGetMetaEncoding
func (doc *HTMLDocument) GetMetaEncoding() string {
	cstr := C.htmlGetMetaEncoding(doc.Ptr)
	return C.GoString(C.to_charptr(cstr))
}

// htmlIsBooleanAttr
func IsBooleanAttr(name string) bool {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	return C.htmlIsBooleanAttr(C.to_xmlcharptr(ptr)) != 0
}

// htmlNewDoc
func NewHTMLDoc(uri string, external_id string) *HTMLDocument {
	ptru := C.CString(uri)
	defer C.free_string(ptru)
	ptre := C.CString(external_id)
	defer C.free_string(ptre)
	doc := C.htmlNewDoc(C.to_xmlcharptr(ptru), C.to_xmlcharptr(ptre))
	return makeHTMLDoc(doc)
}

// htmlNewDocNoDtD
func NewHTMLDocNoDtD() *HTMLDocument {
	doc := C.htmlNewDocNoDtD(nil, nil)
	return makeHTMLDoc(doc)
}

// htmlNodeDump
func (doc *HTMLDocument) NodeDump(buf *Buffer, cur *HTMLNode) int {
	return int(C.htmlNodeDump(buf.Ptr, doc.Ptr, cur.Ptr))
}

// htmlSaveFile
func (doc *HTMLDocument) SaveFile(filename string) int {
	ptrf := C.CString(filename)
	defer C.free_string(ptrf)
	return int(C.htmlSaveFile(ptrf, doc.Ptr))
}

// htmlSaveFileEnc
func (doc *HTMLDocument) SaveFileEnc(filename string, encoding string) int {
	ptrf := C.CString(filename)
	defer C.free_string(ptrf)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	return int(C.htmlSaveFileEnc(ptrf, doc.Ptr, ptre))
}

// htmlSaveFileFormat
func (doc *HTMLDocument) SaveFileFormat(filename string, encoding string, format int) int {
	ptrf := C.CString(filename)
	defer C.free_string(ptrf)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	return int(C.htmlSaveFileFormat(ptrf, doc.Ptr, ptre, C.int(format)))
}

// htmlSetMetaEncoding
func (doc *HTMLDocument) SetMetaEncoding(encoding string) int {
	ptr := C.CString(encoding)
	defer C.free_string(ptr)
	return int(C.htmlSetMetaEncoding(doc.Ptr, C.to_xmlcharptr(ptr)))
}
