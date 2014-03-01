package sharego

import (
	"strconv"
)

//Components contains a Path ["doc", "toto", "0"], that is the list of keys
//to descend to underlying strings and apply inserts and deletes contained
//within either si or Sd.
type Component struct {
	Path []string
	Si   string
	Sd   string
}

//Transforms a component against another one. We use dest to accumulate
//components because the transform of a component may result in several
//components.
func (comp1 Component) transform(dest *Operation, comp2 Component) {
	pos1 := comp1.position()
	if comp1.Si != "" { //Insert
		comp1.setPosition(transformPosition(pos1, comp2))
	} else { //Delete
		if comp2.Si != "" { // Delete vs Insert
			deleted := comp1.Sd
			if pos1 < comp2.position() {
				(*dest).Append(Component{
					Path: comp1.Path,
					Sd:   deleted[:comp2.position()-pos1]})
				deleted = deleted[comp2.position()-pos1:]
			}
			if deleted != "" {
				(*dest).Append(Component{
					Path: append(comp1.Path[:len(comp1.Path)-1], strconv.Itoa(pos1+len(comp2.Si))),
					Sd:   deleted,
				})

			}
		}
	}
	return
}

//Returns the position at which a component is operating.
func (comp Component) position() (pos int) {
	pos, _ = strconv.Atoi(comp.Path[len(comp.Path)-1])
	return
}

//Sets position of component.
func (comp Component) setPosition(newpos int) {
	comp.Path[len(comp.Path)-1] = strconv.Itoa(newpos)
	return
}

//HelperFunction to create a new component. To be used in all cases because
//struct members are all private.
func NewInsertComponent(Path []string, str string) (comp Component) {
	comp.Path = Path
	comp.Si = str
	return
}

//HelperFunction to create a new component. To be used in all cases because
//struct members are all private.
func NewDeleteComponent(path []string, str string) (comp Component) {
	comp.Path = path
	comp.Sd = str
	return
}


type InvalidComponentError struct {
	msg string
}

func (e InvalidComponentError) Error() string {
	return e.msg
}

