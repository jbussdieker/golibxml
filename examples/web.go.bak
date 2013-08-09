package main

import "fmt"
import "flag"
import "html"

import "github.com/hoisie/web"
import . "github.com/jbussdieker/golibxml"

func testapp(webctx *web.Context, val string) (render string) {
	render += "<form><textarea style='width:100%; height:288;' name=document>"
	render += webctx.Params["document"]
	render += "</textarea><br><input name=xpath value=\"" + webctx.Params["xpath"] + "\"></input><input type=submit></form>"

	// Parse the document
	doc := ParseHTMLDoc(webctx.Params["document"], "UTF-8")
	if doc == nil {
		return
	}
	defer doc.Free()

	// Create an xpath context
	ctx := NewXPathContext(doc.Document)
	if ctx == nil {
		return
	}
	defer ctx.Free()

	// Evaluate the xpath search
	result := ctx.Eval(webctx.Params["xpath"])
	if result == nil {
		return
	}
	defer result.Free()

	// Display results
	for node := range result.Results() {
		render += fmt.Sprintln("Name:", node.Name(), "<br>")
		render += fmt.Sprintln("Type:", node.Type(), "<br>")
		render += fmt.Sprintln("Path:", node.Path(), "<br>")
		render += fmt.Sprintln("<br><code>", html.EscapeString(node.String()), "</code><br><br>")
	}

	return
}

func main() {
	var port int
	flag.IntVar(&port, "port", 9999, "Port to run server on")
	flag.Parse()
    web.Get("/(.*)", testapp)
    web.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
