package sharego

//This is document type. Dict is a very flexible struct to contain python like
//dict structure. For now only supports string inserts and deletes. Checksums
//is a mapping between the checksum of a document and the index within the ops
//array. It is used when receiving remote ops that where built against old versions
//of the doc to transform received operation against the operation that occured
//locally in the meantime.
type Document struct {
	content   Dict
	checksums map[string]int
	ops       []Operation
}

//Returns a new document containing initialized map, from dict.
func NewDocument(content Dict) (doc Document) {
	h := hash(content)
	doc = Document{
		content: content,
		checksums: map[string]int{
			h: 0,
		},
	}
	return
}

//Returns the hash of the document. Used to determine version of doc (but
//timeline agnostic as it only depends on the content.
func (doc Document) Checksum() string {
	return hash(doc.content)
}

//Applies an operation.checksum argument represents what
//checksum the document was built against. It is useful when receiving
//remote ops to know how to tranform received op against local ops.
//Apply func is in network.go
func (doc *Document) applyNoRemote(op Operation, checksum string) (err error) {
	last_op_index := doc.checksums[checksum]
	if last_op_index != len(doc.ops) {
		transform_ops := doc.ops[last_op_index:]
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
	doc.checksums[doc.Checksum()] = len(doc.ops)
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
       checksum := doc.Checksum()
       err = doc.applyNoRemote(op, checksum)
       // comment this out until we have network support again
       //if err == nil {
       //        finished = make(chan bool, 1)
       //        doc.sendRemote(op, checksum, finished)
       //}
       return
}
