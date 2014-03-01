package sharego

import (
	"crypto/sha1"
	"fmt"
	"io"
	"unsafe"
)

//An operation is a list of components. To build a complex operation use
// op.Append(component).
type Operation []Component

//Compares two strings to see if they are the same Path.
func PathEquals(strslice1, strslice2 []string) (b bool) {
	b = false
	if len(strslice1) != len(strslice2) {
		return
	}
	for i := 0; i < len(strslice1); i++ {
		el1 := strslice1[i]
		el2 := strslice2[i]
		if el1 != el2 {
			return
		}
	}
	b = true
	return
}

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


//Appends a new component to an operation. If many components already exists
//within op, it will try to compress them in as few components as possible.
func (op Operation) Append(comp Component) {
	op = append(op, comp)
}

//transforms an operation against another one. This basically transform every
//component against every other component
func (op1 Operation) transform(op2 Operation) Operation {
	for _, comp2 := range op2 {
		for _, comp1 := range op1 {
			comp2Path := comp2.Path[:len(comp2.Path)-1]
			comp1Path := comp1.Path[:len(comp1.Path)-1]
			if PathEquals(comp1Path, comp2Path) {
				comp1.transform(&op1, comp2)
			}
		}
	}
	return op1
}

