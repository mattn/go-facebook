package facebook

type Object struct {
	ID   string
	Name string
}

type ObjectCollection struct {
	Data []*Object
}

func (o *ObjectCollection) Objects() []*Object {
	if o.Data == nil {
		o.Data = make([]*Object, 0)
	}
	return o.Data
}

func (o *ObjectCollection) Object(index int) *Object {
	if o.Data == nil {
		return nil
	}
	if len(o.Data) > index {
		return o.Data[index]
	}
	return nil
}
