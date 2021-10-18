package shine

import (
	"fmt"
	"reflect"
)

// CopySlice a slice to a slice
func CopySlice(dest interface{}, src interface{}) error {
	if src == nil {
		return fmt.Errorf("source can not be nil")
	}

	if dest == nil {
		return fmt.Errorf("dest can not be nil")
	}

	sv := reflect.ValueOf(src)
	st := sv.Type()

	if st.Kind() != reflect.Slice {
		return fmt.Errorf("expected src to be a slice, got %s", st.Kind())
	}

	dv := reflect.ValueOf(dest)
	dt := dv.Type()

	if dt.Kind() != reflect.Ptr {
		return fmt.Errorf("expected dest to be a ptr, got %s", dt.Kind())
	}

	dv = dv.Elem()
	if dv.Type() != sv.Type() {
		return fmt.Errorf("expected dest type to be %s, got %s", sv.Type(), dv.Type())
	}

	sl := reflect.MakeSlice(st, 0, sv.Len())
	for i := 0; i < sv.Len(); i++ {
		source := sv.Index(i)
		sl = reflect.Append(sl, source)
	}

	dv.Set(sl)

	return nil
}
