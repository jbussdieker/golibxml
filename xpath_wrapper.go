package golibxml

/*
#cgo pkg-config: libxml-2.0
#include <libxml/xpath.h>

xmlNode* fetchNode(xmlNodeSet *nodeset, int index) {
  	return nodeset->nodeTab[index];
}
*/
import "C"

func (obj *XPathObject) Results() chan *Node {
	channel := make(chan *Node)
	go func(obj *XPathObject, channel chan *Node) {
		if obj.Ptr._type != 1 {
			close(channel)
			return
		}
		for i := 0; i < int(obj.Ptr.nodesetval.nodeNr); i++ {
			channel <- makeNode(C.fetchNode(obj.Ptr.nodesetval, C.int(i)))
		}
		close(channel)
	}(obj, channel)
	return channel
}

