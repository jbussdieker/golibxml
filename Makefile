all:
	cd xmltree && go build -x .
	cd xmlparser && go build -x .
	cd htmltree && go build -x .
	cd htmlparser && go build -x .

install:
	cd xmltree && go install -x .
	cd xmlparser && go install -x .
	cd htmltree && go install -x .
	cd htmlparser && go install -x .

test:
	cd xmltree && go test -v .
	cd xmlparser && go test -v .
	cd htmltree && go test -v .
	cd htmlparser && go test -v .

fmt_all:
	cd xmltree && go fmt .
	cd xmlparser && go fmt .
	cd htmltree && go fmt .
	cd htmlparser && go fmt .

