package leaflet

import (
	"github.com/gowasm/gopherwasm/js"

	"github.com/gowasm/vecty"
)

var gL = js.Global().Get("L")

// NewMap creates a new leaflet map
func NewMap(
	divid string,
	opts MapOptions,
	events Events,
) *Map {

	m := &Map{
		Value:  gL.Call("map", divid, vecty.Value(opts)).Call("setView", vecty.Value(opts.Center), opts.Zoom),
		divid:  divid,
		events: events,
	}

	m.coreEvents().Bind(m.Value)
	m.events.Bind(m.Value)

	return m
}

// Map represents a leaflet map
type Map struct {
	js.Value
	vecty.Core

	divid  string
	events Events
}

func (m *Map) Bounds(b js.Value) {
	m.Value.Call("fitBounds", b)
}

func (m *Map) onZoom(vs []js.Value) {
}

func (m *Map) Zoom() int {
	zoom := m.Get("zoom")
	if zoom == js.Undefined() {
		return 0
	}

	return zoom.Int()
}

func (m *Map) View(v *Coordinate, zoom int) {
	m.Call("setView", vecty.Value(v), zoom)
}

func (m *Map) coreEvents() Events {
	return Events{
		"zoom": m.onZoom,
	}
}

type MapOptions struct {
	Center  *Coordinate `js:"center"`
	Zoom    int         `js:"zoom"`
	MinZoom int         `js:"minZoom"`
	MaxZoom int         `js:"maxZoom"`
}
