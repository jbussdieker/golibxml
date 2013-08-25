package golibxml

/*
#cgo pkg-config: libxml-2.0
#include <libxml/xmlreader.h>

static inline void free_string(char* s) { free(s); }
static inline xmlChar *to_xmlcharptr(const char *s) { return (xmlChar *)s; }
static inline char *to_charptr(const xmlChar *s) { return (char *)s; }
*/
import "C"

////////////////////////////////////////////////////////////////////////////////
// TYPES/STRUCTS
////////////////////////////////////////////////////////////////////////////////

type TextReader struct {
	Ptr    C.xmlTextReaderPtr
	buffer *C.char
}

// xmlReaderTypes
type ReaderType int

const (
	XML_READER_TYPE_NONE                   ReaderType = C.XML_READER_TYPE_NONE
	XML_READER_TYPE_ELEMENT                           = C.XML_READER_TYPE_ELEMENT
	XML_READER_TYPE_ATTRIBUTE                         = C.XML_READER_TYPE_ATTRIBUTE
	XML_READER_TYPE_TEXT                              = C.XML_READER_TYPE_TEXT
	XML_READER_TYPE_CDATA                             = C.XML_READER_TYPE_CDATA
	XML_READER_TYPE_ENTITY_REFERENCE                  = C.XML_READER_TYPE_ENTITY_REFERENCE
	XML_READER_TYPE_ENTITY                            = C.XML_READER_TYPE_ENTITY
	XML_READER_TYPE_PROCESSING_INSTRUCTION            = C.XML_READER_TYPE_PROCESSING_INSTRUCTION
	XML_READER_TYPE_COMMENT                           = C.XML_READER_TYPE_COMMENT
	XML_READER_TYPE_DOCUMENT                          = C.XML_READER_TYPE_DOCUMENT
	XML_READER_TYPE_DOCUMENT_TYPE                     = C.XML_READER_TYPE_DOCUMENT_TYPE
	XML_READER_TYPE_DOCUMENT_FRAGMENT                 = C.XML_READER_TYPE_DOCUMENT_FRAGMENT
	XML_READER_TYPE_NOTATION                          = C.XML_READER_TYPE_NOTATION
	XML_READER_TYPE_WHITESPACE                        = C.XML_READER_TYPE_WHITESPACE
	XML_READER_TYPE_SIGNIFICANT_WHITESPACE            = C.XML_READER_TYPE_SIGNIFICANT_WHITESPACE
	XML_READER_TYPE_END_ELEMENT                       = C.XML_READER_TYPE_END_ELEMENT
	XML_READER_TYPE_END_ENTITY                        = C.XML_READER_TYPE_END_ENTITY
	XML_READER_TYPE_XML_DECLARATION                   = C.XML_READER_TYPE_XML_DECLARATION
)

////////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func makeTextReader(reader C.xmlTextReaderPtr, buffer *C.char) *TextReader {
	if reader == nil {
		return nil
	}
	return &TextReader{reader, buffer}
}

////////////////////////////////////////////////////////////////////////////////
// INTERFACE
////////////////////////////////////////////////////////////////////////////////

func (t ReaderType) String() string {
	switch t {
	case XML_READER_TYPE_NONE:
		return "None"
	case XML_READER_TYPE_ELEMENT:
		return "Element"
	case XML_READER_TYPE_ATTRIBUTE:
		return "Attribute"
	case XML_READER_TYPE_TEXT:
		return "Text"
	case XML_READER_TYPE_CDATA:
		return "CData"
	case XML_READER_TYPE_ENTITY_REFERENCE:
		return "EntityReference"
	case XML_READER_TYPE_ENTITY:
		return "Entity"
	case XML_READER_TYPE_PROCESSING_INSTRUCTION:
		return "ProcessingInstruction"
	case XML_READER_TYPE_COMMENT:
		return "Comment"
	case XML_READER_TYPE_DOCUMENT:
		return "Document"
	case XML_READER_TYPE_DOCUMENT_TYPE:
		return "DocumentType"
	case XML_READER_TYPE_DOCUMENT_FRAGMENT:
		return "DocumentFragment"
	case XML_READER_TYPE_NOTATION:
		return "Notation"
	case XML_READER_TYPE_WHITESPACE:
		return "Whitespace"
	case XML_READER_TYPE_SIGNIFICANT_WHITESPACE:
		return "SignificantWhitespace"
	case XML_READER_TYPE_END_ELEMENT:
		return "EndElement"
	case XML_READER_TYPE_END_ENTITY:
		return "EndEntity"
	case XML_READER_TYPE_XML_DECLARATION:
		return "XMLDeclaration"
	default:
		return "Unknown"
	}
}

