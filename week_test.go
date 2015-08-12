package week

import (
	"testing"

	"github.com/frozzare/go-assert"
	"github.com/frozzare/go-is-type"
)

func TestWeek(t *testing.T) {
	assert.Equal(t, true, is.Int(Now()))
}
