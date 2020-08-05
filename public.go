package meander

// Facade interface
type Facade interface {
	Public() interface{}
}

// Public returns the public version of an interface
func Public(o interface{}) interface{} {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}
	return o
}
