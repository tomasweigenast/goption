package goption

type Option[T any] struct {
	value *T
}

// Some creates an Option that contains a value of type T
func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

// None creates an Option that can hold a value of type T
func None[T any]() Option[T] {
	return Option[T]{}
}

// IsSome returns [true] if the option contains a value.
func (o *Option[T]) IsSome() bool {
	return o.value != nil
}

// IsNone returns [true] if the option does not contain a value.
func (o *Option[T]) IsNone() bool {
	return o.value == nil
}

// MustUnwrap returns the contained value in the option. This will panic if it does not contain a value.
func (o *Option[T]) MustUnwrap() T {
	if o.value == nil {
		panic("Unwrap() called on a None Option")
	}

	return *o.value
}

// Unwrap returns the underlaying pointer value in the option. It will return [nil] if the Option does not contain
// a value.
func (o *Option[T]) Unwrap() *T {
	return o.value
}

// UnwrapOr returns the contained value in the option or [defaultValue] is the option does not contain a value.
func (o *Option[T]) UnwrapOr(defaultValue T) T {
	if o.value == nil {
		return defaultValue
	}

	return *o.value
}

// UnwrapOrElse returns the contained value in the option or computes [f], returning its value
func (o *Option[T]) UnwrapOrElse(f func() T) T {
	if o.value == nil {
		return f()
	}

	return *o.value
}

// Insert inserts the given [value] into the Option. It the option already contains a value, its dropped
func (o *Option[T]) Insert(value T) {
	o.value = &value
}

// Take takes the value out of the option, leaving a None in its place
func (o *Option[T]) Take() Option[T] {
	new := Option[T]{value: o.value}
	o.value = nil
	return new
}

// Mutate transforms the value contained in the option, if any, appliying [f]. Otherwise it is a no-op.
func (o *Option[T]) Mutate(f func(*T)) {
	if o.value != nil {
		f(o.value)
	}
}

// Copy copies the entire
func (o *Option[T]) Copy() Option[T] {
	var newVal *T
	if o.value != nil {
		val := *o.value
		newVal = &val
	}

	return Option[T]{newVal}
}
