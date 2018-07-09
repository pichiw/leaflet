package leaflet

import (
	"github.com/gowasm/gopherwasm/js"

	"github.com/gowasm/vecty"
)

func NewTileLayer(o TileLayerOptions) *TileLayer {
	return &TileLayer{
		opts: o,
	}
}

func (l *TileLayer) JSValue() js.Value {
	return gL.Call(
		"tileLayer",
		"https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
		vecty.Value(l.opts),
	)
}

// AddTo add the receiver to the specified Map.
func (l *TileLayer) AddTo(m *Map) {
	l.JSValue().Call("addTo", vecty.Value(m))
}

type TileLayer struct {
	// Layer
	opts TileLayerOptions
}

// TileLayerOptions are tile layer options
type TileLayerOptions struct {
	ID          string `js:"id"`
	Attribution string `js:"attribution"`
	MaxZoom     int    `js:"maxZoom"`
}
