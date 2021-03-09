package rand

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	require.Len(t, String(1000), 1000)
	require.Len(t, String(10), 10)
	require.Len(t, String(0), 0)
}
