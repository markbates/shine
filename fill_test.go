package shine

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_FillSlice(t *testing.T) {
	t.Parallel()

	exp := []interface{}{
		&Beatle{Name: "John"},
		&Beatle{Name: "Paul"},
		&Beatle{Name: "George"},
		&Beatle{Name: "Ringo"},
	}

	table := []struct {
		dest interface{}
		err  bool
		name string
		src  []interface{}
	}{
		{name: "all good", src: exp, dest: &Beatles{}},
		{name: "nil dest", src: exp, dest: nil, err: true},
		{name: "nil nil", src: nil, dest: nil, err: true},
		{name: "nil src", src: nil, dest: Beatles{}},
		{name: "non-ptr dest", src: exp, dest: Beatles{}, err: true},
		{name: "non-slice dest", src: exp, dest: 42, err: true},
		{name: "ptr non-slice dest", src: exp, dest: &Beatle{}, err: true},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			err := FillSlice(tt.dest, tt.src...)
			if tt.err {
				r.Error(err)
				return
			}

			r.NoError(err)

			Equal(t, tt.src, tt.dest)

		})
	}

}

func Test_FillSlice_Interfaces(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	beatles := Beatles{
		&Beatle{Name: "John"},
		&Beatle{Name: "Paul"},
		&Beatle{Name: "George"},
		&Beatle{Name: "Ringo"},
	}

	exp := []fmt.Stringer{}
	for _, b := range beatles {
		exp = append(exp, b)
	}

	var dest Beatles
	r.NoError(FillSlice(&dest, exp[0], exp[1], exp[2], exp[3]))

	Equal(t, exp, dest)
}
