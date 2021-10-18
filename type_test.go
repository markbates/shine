package shine

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SliceElemType(t *testing.T) {
	t.Parallel()

	table := []struct {
		err  bool
		exp  reflect.Type
		name string
		sl   interface{}
	}{
		{name: "all good slice", sl: []string{}, exp: reflect.TypeOf("string")},
		{name: "all good ptr slice", sl: &[]string{}, exp: reflect.TypeOf("string")},
		{name: "all good array", sl: [1]string{}, exp: reflect.TypeOf("string")},
		{name: "all good ptr array", sl: &[1]string{}, exp: reflect.TypeOf("string")},
		{name: "nil", sl: nil, err: true},
		{name: "non-slice/array", sl: 42, err: true},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			act, err := SliceElemType(tt.sl)
			if tt.err {
				r.Error(err)
				return
			}

			r.NoError(err)
			r.Equal(tt.exp, act)

		})
	}

}
