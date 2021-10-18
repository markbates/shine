package shine

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type Beatle struct {
	Name string
}

func (b Beatle) String() string {
	return b.Name
}

type Beatles []*Beatle

func Equal(t testing.TB, exp interface{}, act interface{}) {
	t.Helper()

	ev := reflect.ValueOf(exp)
	if ev.Kind() == reflect.Ptr {
		ev = ev.Elem()
	}

	av := reflect.ValueOf(act)
	if av.Kind() == reflect.Ptr {
		av = av.Elem()
	}

	r := require.New(t)

	es := fmt.Sprint(ev.Interface())
	as := fmt.Sprint(av.Interface())
	r.Equal(es, as)
}
