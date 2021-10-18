package shine

import (
	"fmt"
	"reflect"
)

// FillSlice fills the destination slice with the src interfaces
func FillSlice(dest interface{}, src ...interface{}) error {

	if dest == nil {
		return fmt.Errorf("dest must not be nil")
	}

	if len(src) == 0 {
		return nil
	}

	dv := reflect.ValueOf(dest)
	if dv.Kind() != reflect.Ptr {
		return fmt.Errorf("dest must be a pointer to slice: %q", dv.Kind())
	}

	dv = dv.Elem()
	if dv.Kind() != reflect.Slice {
		return fmt.Errorf("dest must be a pointer to slice %v", dv.Kind())
	}

	dv.Set(reflect.MakeSlice(dv.Type(), len(src), len(src)))

	for i, w := range src {
		dv.Index(i).Set(reflect.ValueOf(w))
	}

	return nil
}
