package leaflet

import (
	"sync"

	"github.com/gowasm/gopherwasm/js"
)

// NewCoordinate creates a new coordinate
func NewCoordinate(lat, lng float64) *Coordinate {
	return &Coordinate{
		lat: lat,
		lng: lng,
	}
}

func (c *Coordinate) JSValue() js.Value {
	c.valueOnce.Do(func() {
		v := gL.Call("latLng", c.lat, c.lng)
		c.v = &v
	})
	return *c.v
}

// NewCoordinates creates a set of coordinates with alternating lats and longs
// If the number of values passed in isn't even it won't create the last coordinate
func NewCoordinates(latLngs ...float64) []*Coordinate {
	var coords []*Coordinate
	isLat := true
	var lat float64
	for _, ll := range latLngs {
		if isLat {
			lat = ll
			isLat = false
		} else {
			coords = append(coords, NewCoordinate(lat, ll))
			isLat = true
		}
	}

	return coords
}

type Coordinate struct {
	v         *js.Value
	valueOnce sync.Once
	lat       float64
	lng       float64
}

func (c *Coordinate) Lat() float64 {
	return c.lat
}

func (c *Coordinate) Lng() float64 {
	return c.lng
}
