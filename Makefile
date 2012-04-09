include $(GOROOT)/src/Make.inc

TARG=golibxml/xmltree

CGOFILES=\
  xmltree.go\
  xmlparser.go\

include $(GOROOT)/src/Make.pkg

