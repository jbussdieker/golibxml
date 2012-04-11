package golibxml

/*
#cgo pkg-config: libxml-2.0
#include <libxml/xpath.h>

static inline void free_string(char* s) { free(s); }
static inline xmlChar *to_xmlcharptr(const char *s) { return (xmlChar *)s; }
static inline char *to_charptr(const xmlChar *s) { return (char *)s; }

*/
import "C"
import "unsafe"

////////////////////////////////////////////////////////////////////////////////
// TYPES/STRUCTS
////////////////////////////////////////////////////////////////////////////////

type Xpath struct {
	Ptr C.xmlXPathCompExprPtr
}

type XpathCtx struct {
	Ptr C.xmlXPathContextPtr
}

type XPathObjectPtr struct {
	Ptr C.xmlXPathObjectPtr
}

type NodeSet struct {
	Ptr C.xmlNodeSetPtr
}

////////////////////////////////////////////////////////////////////////////////
// INTERFACE
////////////////////////////////////////////////////////////////////////////////

// xmlXPathCastToNumber
func (obj *XPathObjectPtr) CastToNumber() float32 {
	cdbl := C.xmlXPathCastToNumber(obj.Ptr)
	return float32(cdbl)
}

// xmlXPathCastToString
func (obj *XPathObjectPtr) CastToString() string {
	cstr := C.xmlXPathCastToString(obj.Ptr)
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(C.to_charptr(cstr))
}

// xmlXPathCompile
func XPathCompile(str string) *Xpath {
	ptr := C.CString(str)
	defer C.free_string(ptr)
	return &Xpath{C.xmlXPathCompile(C.to_xmlcharptr(ptr))}
}

// xmlXPathConvertBoolean
func (obj *XPathObjectPtr) ConvertBoolean() *XPathObjectPtr {
	return &XPathObjectPtr{C.xmlXPathConvertBoolean(obj.Ptr)}
}

// xmlXPathConvertNumber
func (obj *XPathObjectPtr) ConvertNumber() *XPathObjectPtr {
	return &XPathObjectPtr{C.xmlXPathConvertNumber(obj.Ptr)}
}

// xmlXPathConvertString
func (obj *XPathObjectPtr) ConvertString() *XPathObjectPtr {
	return &XPathObjectPtr{C.xmlXPathConvertString(obj.Ptr)}
}

// xmlXPathCtxtCompile
func (ctx *XpathCtx) XPathCompile(str string) *Xpath {
	ptr := C.CString(str)
	defer C.free_string(ptr)
	return &Xpath{C.xmlXPathCtxtCompile(ctx.Ptr, C.to_xmlcharptr(ptr))}
}

// xmlXPathEvalExpression
func (ctx *XpathCtx) EvalExpression(str string) *XPathObjectPtr {
	ptr := C.CString(str)
	defer C.free_string(ptr)
	return &XPathObjectPtr{C.xmlXPathEvalExpression(C.to_xmlcharptr(ptr), ctx.Ptr)}
}

// xmlXPathFreeCompExpr
func (xpath *Xpath) Free() {
	C.xmlXPathFreeCompExpr(xpath.Ptr)
	xpath.Ptr = nil
}

// xmlXPathFreeContext
func (ctx *XpathCtx) Free() {
	C.xmlXPathFreeContext(ctx.Ptr)
	ctx.Ptr = nil
}

// xmlXPathFreeNodeSet
func (nodeset *NodeSet) Free() {
	C.xmlXPathFreeNodeSet(nodeset.Ptr)
	nodeset.Ptr = nil
}

// xmlXPathIsInf
func xmlXPathIsInf(val float32) int {
	return int(C.xmlXPathIsInf(C.double(val)))
}

// xmlXPathIsNaN
func XpathIsNaN(val float32) int {
	return int(C.xmlXPathIsNaN(C.double(val)))
}

// xmlXPathNewContext
func XpathNewContext(doc *Document) *XpathCtx {
	return &XpathCtx{C.xmlXPathNewContext(doc.Ptr)}
}

// xmlXPathNodeSetCreate
func NodeSetCreate(node *Node) *NodeSet {
	return &NodeSet{C.xmlXPathNodeSetCreate(node.Ptr)}
}

// xmlXPathObjectCopy
func (obj *XPathObjectPtr) Copy() *XPathObjectPtr {
	return &XPathObjectPtr{C.xmlXPathObjectCopy(obj.Ptr)}
}

// xmlXPathOrderDocElems
func (doc *Document) OrderDocElems() int {
	return int(C.xmlXPathOrderDocElems(doc.Ptr))
}

