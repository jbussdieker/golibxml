package main

import "os"

import . "golibxml"

func print_element_names(a_node *Node) {
	for cur_node := a_node; cur_node != nil; cur_node = cur_node.Next() {
		if (cur_node.Type() == XML_ELEMENT_NODE) {
			println("node type: Element, name:", cur_node.Name())
		}

		print_element_names(cur_node.Children())
	}
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	/*parse the file and get the DOM */
	doc := ReadFile(os.Args[1], "", 0)

	if doc == nil {
		println("error: could not parse file", os.Args[1])
		os.Exit(1)
	}

	/*Get the root element node */
	root_element := doc.Root()

	print_element_names(root_element)

	/*free the document */
	doc.Free()

	/*
	 *Free the global variables that may
	 *have been allocated by the parser.
	 */
	CleanupParser()
}
