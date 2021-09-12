package ch8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTripleStep(t *testing.T) {
	require.Equal(t, 4, TripleStep(3))
}
