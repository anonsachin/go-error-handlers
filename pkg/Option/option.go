package Option

// Main interface for the external packages
type Type[T any] interface {
	isOption()
	Value() *T
	IsNone() bool
}

// The type that indicates something was created
type some[T any] struct {
	value *T
}

// This creates a useful result to be used
func Some[T any](v *T) Type[T] {
	return &some[T]{
		value: v,
	}
}

func (_ *some[T]) isOption(){}

func (r *some[T]) Value() *T{
	return r.value
}

func (_ *some[T]) IsNone() bool{
	return false
}

// This type represents empty ness
type none[T any] struct {
}

// Used to create a none type of a particular type
func None[T any]() Type[T] {
	return &none[T]{}
}

func (_ *none[T]) isOption(){}

func (_ *none[T]) Value() *T{
	return nil
}

func (_ *none[T]) IsNone() bool{
	return true
}
