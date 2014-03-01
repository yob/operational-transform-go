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
	component := sharego.NewInsertComponent([]string{"doc","0"}, "aaa")
	op := sharego.Operation{}
	op = append(op, component)
	doc.Apply(op)
	subdoc, err := doc.Get([]string{"doc"})
	if (err != nil) {
		log.Fatal("Error getting doc")
	}
	log.Println(subdoc)
}

