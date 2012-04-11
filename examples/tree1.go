package main

import "os"

import . "github.com/jbussdieker/golibxml"

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	/*parse the file and get the DOM */
	doc := ReadFile(os.Args[1], "", 0);

	if doc == nil {
		println("error: could not parse file", os.Args[1])
		os.Exit(1)
	}

	/*Get the root element node */
	root_element := doc.GetRoot()

	/*free the document */
	doc.Free()
}
