package golibxml

import "testing"
import "io/ioutil"

func TestGetFeaturesList(t *testing.T) {
	features := GetFeaturesList()
	if len(features) != 42 {
		t.Fail()
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
	if doc.String() != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<catalog/>\n" {
		t.Fail()
	}
}

func TestParseEntity(t *testing.T) {
	doc := ParseEntity("test/test.xml")
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	if doc.String() != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<catalog/>\n" {
		t.Fail()
	}
}

func TestParseFile(t *testing.T) {
	doc := ParseFile("test/test.xml")
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	if doc.String() != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<catalog/>\n" {
		t.Fail()
	}
}

func TestParseMemory(t *testing.T) {
	buf, err := ioutil.ReadFile("test/test.xml")
	if err != nil {
		t.Fatal(err)
	}
	doc := ParseMemory(buf)
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	if doc.String() != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<catalog/>\n" {
		t.Fail()
	}
}

func TestReadDoc(t *testing.T) {
	doc := ReadDoc("<catalog/>", "", "", 0)
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	if doc.String() != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<catalog/>\n" {
		t.Fail()
	}
}

func TestReadFile(t *testing.T) {
	doc := ReadFile("test/test.xml", "UTF-8", 0)
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	if doc.String() != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<catalog/>\n" {
		t.Fail()
	}
}

func TestReadMemory(t *testing.T) {
	buf, err := ioutil.ReadFile("test/test.xml")
	if err != nil {
		t.Fatal(err)
	}
	doc := ReadMemory(buf, "/", "UTF-8", 0)
	if doc == nil {
		t.Fail()
	}
	defer doc.Free()
	if doc.String() != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<catalog/>\n" {
		t.Fail()
	}
}

func TestRecoverDoc(t *testing.T) {
	buf, err := ioutil.ReadFile("test/test.xml")
	if err != nil {
		t.Fatal(err)
	}
	doc := RecoverDoc(string(buf))
	if doc.String() != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<catalog/>\n" {
		t.Fail()
	}
}

func TestRecoverFile(t *testing.T) {
	doc := RecoverFile("test/test.xml")
	if doc == nil {
		t.Fail()
	}
	if doc.String() != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<catalog/>\n" {
		t.Fail()
	}
}

func TestRecoverMemory(t *testing.T) {
	buf, err := ioutil.ReadFile("test/test.xml")
	if err != nil {
		t.Fatal(err)
	}
	doc := RecoverMemory(buf)
	if doc.String() != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<catalog/>\n" {
		t.Fail()
	}
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
