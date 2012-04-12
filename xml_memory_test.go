package golibxml

import "testing"

func TestCleanupMemory(t *testing.T) {
	CleanupMemory()
}

func TestInitMemory(t *testing.T) {
	InitMemory()
}

func TestMemBlocks(t *testing.T) {
	if MemBlocks() != 0 {
		t.Fail()
	}
}

func TestMemUsed(t *testing.T) {
	if MemUsed() != 0 {
		t.Fail()
	}
}

func TestMemoryDump(t *testing.T) {
	MemoryDump()
}
