package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsyncSliceAny(t *testing.T) {
	f := func(check func(t T) bool, given []T, expected bool) {
		s := AsyncSlice{data: given, workers: 2}
		actual := s.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t T) bool { return (t % 2) == 0 }

	f(isEven, []T{}, false)
	f(isEven, []T{1}, false)
	f(isEven, []T{1, 3}, false)
	f(isEven, []T{2}, true)
	f(isEven, []T{1, 2}, true)
	f(isEven, []T{1, 3, 5, 7, 9, 11}, false)
	f(isEven, []T{1, 3, 5, 7, 9, 12}, true)
}
