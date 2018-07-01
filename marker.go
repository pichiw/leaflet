package leaflet

func NewMarker(c *Coordinate) *Marker {
	return &Marker{
		Layer: Layer{
			Value: gL.Call("marker", c.Value),
		},
	}
}

type Marker struct{ Layer }
