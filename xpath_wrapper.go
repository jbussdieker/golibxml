package golibxml

/*
#cgo pkg-config: libxml-2.0
#include <libxml/xpath.h>

xmlNode* fetchNode(xmlNodeSet *nodeset, int index) {
  	return nodeset->nodeTab[index];
}
*/
import "C"

func (obj *XPathObject) Type() XpathObjectType {
	return XpathObjectType(obj.Ptr._type)
}

func (obj *XPathObject) NodeSet() *NodeSet {
	return makeNodeSet(obj.Ptr.nodesetval)
}

func (nodeset *NodeSet) Size() int {
	return int(nodeset.Ptr.nodeNr)
}

func (nodeset *NodeSet) Item(index int) *Node {
	return makeNode(C.fetchNode(nodeset.Ptr, C.int(index)))
}
