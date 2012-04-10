package golibxml

import "testing"
import "io/ioutil"

const TEST_HTML_FILE = "test/test.html"
const TEST_HTML_EXPECTED = "<!DOCTYPE html PUBLIC \"-//W3C//DTD HTML 4.0 Transitional//EN\" \"http://www.w3.org/TR/REC-html40/loose.dtd\">\n<html></html>\n"

func TestParseHTMLDoc(t *testing.T) {
	doc := ParseHTMLDoc("<html></html>", "UTF-8")
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	compareResult(t, doc.String(), TEST_HTML_EXPECTED)
}

func TestParseHTMLFile(t *testing.T) {
	doc := ParseHTMLFile(TEST_HTML_FILE, "UTF-8")
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	compareResult(t, doc.String(), TEST_HTML_EXPECTED)
}

func TestReadHTMLDoc(t *testing.T) {
	doc := ReadHTMLDoc("<html></html>", "", "", 0)
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	compareResult(t, doc.String(), TEST_HTML_EXPECTED)
}

func TestReadHTMLFile(t *testing.T) {
	doc := ReadHTMLFile(TEST_HTML_FILE, "UTF-8", 0)
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	compareResult(t, doc.String(), TEST_HTML_EXPECTED)
}

func TestReadHTMLMemory(t *testing.T) {
	buf, err := ioutil.ReadFile(TEST_HTML_FILE)
	if err != nil {
		t.Fatal(err)
	}
	doc := ReadHTMLMemory(buf, "/", "UTF-8", 0)
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	compareResult(t, doc.String(), TEST_HTML_EXPECTED)
}

