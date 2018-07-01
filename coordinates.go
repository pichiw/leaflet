package leaflet

import "syscall/js"

func NewCoordinate(lat, lng float64) *Coordinate {
	return &Coordinate{
		lat:   lat,
		lng:   lng,
		Value: gL.Call("latLng", lat, lng),
	}
}

type Coordinate struct {
	js.Value

	lat float64
	lng float64
}

func (c *Coordinate) Lat() float64 {
	return c.lat
}

func (c *Coordinate) Lng() float64 {
	return c.lng
}
