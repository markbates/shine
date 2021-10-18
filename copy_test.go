package shine

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CopySlice(t *testing.T) {
	t.Parallel()

	exp := Beatles{
		{Name: "John"},
		{Name: "Paul"},
		{Name: "George"},
		{Name: "Ringo"},
	}

	table := []struct {
		dest interface{}
		err  bool
		name string
		src  interface{}
	}{
		{name: "all good", src: exp, dest: &Beatles{}},
		{name: "bad dest", src: exp, dest: 42, err: true},
		{name: "nil dest", src: exp, dest: nil, err: true},
		{name: "nil nil", src: nil, dest: nil, err: true},
		{name: "nil src", src: nil, dest: &Beatles{}, err: true},
		{name: "non-ptr dest", src: exp, dest: Beatles{}, err: true},
		{name: "ptr src", src: &exp, dest: &Beatles{}, err: true},
		{name: "wrong dest type", src: exp, dest: &[]string{}, err: true},
		{name: "wrong src type", src: []string{}, dest: &Beatles{}, err: true},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			err := CopySlice(tt.dest, tt.src)
			if tt.err {
				r.Error(err)
				return
			}

			r.NoError(err)

			Equal(t, tt.src, tt.dest)

		})
	}

}
