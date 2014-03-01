package sharego

//This is document type. Dict is a very flexible struct to contain python like
//dict structure. For now only supports string inserts and deletes.
type Document struct {
	content   Dict
	ops       []Operation
}

//Returns a new document containing initialized map, from dict.
func NewDocument(content Dict) (doc Document) {
	doc = Document{
		ops: make([]Operation, 0),
		content: content,
	}
	return
}

//Every operation applied to the document will increment the version number by 1
func (doc Document) Version() int {
	return len(doc.ops)
}

//Applies an operation to this document. Version argument indicates what
//doc version the operation was built against. It is useful when receiving
//remote ops to know how to tranform received op against local ops.
func (doc *Document) ApplyToVersion(op Operation, version int) (err error) {
	if version != len(doc.ops) {
		transform_ops := doc.ops[version:]
		for i := 0; i < len(transform_ops); i++ {
			top := transform_ops[i]
			op = op.transform(top)
		}
	}
	content := doc.content
	for c := 0; c < len(op); c++ {
		comp := op[c]
		if comp.Si != "" {
			index := comp.position()
			str, err := content.get(comp.Path[:len(comp.Path)-1])
			if err != nil {
				return InvalidComponentError{msg: str}
			}
			str = str[:index] + comp.Si + str[index:]
			content.set(comp.Path[:len(comp.Path)-1], str)
		}
		if comp.Sd != "" {
			str, err := content.get(comp.Path[:len(comp.Path)-1])
			if err != nil {
				return err
			}
			str_length := len(comp.Sd)
			index := comp.position()
			deleted := str[index : index+str_length]
			if deleted != comp.Sd {
				return InvalidComponentError{"Trying to delete '" + comp.Sd + "' but found '" + deleted + "' instead"}
			}
			new_str := str[:index] + str[index+str_length:]
			content.set(comp.Path[:len(comp.Path)-1], new_str)
		}
	}
	doc.ops = append(doc.ops, op)
	return nil
}

//In order to access portions of the document. path is the list of keys in
//descending order to access final string.
func (doc *Document) Get(path []string) (inner string, err error) {
	return doc.content.get(path)
}

//Applies an operation to a document. This one will apply only local ops to
//the last version of the document. It will automatically send the op to
//connected documents.
func (doc *Document) Apply(op Operation) (err error, finished chan bool) {
       version := doc.Version()
       err = doc.ApplyToVersion(op, version)
       // comment this out until we have network support again
       //if err == nil {
       //        finished = make(chan bool, 1)
       //        doc.sendRemote(op, version, finished)
       //}
       return
}
