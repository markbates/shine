package shine

import (
	"fmt"
	"reflect"
)

// SliceElemType wiill return the type for the slice or array.
// ex. []string => reflect.TypeOf("string")
// ex. *[]Beatle => reflect.TypeOf(Beatle)
func SliceElemType(i interface{}) (reflect.Type, error) {
	rv := reflect.ValueOf(i)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	var sl reflect.Value
	switch rv.Kind() {
	case reflect.Array:
		sl = reflect.New(rv.Type())
		sl = reflect.Indirect(sl)
	case reflect.Slice:
		sl = reflect.MakeSlice(rv.Type(), 1, 1)
	default:
		return nil, fmt.Errorf("expected an array/slice, got %s", rv.Kind())
	}

	rv = sl.Index(0)
	rt := rv.Type()

	return rt, nil
}
