package golibxml

/*
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

type HTMLParserOption int

const (
	HTML_PARSE_RECOVER   HTMLParserOption = C.HTML_PARSE_RECOVER   //: Relaxed parsing
	HTML_PARSE_NODEFDTD                   = C.HTML_PARSE_NODEFDTD  //: do not default a doctype if not found
	HTML_PARSE_NOERROR                    = C.HTML_PARSE_NOERROR   //: suppress error reports
	HTML_PARSE_NOWARNING                  = C.HTML_PARSE_NOWARNING //: suppress warning reports
	HTML_PARSE_PEDANTIC                   = C.HTML_PARSE_PEDANTIC  //: pedantic error reporting
	HTML_PARSE_NOBLANKS                   = C.HTML_PARSE_NOBLANKS  //: remove blank nodes
	HTML_PARSE_NONET                      = C.HTML_PARSE_NONET     //: Forbid network access
	HTML_PARSE_NOIMPLIED                  = C.HTML_PARSE_NOIMPLIED //: Do not add implied html/body... elements
	HTML_PARSE_COMPACT                    = C.HTML_PARSE_COMPACT   //: compact small text nodes
)

type ElemDesc struct {
	Ptr C.htmlElemDescPtr
}

type HTMLDocument struct {
	*Document
	*HTMLNode
	Ptr C.htmlDocPtr
}

type HTMLParser struct {
	Ptr C.htmlParserCtxtPtr
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func makeHTMLDoc(doc C.htmlDocPtr) *HTMLDocument {
	if doc == nil {
		return nil
	}
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

func makeHTMLParser(parser C.htmlParserCtxtPtr) *HTMLParser {
	if parser == nil {
		return nil
	}
	return &HTMLParser{parser}
}

func makeElemDesc(desc C.htmlElemDescPtr) *ElemDesc {
	if desc == nil {
		return nil
	}
	return &ElemDesc{desc}
}

////////////////////////////////////////////////////////////////////////////////
// INTERFACE
////////////////////////////////////////////////////////////////////////////////

// htmlAutoCloseTag
func (doc *HTMLDocument) AutoCloseTag(name string, node *Node) bool {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	return int(C.htmlAutoCloseTag(doc.Document.Ptr, C.to_xmlcharptr(ptr), node.Ptr)) == 1
}

// htmlCtxtReadDoc
func (p *HTMLParser) ReadDoc(input string, url string, encoding string, options ParserOption) *HTMLDocument {
	ptri := C.CString(input)
	defer C.free_string(ptri)
	ptru := C.CString(url)
	defer C.free_string(ptru)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	doc := C.htmlCtxtReadDoc(p.Ptr, C.to_xmlcharptr(ptri), ptru, ptre, C.int(options))
	return makeHTMLDoc(doc)
}

// htmlCtxtReset
func (p *HTMLParser) Reset() {
	C.htmlCtxtReset(p.Ptr)
}

// htmlCtxtUseOptions
func (p *HTMLParser) UseOptions(options HTMLParserOption) int {
	return int(C.htmlCtxtUseOptions(p.Ptr, C.int(options)))
}

// htmlFreeParserCtxt
func (p *HTMLParser) Free() {
	C.htmlFreeParserCtxt(p.Ptr)
}

// htmlNewParserCtxt
func NewHTMLParserCtxt() *HTMLParser {
	pctx := C.htmlNewParserCtxt()
	return makeHTMLParser(pctx)
}

// htmlParseDoc
func ParseHTMLDoc(cur string, encoding string) *HTMLDocument {
	ptrc := C.CString(cur)
	defer C.free_string(ptrc)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	doc := C.htmlParseDoc(C.to_xmlcharptr(ptrc), ptre)
	return makeHTMLDoc(doc)
}

// htmlParseFile
func ParseHTMLFile(filename string, encoding string) *HTMLDocument {
	ptrf := C.CString(filename)
	defer C.free_string(ptrf)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	doc := C.htmlParseFile(ptrf, ptre)
	return makeHTMLDoc(doc)
}

// htmlReadDoc
func ReadHTMLDoc(cur string, url string, encoding string, options HTMLParserOption) *HTMLDocument {
	ptrc := C.CString(cur)
	defer C.free_string(ptrc)
	ptru := C.CString(url)
	defer C.free_string(ptru)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	doc := C.htmlReadDoc(C.to_xmlcharptr(ptrc), ptru, ptre, C.int(options))
	return makeHTMLDoc(doc)
}

// htmlReadFile
func ReadHTMLFile(filename string, encoding string, options HTMLParserOption) *HTMLDocument {
	ptrf := C.CString(filename)
	defer C.free_string(ptrf)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	doc := C.htmlReadFile(ptrf, ptre, C.int(options))
	return makeHTMLDoc(doc)
}

// htmlReadMemory
func ReadHTMLMemory(buffer []byte, url string, encoding string, options HTMLParserOption) *HTMLDocument {
	ptru := C.CString(url)
	defer C.free_string(ptru)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	doc := C.htmlReadMemory((*C.char)(unsafe.Pointer(&buffer[0])), C.int(len(buffer)), ptru, ptre, C.int(options))
	return makeHTMLDoc(doc)
}

// htmlTagLookup
func TagLookup(tag string) *ElemDesc {
	ptr := C.CString(tag)
	defer C.free_string(ptr)
	cdesc := C.htmlTagLookup(C.to_xmlcharptr(ptr))
	return makeElemDesc(cdesc)
}
