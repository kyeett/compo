package component

import "sync/atomic"

var (
	id uint64
)

type BasicEntity struct {
	id uint64
}

func NewBasic() BasicEntity {
	return BasicEntity{id: atomic.AddUint64(&id, 1)}
}
