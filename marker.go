package leaflet

import (
	"sync"

	"github.com/gowasm/gopherwasm/js"

	"github.com/gowasm/vecty"
)

func NewMarker(c *Coordinate, events Events) *Marker {
	return &Marker{
		coord:  c,
		events: events,
	}
}

func (m *Marker) JSValue() js.Value {
	m.valueOnce.Do(func() {
		marker := gL.Call("marker", vecty.Value(m.coord))
		m.events.Bind(marker)
		m.v = &marker
	})
	return *m.v
}

// AddTo add the receiver to the specified Map.
func (l *Marker) AddTo(m *Map) {
	l.JSValue().Call("addTo", vecty.Value(m))
}

func (l *Marker) Remove() {
	l.JSValue().Call("remove")
}

type Marker struct {
	v         *js.Value
	valueOnce sync.Once
	coord     *Coordinate
	events    Events
}
