package golibxml

import "testing"
import "io/ioutil"

func TestReadXmlFile(t *testing.T) {
	out, err := ioutil.ReadFile("test/test.xml")
	if err != nil {
		t.Fatal("Failed to read test source")
	}
	doc := ParseDoc(string(out))
	if doc.String() != string(out) {
		t.Fail()
	}
}

func TestReadHtmlFile(t *testing.T) {
	out, err := ioutil.ReadFile("test/test.html")
	if err != nil {
		t.Fatal("Failed to read test source")
	}
	doc := ParseHTMLDoc(string(out), "UTF-8")
	if doc.String() != string(out) {
		t.Fail()
	}
}

