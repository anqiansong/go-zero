package generate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	err := Do(&Context{
		Type:   "User",
		Time:   false,
		Cache:  false,
		Output: t.TempDir(),
	})

	assert.Nil(t, err)
}
