package main

import "golibxml"

func example1Func(filename string) {
    doc := ReadFile(filename, "", 0);

    if (doc == NULL) {
        println("Failed to parse", filename);
		return;
    }
    doc.Free()
}

func main() {
    if len(os.Args) != 2 {
        os.Exit(1)
	}

    example1Func(os.Args(1))

    /*
     * Cleanup function for the XML library.
     */
    CleanupParser()
    /*
     * this is to debug memory for regression tests
     */
    MemoryDump()
}