// xmlReaderForMemory
func ReaderForMemory(buffer string, url, encoding string, options ParserOption) *TextReader {

	cbuffer := C.CString(buffer)

	clen := C.int(len(buffer))

	curl := C.CString(url)
	defer C.free_string(curl)

	cencoding := C.CString(encoding)
	defer C.free_string(cencoding)

	coptions := C.int(options)

	ptr := C.xmlReaderForMemory(cbuffer, clen, curl, cencoding, coptions)

	return makeTextReader(ptr, cbuffer)
}

// xmlTextReaderRead
func (r *TextReader) Read() int {
	return int(C.xmlTextReaderRead(r.Ptr))
}

// xmlTextReaderNext
func (r *TextReader) Next() int {
	return int(C.xmlTextReaderNext(r.Ptr))
}

// xmlTextReaderMoveToFirstAttribute
func (r *TextReader) MoveToFirstAttribute() int {
	return int(C.xmlTextReaderMoveToFirstAttribute(r.Ptr))
}

// xmlTextReaderMoveToNextAttribute
func (r *TextReader) MoveToNextAttribute() int {
	return int(C.xmlTextReaderMoveToNextAttribute(r.Ptr))
}

// xmlTextReaderNodeType
func (r *TextReader) NodeType() ReaderType {
	return ReaderType(C.xmlTextReaderNodeType(r.Ptr))
}

// xmlTextReaderConstLocalName
func (r *TextReader) LocalName() string {
	return C.GoString(C.to_charptr(C.xmlTextReaderConstLocalName(r.Ptr)))
}

// xmlTextReaderConstName
func (r *TextReader) Name() string {
	return C.GoString(C.to_charptr(C.xmlTextReaderConstName(r.Ptr)))
}

// xmlTextReaderValue
func (r *TextReader) Value() string {
	cvalue := C.to_charptr(C.xmlTextReaderValue(r.Ptr))
	defer C.free_string(cvalue)
	return C.GoString(cvalue)
}

// xmlTextReaderConstNamespaceUri
func (r *TextReader) NamespaceUri() string {
	return C.GoString(C.to_charptr(C.xmlTextReaderConstNamespaceUri(r.Ptr)))
}

// xmlTextReaderConstPrefix
func (r *TextReader) Prefix() string {
	return C.GoString(C.to_charptr(C.xmlTextReaderConstPrefix(r.Ptr)))
}

// xmlTextReaderReadString
func (r *TextReader) ReadString() string {
	cstr := C.to_charptr(C.xmlTextReaderReadString(r.Ptr))
	defer C.free_string(cstr)
	return C.GoString(cstr)
}

// xmlTextReaderDepth
func (r *TextReader) Depth() int {
	return int(C.xmlTextReaderDepth(r.Ptr))
}

// xmlTextReaderGetAttribute
func (r *TextReader) GetAttribute(name string) *string {
	cname := C.CString(name)
	defer C.free_string(cname)

	cvalue := C.xmlTextReaderGetAttribute(r.Ptr, C.to_xmlcharptr(cname))

	if cvalue != nil {
		value := C.GoString(C.to_charptr(cvalue))
		return &value
	}
	return nil
}

// xmlFreeTextReader
func (r *TextReader) Free() {
	C.xmlFreeTextReader(r.Ptr)
	C.free_string(r.buffer)
}
