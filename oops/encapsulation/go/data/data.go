package data

type ExportedFields struct {
	Field1 string
	Field2 string
}

type unexportedFields struct {
	field1 string
	field2 string
}

func NewUnexportedFields(val1, val2 string) *unexportedFields {
	return &unexportedFields{
		field1: val1,
		field2: val2,
	}
}
