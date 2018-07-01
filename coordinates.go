package leaflet

import "syscall/js"

// NewCoordinate creates a new coordinate
func NewCoordinate(lat, lng float64) *Coordinate {
	return &Coordinate{
		lat:   lat,
		lng:   lng,
		Value: gL.Call("latLng", lat, lng),
	}
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
