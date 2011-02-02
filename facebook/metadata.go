package facebook

type Metadata struct {
	Connections map[string] string
	Fields []*Field
}

type Field struct {
	Name string
	Description string
}