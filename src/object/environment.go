package object

// NewEnvironment is [TODO]
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment is [TODO]
type Environment struct {
	store map[string]Object
}

// Get is [TODO]
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set is [TODO]
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
