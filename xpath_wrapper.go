package golibxml

func (obj *XPathObject) Type() XpathObjectType {
	return XpathObjectType(obj.Ptr._type)
}

func (obj *XPathObject) NodeSet() *NodeSet {
	return makeNodeSet(obj.Ptr.nodesetval)
}
