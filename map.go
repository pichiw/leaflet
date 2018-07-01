package leaflet

import (
	"sync"
	"syscall/js"

	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

var gL = js.Global().Get("L")

// NewMap creates a new leaflet map
func NewMap(divid string, adders ...MapAdderTo) *Map {
	return &Map{
		divid:  divid,
		adders: adders,
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

// Mount is called after everything renders and the dom is fully mounted
func (m *Map) Mount() {
	m.Value = gL.Call("map", m.divid)

	m.addersMutex.Lock()
	defer m.addersMutex.Unlock()
	for _, a := range m.adders {
		a.AddTo(m)
	}

	m.Call("setView", NewCoordinate(49.8951, -97.1384).Value, 5)

}

// Render renders the map
func (m *Map) Render() vecty.ComponentOrHTML {
	return elem.Div(vecty.Markup(vecty.Attribute("id", m.divid)))
}