package golibxml

import "testing"

func TestDetectCharEncoding(t *testing.T) {
	enc := DetectCharEncoding([]byte("testing"))
	if enc != 0 {
		t.Fatal(enc)
	}
}

func TestDetectCharEncodingUTF8(t *testing.T) {
	enc := DetectCharEncoding([]byte("\xEF\xBB\xBF"))
	if enc != XML_CHAR_ENCODING_UTF8 {
		t.Fatal(enc)
	}
}

func TestDetectCharEncodingUTF16BE(t *testing.T) {
	enc := DetectCharEncoding([]byte("\xFE\xFF"))
	if enc != XML_CHAR_ENCODING_UTF16BE {
		t.Fatal(enc)
	}
}

func TestDetectCharEncodingUTF16LE(t *testing.T) {
	enc := DetectCharEncoding([]byte("\xFF\xFE"))
	if enc != XML_CHAR_ENCODING_UTF16LE {
		t.Fatal(enc)
	}
}

