package leaflet

// PathOptions are the options for the path type
type PathOptions struct {
	MaxZoom             int    `js:"maxZoom"`
	Stroke              bool   `js:"stroke"`
	Color               string `js:"color"`
	Weight              int    `js:"weight"`
	Opacity             int    `js:"opacity"`
	LineCap             string `js:"lineCap"`
	LineJoin            string `js:"lineJoin"`
	DashArray           string `js:"dashArray"`
	DashOffset          string `js:"dashOffset"`
	Fill                bool   `js:"fill"`
	FillColor           string `js:"fillColor"`
	FillOpacity         int    `js:"fillOpacity"`
	FillRule            string `js:"fillRule"`
	BubblingMouseEvents bool   `js:"bubblingMouseEvents"`
	ClassName           string `js:"className"`
}
