package sharego


//Compares two strings to see if they are the same Path.
func pathEquals(strslice1, strslice2 []string) (b bool) {
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

