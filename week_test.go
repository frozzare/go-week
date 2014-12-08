package week

import (
	"github.com/bmizerany/assert"
	"github.com/frozzare/go-is-type"
	"testing"
)

func TestWeek(t *testing.T) {
	assert.Equal(t, true, is.Int(Now()))
}
