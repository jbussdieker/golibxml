package golibxml

/*
#cgo pkg-config: libxml-2.0
#include <libxml/parser.h>

static inline void free_string(char* s) { free(s); }
static inline xmlChar *to_xmlcharptr(const char *s) { return (xmlChar *)s; }
static inline char *to_charptr(const xmlChar *s) { return (char *)s; }
static inline int *new_int_ptr(int value) { 
	int *ptr = calloc(sizeof(int), 1);
	*ptr = value;
	return ptr;
}
static inline char **new_char_array(int size) { return calloc(sizeof(char *), size); } 
static inline void set_char_array_string(char **ptr, char *str, int n) { ptr[n] = str; } 
static inline char *get_char_array_string(char **ptr, int n) {
	return ptr[n];
}
*/
import "C"
import "unsafe"

////////////////////////////////////////////////////////////////////////////////
// TYPES/STRUCTS
////////////////////////////////////////////////////////////////////////////////

type ParserOption int

const (
	XML_PARSE_RECOVER    ParserOption = C.XML_PARSE_RECOVER    //: recover on errors
	XML_PARSE_NOENT                   = C.XML_PARSE_NOENT      //: substitute entities
	XML_PARSE_DTDLOAD                 = C.XML_PARSE_DTDLOAD    //: load the external subset
	XML_PARSE_DTDATTR                 = C.XML_PARSE_DTDATTR    //: default DTD attributes
	XML_PARSE_DTDVALID                = C.XML_PARSE_DTDVALID   //: validate with the DTD
	XML_PARSE_NOERROR                 = C.XML_PARSE_NOERROR    //: suppress error reports
	XML_PARSE_NOWARNING               = C.XML_PARSE_NOWARNING  //: suppress warning reports
	XML_PARSE_PEDANTIC                = C.XML_PARSE_PEDANTIC   //: pedantic error reporting
	XML_PARSE_NOBLANKS                = C.XML_PARSE_NOBLANKS   //: remove blank nodes
	XML_PARSE_SAX1                    = C.XML_PARSE_SAX1       //: use the SAX1 interface internally
	XML_PARSE_XINCLUDE                = C.XML_PARSE_XINCLUDE   //: Implement XInclude substitition
	XML_PARSE_NONET                   = C.XML_PARSE_NONET      //: Forbid network access
	XML_PARSE_NODICT                  = C.XML_PARSE_NODICT     //: Do not reuse the context dictionnary
	XML_PARSE_NSCLEAN                 = C.XML_PARSE_NSCLEAN    //: remove redundant namespaces declarations
	XML_PARSE_NOCDATA                 = C.XML_PARSE_NOCDATA    //: merge CDATA as text nodes
	XML_PARSE_NOXINCNODE              = C.XML_PARSE_NOXINCNODE //: do not generate XINCLUDE START/END nodes
	XML_PARSE_COMPACT                 = C.XML_PARSE_COMPACT    //: compact small text nodes; no modification of the tree allowed afterwards (will possibly crash if you try to modify the tree)
	XML_PARSE_OLD10                   = C.XML_PARSE_OLD10      //: parse using XML-1.0 before update 5
	XML_PARSE_NOBASEFIX               = C.XML_PARSE_NOBASEFIX  //: do not fixup XINCLUDE xml:base uris
	XML_PARSE_HUGE                    = C.XML_PARSE_HUGE       //: relax any hardcoded limit from the parser
	XML_PARSE_OLDSAX                  = C.XML_PARSE_OLDSAX     //: parse using SAX2 interface from before 2.7.0
)

