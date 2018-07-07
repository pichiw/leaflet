package leaflet

import (
	"sync"

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
		Value:  gL.Call("map", divid, vecty.Value(opts)),
		divid:  divid,
		opts:   opts,
		events: events,
	}

	m.coreEvents().Bind(m.Value)
	m.events.Bind(m.Value)

	return m
}

// MapAdderTo defines things that can be added to a map
type MapAdderTo interface {
	AddTo(m *Map)
}

// Map represents a leaflet map
type Map struct {
	js.Value

	vecty.Core

	divid string

	opts      MapOptions
	optsMutex sync.RWMutex

	events Events
}

func (m *Map) Remove() {

}

// Add adders to map
func (m *Map) Add(as ...MapAdderTo) {
	for _, a := range as {
		a.AddTo(m)
	}
}

func (m *Map) onZoom(vs []js.Value) {
	if len(vs) == 0 {
		return
	}

	event := vs[0]
	target := event.Get("target")
	if target == js.Undefined() {
		return
	}
	zoom := target.Get("_zoom")
	if zoom == js.Undefined() {
		return
	}

	m.optsMutex.Lock()
	defer m.optsMutex.Unlock()
	m.opts.Zoom = zoom.Int()
}

func (m *Map) Zoom() int {
	m.optsMutex.RLock()
	defer m.optsMutex.RUnlock()
	return m.opts.Zoom
}

func (m *Map) Center() *Coordinate {
	m.optsMutex.RLock()
	defer m.optsMutex.RUnlock()
	return m.opts.Center
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
