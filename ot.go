package sharego

import (
	"crypto/sha1"
	"fmt"
	"io"
	"unsafe"
)

//hashes a Dict, to produce checksums used within Document struct. hashes reflects
//the whole dict, both values and keys to be unique for each document.
func hash(content Dict) string {
	h := sha1.New()
	for key, val := range content {
		io.WriteString(h, key)
		switch value := val.(type) {
		case Dict:
			io.WriteString(h, hash(value))
		case []interface{}:
			for _, el := range value {
				switch element := el.(type) {
				case Dict:
					io.WriteString(h, hash(element))
				case fmt.Stringer:
					io.WriteString(h, element.String())
				}

			}
		case fmt.Stringer:
			io.WriteString(h, value.String())
		case unsafe.Pointer:
			io.WriteString(h, *(*string)(value))
		}
	}
	return string(h.Sum(nil))
}

//Given the old position of an insert operation returns its new position
//when transforming against another component.
func transformPosition(oldpos int, comp Component) (newpos int) {
	newpos = oldpos
	compos := comp.position()
	if comp.Si != "" {
		if compos <= oldpos {
			newpos += len(comp.Si)
		}
	} else {
		if oldpos <= compos {
			newpos = oldpos
		} else if oldpos <= compos+len(comp.Sd) {
			newpos = compos
		} else {
			newpos = oldpos - len(comp.Sd)
		}
	}
	return

}

