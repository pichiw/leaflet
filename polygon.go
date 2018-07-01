package leaflet

import (
	"syscall/js"

	"github.com/gowasm/vecty"
)

func NewPolyline(opts PolylineOptions, coords ...*Coordinate) *Polyline {
	o := js.Global().Get("Array").New()
	for _, c := range coords {
		o.Call("push", c.Value)
	}

	return &Polyline{
		Layer: Layer{
			Value: gL.Call("polyline", o, vecty.Value(opts)),
		},
		coordinates: coords,
	}
}

type Polyline struct {
	Layer
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
