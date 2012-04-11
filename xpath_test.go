package golibxml

import "testing"

func TestXPathCompile(t *testing.T) {
	xpath := XPathCompile("*")
	if xpath == nil {
		t.Fail()
	}
}

func TestXPathCompileNeg(t *testing.T) {
	xpath := XPathCompile("")
	if xpath != nil {
		t.Fail()
	}
}

