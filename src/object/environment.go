package object

// NewEnclosedEnvironment is [TODO]
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment is [TODO]
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment is [TODO]
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get is [TODO]
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set is [TODO]
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
