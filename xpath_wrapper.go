package golibxml

func (obj *XPathObject) Type() XpathObjectType {
	return XpathObjectType(obj.Ptr._type)
}
