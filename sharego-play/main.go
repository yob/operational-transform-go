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
	op := sharego.Operation{}
	comp1 := sharego.NewInsertComponent([]string{"doc","0"}, "aaa ")
	comp2 := sharego.NewDeleteComponent([]string{"doc","14"}, "is ")
	op = append(op, comp1)
	op = append(op, comp2)
	doc.Apply(op)
	subdoc, err := doc.Get([]string{"doc"})
	if (err != nil) {
		log.Fatal("Error getting doc")
	}
	log.Println(subdoc)
}

