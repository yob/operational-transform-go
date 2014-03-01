package sharego

import (
	"testing"
)

func TestDictGetting(t *testing.T) {
	dict := Dict{
		"doc": "Haha this is is some text",
	}
	str, err := dict.get([]string{"doc"})
	if err != nil {
		t.Errorf("Cannot correctly get from dict")
	}
	if str != "Haha this is is some text" {
		t.Errorf("Cannot correctly get from dict")
	}
}

func TestDictSetting(t *testing.T) {
	dict := Dict{
		"doc": "Haha this is is some text",
	}
	dict.set([]string{"doc"}, "New string")
	str, err := dict.get([]string{"doc"})
	if err != nil {
		t.Errorf("Cannot properly set to dict (%s)", err)
	}
	if str != "New string" {
		t.Errorf("Cannot properly set to dict")
	}
}
