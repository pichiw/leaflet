package leaflet

import (
	"syscall/js"

	"github.com/gowasm/vecty"
)

func NewPolyline(opts PolylineOptions, coords ...*Coordinate) *Polyline {
	return &Polyline{
		opts:        opts,
		coordinates: coords,
	}
}

func (l *Polyline) JSValue() js.Value {
	if l.v != nil {
		return *l.v
	}

	o := js.Global().Get("Array").New()
	for _, c := range l.coordinates {
		o.Call("push", vecty.Value(c))
	}

	v := gL.Call("polyline", o, vecty.Value(l.opts))
	l.v = &v
	return v
}

// AddTo add the receiver to the specified Map.
func (l *Polyline) AddTo(m *Map) {
	l.JSValue().Call("addTo", vecty.Value(m))
}

func (l *Polyline) Remove() {
	l.JSValue().Call("remove")
}

type Polyline struct {
	v           *js.Value
	opts        PolylineOptions
	coordinates []*Coordinate
}

func (p *Polyline) Coordinates() []*Coordinate {
	return p.coordinates
}

// PolylineOptions are options that can be applied to a Polyline
type PolylineOptions struct {
	PathOptions

	SmoothFactor int  `js:"smoothFactor"`
	NoClip       bool `js:"noClip"`
}