type Parser struct {
	Ptr C.xmlParserCtxtPtr
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func makeParser(parser C.xmlParserCtxtPtr) *Parser {
	if parser == nil {
		return nil
	}
	return &Parser{parser}
}

////////////////////////////////////////////////////////////////////////////////
// INTERFACE
////////////////////////////////////////////////////////////////////////////////

// xmlByteConsumed
func (p *Parser) ByteConsumed() int {
	return int(C.xmlByteConsumed(p.Ptr))
}

// xmlCleanupParser
func CleanupParser() {
	C.xmlCleanupParser()
}

// xmlCtxtReadDoc
func (p *Parser) ReadDoc(input string, url string, encoding string, options ParserOption) *Document {
	ptri := C.CString(input)
	defer C.free_string(ptri)
	ptru := C.CString(url)
	defer C.free_string(ptru)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	doc := C.xmlCtxtReadDoc(p.Ptr, C.to_xmlcharptr(ptri), ptru, ptre, C.int(options))
	return makeDoc(doc)
}

// xmlCtxtReset
func (p *Parser) Reset() {
	C.xmlCtxtReset(p.Ptr)
}

// xmlCtxtUseOptions
func (p *Parser) UseOptions(options ParserOption) int {
	return int(C.xmlCtxtUseOptions(p.Ptr, C.int(options)))
}

// xmlFreeParserCtxt
func (p *Parser) Free() {
	C.xmlFreeParserCtxt(p.Ptr)
}

// xmlGetFeaturesList
func GetFeaturesList() []string {
	// Get list in C land
	clength := C.new_int_ptr(255)
	defer C.free(unsafe.Pointer(clength))
	clist := C.new_char_array(255)
	defer C.free(unsafe.Pointer(clist))
	result := C.xmlGetFeaturesList(clength, clist)
	if result < 0 {
		panic("ERROR TO BE HANDLED")
	}

	// Convert to Go land
	length := int(*clength)
	list := make([]string, length)
	for i := 0; i < length; i++ {
		list[i] = C.GoString(C.get_char_array_string(clist, C.int(i)))
	}

	return list
}

// xmlNewParserCtxt
func NewParser() *Parser {
	pctx := C.xmlNewParserCtxt()
	return makeParser(pctx)
}

// xmlParseDTD
func ParseDTD(ExternalID string, SystemID string) *Dtd {
	ptre := C.CString(ExternalID)
	defer C.free_string(ptre)
	ptrs := C.CString(SystemID)
	defer C.free_string(ptrs)
	cdtd := C.xmlParseDTD(C.to_xmlcharptr(ptre), C.to_xmlcharptr(ptrs))
	return makeDtd(cdtd)
}

// xmlParseDoc
func ParseDoc(cur string) *Document {
	ptr := C.CString(cur)
	defer C.free_string(ptr)
	doc := C.xmlParseDoc(C.to_xmlcharptr(ptr))
	return makeDoc(doc)
}

// xmlParseDocument
func (p *Parser) Parse() int {
	return int(C.xmlParseDocument(p.Ptr))
}

// xmlParseEntity
func ParseEntity(filename string) *Document {
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	doc := C.xmlParseEntity(ptr)
	return makeDoc(doc)
}

// xmlParseFile
func ParseFile(filename string) *Document {
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	doc := C.xmlParseFile(ptr)
	return makeDoc(doc)
}

// xmlParseMemory
func ParseMemory(buffer []byte) *Document {
	doc := C.xmlParseMemory((*C.char)(unsafe.Pointer(&buffer[0])), C.int(len(buffer)))
	return makeDoc(doc)
}

// xmlReadDoc
func ReadDoc(input string, url string, encoding string, options ParserOption) *Document {
	ptri := C.CString(input)
	defer C.free_string(ptri)
	ptru := C.CString(url)
	defer C.free_string(ptru)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	doc := C.xmlReadDoc(C.to_xmlcharptr(ptri), ptru, ptre, C.int(options))
	return makeDoc(doc)
}

// xmlReadFile
func ReadFile(filename string, encoding string, options ParserOption) *Document {
	ptrf := C.CString(filename)
	defer C.free_string(ptrf)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	doc := C.xmlReadFile(ptrf, ptre, C.int(options))
	return makeDoc(doc)
}

// xmlReadMemory
func ReadMemory(buffer []byte, url string, encoding string, options ParserOption) *Document {
	ptru := C.CString(url)
	defer C.free_string(ptru)
	ptre := C.CString(encoding)
	defer C.free_string(ptre)
	doc := C.xmlReadMemory((*C.char)(unsafe.Pointer(&buffer[0])), C.int(len(buffer)), ptru, ptre, C.int(options))
	return makeDoc(doc)
}

// xmlRecoverDoc
func RecoverDoc(cur string) *Document {
	ptr := C.CString(cur)
	defer C.free_string(ptr)
	doc := C.xmlRecoverDoc(C.to_xmlcharptr(ptr))
	return makeDoc(doc)
}

// xmlRecoverFile
func RecoverFile(filename string) *Document {
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	doc := C.xmlRecoverFile(ptr)
	return makeDoc(doc)
}

// xmlRecoverMemory
func RecoverMemory(buffer []byte) *Document {
	doc := C.xmlRecoverMemory((*C.char)(unsafe.Pointer(&buffer[0])), C.int(len(buffer)))
	return makeDoc(doc)
}

// xmlStopParser
func (p *Parser) Stop() {
	C.xmlStopParser(p.Ptr)
}

// xmlSubstituteEntitiesDefault
func SubstituteEntitiesDefault(val int) int {
	return int(C.xmlSubstituteEntitiesDefault(C.int(val)))
}
