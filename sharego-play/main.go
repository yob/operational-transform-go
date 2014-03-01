package main

// An example app that uses sharego. This is just for me to learn how things
// work, so you can safely ignore it.
//
// USAGE:
//
//    go get github.com/yob/sharego
//    go install github.com/yob/sharego/sharego-play
//    ./bin/sharego-play
//

import (
	"github.com/yob/sharego"
	"log"
)

// it's alive!
func main() {
	dict := sharego.Dict{
			"doc": "Haha this is is some text",
	}
	doc := sharego.NewDocument(dict)
	subdoc, _ := doc.Get([]string{"doc"})
	log.Println(subdoc)
}

