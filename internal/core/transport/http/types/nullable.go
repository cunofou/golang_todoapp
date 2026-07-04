package types

type Nullable[T any] struct {
    Value *T
    Set   bool
}

func (n Nullable[T]) ToDomain() Nullable[T] {
    return n
}