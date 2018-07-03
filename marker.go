package leaflet

func NewMarker(c *Coordinate, events Events) *Marker {
	m := gL.Call("marker", c.Value)
	events.Bind(m)
	return &Marker{
		Layer: Layer{
			Value: m,
		},
	}
}

type Marker struct{ Layer }
