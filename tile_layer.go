package leaflet

func NewTileLayer(o TileLayerOptions) *TileLayer {
	return &TileLayer{
		Layer: Layer{
			Value: gL.Call(
				"tileLayer",
				"https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
				Value(o),
			),
		},
	}
}

type TileLayer struct{ Layer }

// TileLayerOptions are tile layer options
type TileLayerOptions struct {
	ID          string `js:"id"`
	Attribution string `js:"attribution"`
	MaxZoom     int    `js:"maxZoom"`
}
