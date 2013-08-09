package golibxml

/*
#cgo pkg-config: libxml-2.0
#include <libxml/tree.h>

static inline void free_string(char* s) { free(s); }
static inline xmlChar *to_xmlcharptr(const char *s) { return (xmlChar *)s; }
static inline char *to_charptr(const xmlChar *s) { return (char *)s; }
*/
import "C"
import "unsafe"

func (doc *Document) String() string {
	buf := NewBuffer()
	defer buf.Free()
	doc.NodeDump(buf, doc.Node, 0, 0)
	return buf.Content()
}

func (node *Node) Document() *Document {
	return makeDoc(_Ctype_xmlDocPtr(unsafe.Pointer(node.Ptr.doc)))
}

func (node *Node) String() string {
	buf := NewBuffer()
	defer buf.Free()
	node.Document().NodeDump(buf, node, 0, 0)
	return buf.Content()
}

func (node *Node) Children() *Node {
	return makeNode(_Ctype_xmlNodePtr(unsafe.Pointer(node.Ptr.children)))
}

func (node *Node) Type() ElementType {
	return ElementType(node.Ptr._type)
}

func (node *Node) Name() string {
	return C.GoString(C.to_charptr(node.Ptr.name))
}

func (node *Node) Next() *Node {
	return makeNode(_Ctype_xmlNodePtr(unsafe.Pointer(node.Ptr.next)))
}

func (node *Node) Attributes() *Attribute {
	return makeAttribute(_Ctype_xmlAttrPtr(unsafe.Pointer(node.Ptr.properties)))
}

func (node *Node) Namespace() *Namespace {
	return makeNamespace(_Ctype_xmlNsPtr(unsafe.Pointer(node.Ptr.ns)))
}

func (attr *Attribute) Type() ElementType {
	return ElementType(attr.Ptr._type)
}

func (attr *Attribute) Name() string {
	return C.GoString(C.to_charptr(attr.Ptr.name))
}

func (attr *Attribute) Children() *Node {
	return makeNode(_Ctype_xmlNodePtr(unsafe.Pointer(attr.Ptr.children)))
}

func (attr *Attribute) Next() *Attribute {
	return makeAttribute(_Ctype_xmlAttrPtr(unsafe.Pointer(attr.Ptr.next)))
}

func (attr *Attribute) Namespace() *Namespace {
	return makeNamespace(_Ctype_xmlNsPtr(unsafe.Pointer(attr.Ptr.ns)))
}

func (ns *Namespace) Href() string {
	return C.GoString(C.to_charptr(ns.Ptr.href))
}

func (ns *Namespace) Prefix() string {
	return C.GoString(C.to_charptr(ns.Ptr.prefix))
}

func (elem ElementType) GoString() string {
	return elem.String()
}

func (elem ElementType) String() string {
	switch (elem) {
		case XML_ELEMENT_NODE:
			return "Node"
		case XML_ATTRIBUTE_NODE:
			return "Attribute"
		case XML_TEXT_NODE:
			return "Text"
	}
	return "Unknown Type"
}

