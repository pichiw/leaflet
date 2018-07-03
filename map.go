package leaflet

import (
	"sync"

	"github.com/gowasm/gopherwasm/js"

	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

var gL = js.Global().Get("L")

// NewMap creates a new leaflet map
func NewMap(
	divid string,
	opts MapOptions,
	events Events,
	adders ...MapAdderTo,
) *Map {
	return &Map{
		divid:  divid,
		adders: adders,
		opts:   opts,
		events: events,
	}
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

	adders      []MapAdderTo
	addersMutex sync.Mutex

	opts      MapOptions
	optsMutex sync.RWMutex

	events Events
}

// Add adders to map
func (m *Map) Add(as ...MapAdderTo) {
	if len(as) == 0 {
		return
	}

	m.addersMutex.Lock()
	defer m.addersMutex.Unlock()

	m.adders = append(m.adders, as...)
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

// Mount is called after everything renders and the dom is fully mounted
func (m *Map) Mount() {
	m.Value = gL.Call("map", m.divid, vecty.Value(m.opts))

	m.addersMutex.Lock()
	defer m.addersMutex.Unlock()
	for _, a := range m.adders {
		a.AddTo(m)
	}

	m.coreEvents().Bind(m.Value)
	m.events.Bind(m.Value)

	m.Value.Call("invalidateSize")
}

// Render renders the map
func (m *Map) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Attribute("id", m.divid)))
}

type MapOptions struct {
	Center  *Coordinate `js:"center"`
	Zoom    int         `js:"zoom"`
	MinZoom int         `js:"minZoom"`
	MaxZoom int         `js:"maxZoom"`
}
