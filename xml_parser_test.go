package golibxml

import "testing"
import "io/ioutil"

const TEST_XML_FILE = "test/test.xml"
const TEST_XML_EXPECTED = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<catalog/>\n"

func compareResult(t *testing.T, got string, expected string) {
	if got != expected {
		t.Fatal("\nGot:\n" + got + "\nExpected:\n" + expected)
	}
}

func TestParseDtd(t *testing.T) {
	dtd := ParseDTD("1", "test/test.dtd")
	if dtd == nil {
		t.Fail()
	}
}

func TestParseDoc(t *testing.T) {
	doc := ParseDoc("<catalog/>")
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	compareResult(t, doc.String(), TEST_XML_EXPECTED)
}

func TestParseEntity(t *testing.T) {
	doc := ParseEntity(TEST_XML_FILE)
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	compareResult(t, doc.String(), TEST_XML_EXPECTED)
}

func TestParseFile(t *testing.T) {
	doc := ParseFile(TEST_XML_FILE)
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	compareResult(t, doc.String(), TEST_XML_EXPECTED)
}

func TestParseMemory(t *testing.T) {
	buf, err := ioutil.ReadFile(TEST_XML_FILE)
	if err != nil {
		t.Fatal(err)
	}
	doc := ParseMemory(buf)
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	compareResult(t, doc.String(), TEST_XML_EXPECTED)
}

func TestReadDoc(t *testing.T) {
	doc := ReadDoc("<catalog/>", "", "", 0)
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	compareResult(t, doc.String(), TEST_XML_EXPECTED)
}

func TestReadFile(t *testing.T) {
	doc := ReadFile(TEST_XML_FILE, "UTF-8", 0)
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	compareResult(t, doc.String(), TEST_XML_EXPECTED)
}

func TestReadMemory(t *testing.T) {
	buf, err := ioutil.ReadFile(TEST_XML_FILE)
	if err != nil {
		t.Fatal(err)
	}
	doc := ReadMemory(buf, "/", "UTF-8", 0)
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	compareResult(t, doc.String(), TEST_XML_EXPECTED)
}

func TestRecoverDoc(t *testing.T) {
	buf, err := ioutil.ReadFile(TEST_XML_FILE)
	if err != nil {
		t.Fatal(err)
	}
	doc := RecoverDoc(string(buf))
	compareResult(t, doc.String(), TEST_XML_EXPECTED)
}

func TestRecoverFile(t *testing.T) {
	doc := RecoverFile(TEST_XML_FILE)
	if doc == nil {
		t.Fail()
	}
	compareResult(t, doc.String(), TEST_XML_EXPECTED)
}

func TestRecoverMemory(t *testing.T) {
	buf, err := ioutil.ReadFile(TEST_XML_FILE)
	if err != nil {
		t.Fatal(err)
	}
	doc := RecoverMemory(buf)
	compareResult(t, doc.String(), TEST_XML_EXPECTED)
}

func TestSubstituteEntitiesDefault(t *testing.T) {
	start := SubstituteEntitiesDefault(0)
	cur := SubstituteEntitiesDefault(1)
	if cur != 0 {
		t.Fail()
	}
	cur = SubstituteEntitiesDefault(0)
	if cur != 1 {
		t.Fail()
	}
	cur = SubstituteEntitiesDefault(start)
	if cur != 0 {
		t.Fail()
	}
}
