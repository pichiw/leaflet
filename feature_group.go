package leaflet

import (
	"sync"

	"github.com/gowasm/gopherwasm/js"

	"github.com/gowasm/vecty"
)

func NewFeatureGroup(vs ...vecty.JSValuer) *FeatureGroup {
	return &FeatureGroup{
		vs: vs,
	}
}

type FeatureGroup struct {
	v         *js.Value
	valueOnce sync.Once
	vs        []vecty.JSValuer
}

func (f *FeatureGroup) JSValue() js.Value {
	f.valueOnce.Do(func() {
		a := js.Global().Get("Array").New()
		for _, v := range f.vs {
			a.Call("push", v.JSValue())
		}
		v := gL.Call("featureGroup", a)
		f.v = &v
	})
	return *f.v
}

func (f *FeatureGroup) Bounds() js.Value {
	return f.JSValue().Call("getBounds")
}
