package golibxml

import "testing"

const TEST_XML_STRING = `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:foo="http://bar">
	<channel>
		<item foo="bar">
			<title>test</title>
			<foo:xx />
		</item>
	</channel>
</rss>
`

func TestXmlReader(t *testing.T) {

		r := ReaderForMemory(TEST_XML_STRING, "", "", ParserOption(0))

		if r == nil {
			t.Fail()
		}

		defer r.Free()

		if r.Read() != 1 {
			t.Fail()
		}

		if r.NodeType() != XML_READER_TYPE_ELEMENT {
			t.Fail()
		}

		if r.Name() != "rss" {
			t.Fail()
		}
}

