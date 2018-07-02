package leaflet

import "github.com/gowasm/gopherwasm/js"

// Events define events that can be added to an object
type Events map[string]func(vs []js.Value)

func (e Events) Bind(o js.Value) {
	for n, ev := range e {
		o.Call("on", n, js.NewCallback(ev))
	}
}
