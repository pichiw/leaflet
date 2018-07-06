package leaflet

import (
	"syscall/js"

	"github.com/gowasm/vecty"
)

func NewMarker(c *Coordinate, events Events) *Marker {
	return &Marker{
		coord:  c,
		events: events,
	}
}

func (m *Marker) JSValue() js.Value {
	if m.v != nil {
		return *m.v
	}
	marker := gL.Call("marker", vecty.Value(m.coord))
	m.events.Bind(marker)
	m.v = &marker
	return marker
}

// AddTo add the receiver to the specified Map.
func (l *Marker) AddTo(m *Map) {
	l.JSValue().Call("addTo", vecty.Value(m))
}

func (l *Marker) Remove() {
	l.JSValue().Call("remove")
}

type Marker struct {
	// Layer
	v      *js.Value
	coord  *Coordinate
	events Events
}
