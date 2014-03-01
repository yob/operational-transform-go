package sharego

//An operation is a list of components. To build a complex operation use
// op.Append(component).
type Operation []Component

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
