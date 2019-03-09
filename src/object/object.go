package object

import "fmt"

// ObjectType is [TODO]
type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
)

// Object is [TODO]
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer is [TODO]
type Integer struct {
	Value int64
}

// Inspect is [TODO]
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type is [TODO]
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// Boolean is [TODO]
type Boolean struct {
	Value bool
}

// Type is [TODO]
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// Inspect is [TODO]
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Null is [TODO]
type Null struct{}

// Type is [TODO]
func (n *Null) Type() ObjectType { return NULL_OBJ }

// Inspect is [TODO]
func (n *Null) Inspect() string { return "null" }

// ReturnValue is [TODO]
type ReturnValue struct {
	Value Object
}

// Type is [TODO]
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

// Inspect is [TODO]
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }
