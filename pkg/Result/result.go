package Result

// Main interface for the external packages
type Type[T any] interface {
	isResult()
	Value() *T
	IsErr() bool
	Error() error
}

// The type that indicates something was created
type some[T any] struct {
	value *T
}

// This creates a useful result to be used
func Value[T any](v *T) Type[T] {
	return &some[T]{
		value: v,
	}
}

func (_ *some[T]) isResult(){}

func (r *some[T]) Value() *T{
	return r.value
}

func (_ *some[T]) IsErr() bool{
	return false
}

func (_ *some[T]) Error() error {
	return nil
}

// This type represents error
type errors[T any] struct {
	err error
}

// Used to create a Error type of a particular type
func Err[T any](err error) Type[T] {
	return &errors[T]{
		err: err,
	}
}

func (_ *errors[T]) isResult(){}

func (_ *errors[T]) Value() *T{
	return nil
}

func (_ *errors[T]) IsErr() bool{
	return true
}

func (n *errors[T]) Error() error{
	return n.err
}
