package component

type Basic interface {
	ID() uint64
}

type Controllable interface {
	Basic
}
