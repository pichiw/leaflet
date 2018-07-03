package leaflet

import (
	"syscall/js"

	"github.com/gowasm/vecty"
)

type Layer struct {
	js.Value
}

// AddTo add the receiver to the specified Map.
func (l *Layer) AddTo(m *Map) {
	l.Value.Call("addTo", vecty.Value(m))
}
