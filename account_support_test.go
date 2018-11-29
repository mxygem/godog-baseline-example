package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBalance(t *testing.T) {
	tc := &TestContext{}

	SetBalance(float64(100.00), tc)

	assert.Equal(t, 100.00, tc.Balance)
}
