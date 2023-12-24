package pdfgen

import (
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/customerr"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/elem"
)

type ElementObject struct {
	Tp   elem.ElType
	Data interface{}
}
type DataObject struct {
	MapObject map[string]ElementObject
}

func NewDataObject() DataObject {
	m := make(map[string]ElementObject)
	return DataObject{MapObject: m}
}

func (do DataObject) Save(id string, tp elem.ElType, obj interface{}) {
	if _, f := do.MapObject[id]; f {
		return
	}
	do.MapObject[id] = ElementObject{
		Tp:   tp,
		Data: obj,
	}
}

func (do DataObject) Load(id string) (elem.ElType, interface{}, error) {
	if v, f := do.MapObject[id]; !f {
		return "", nil, customerr.ErrInvalidObjectID
	} else {
		return v.Tp, v.Data, nil
	}
}

//func LoadSavedObject[T any](do DataObject, id string) (*T, bool, error) {
//	if obj, f := do.MapObject[id]; !f {
//		return nil, false, nil
//	} else {
//		if data, ok := obj.Data.(T); !ok {
//			return nil, true, customerr.ErrMissMatchObjectAndType
//		} else {
//			return &data, true, nil
//		}
//	}
//}
