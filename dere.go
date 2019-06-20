// Package dere implements deep compound and aggregate data type zero values using reflection.
package dere

import (
	"reflect"
)

// DeepZero returns a deep zero value for val.
func DeepZero(val interface{}) interface{} {
	if val == nil {
		return nil
	}
	t := reflect.TypeOf(val)
	v := reflect.New(t)
	vv := v.Elem()

	ts := map[reflect.Type]bool{}
	setValue(ts, t, vv)
	return vv.Interface()
}

func setStruct(ts map[reflect.Type]bool, t reflect.Type, v reflect.Value) {
	ts[t] = true
	n := v.NumField()
	for i := 0; i < n; i++ {
		f := v.Field(i)
		ft := t.Field(i)
		setValue(ts, ft.Type, f)
	}
	delete(ts, t)
}

func setValue(ts map[reflect.Type]bool, t reflect.Type, v reflect.Value) {
	switch t.Kind() {
	case reflect.Map:
		m := reflect.MakeMap(t)
		v.Set(m)
	case reflect.Ptr:
		t = t.Elem()
		if ts[t] {
			return
		}
		vv := reflect.New(t)
		setValue(ts, t, vv.Elem())
		v.Set(vv)
	case reflect.Struct:
		if ts[t] {
			return
		}
		setStruct(ts, t, v)
	default:
	}
}
