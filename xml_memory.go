package golibxml

/*
#cgo pkg-config: libxml-2.0
#include <libxml/parser.h>
*/
import "C"

////////////////////////////////////////////////////////////////////////////////
// INTERFACE
////////////////////////////////////////////////////////////////////////////////

// xmlCleanupMemory
func CleanupMemory() {
	C.xmlCleanupMemory()
}

// xmlInitMemory
func InitMemory() {
	C.xmlInitMemory()
}

// xmlMemBlocks
func MemBlocks() int {
	return int(C.xmlMemBlocks())
}

// xmlMemUsed
func MemUsed() int {
	return int(C.xmlMemUsed())
}

// xmlMemoryDump
func MemoryDump() {
	C.xmlMemoryDump()
}
