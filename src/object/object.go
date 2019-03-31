package object

import (
	"ast"
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"
)

// ObjectType is [TODO]
type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBj         = "HASH"
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

// Error is [TODO]
type Error struct {
	Message string
}

// Type is [TODO]
func (e *Error) Type() ObjectType { return ERROR_OBJ }

// Inspect is [TODO]
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

// Function is [TODO]
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type is [TODO]
func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

// Inspect is [TODO]
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// String is [TODO]
type String struct {
	Value string
}

// Type is [TODO]
func (s *String) Type() ObjectType { return STRING_OBJ }

// Inspect is [TODO]
func (s *String) Inspect() string { return s.Value }

// BuiltinFunction is [TODO]
type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

// Type is [TODO]
func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }

// Inspect is [TODO]
func (b *Builtin) Inspect() string { return "builtin function" }

// Array is [TODO]
type Array struct {
	Elements []Object
}

// Type is [TODO]
func (ao *Array) Type() ObjectType { return ARRAY_OBJ }

// Inspect is [TODO]
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// HashKey is [TODO]
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashKey is [TODO]
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

// HashKey is [TODO]
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// HashKey is [TODO]
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// HashPair is [TODO]
type HashPair struct {
	Key   Object
	Value Object
}

// Hash is [TODO]
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type is [TODO]
func (h *Hash) Type() ObjectType { return HASH_OBj }

// Inspect is [TODO]
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// Hashable is [TODO]
type Hashable interface {
	HashKey() HashKey
}
