package chezmoi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var _ PersistentState = &SQLitePersistentState{}

func TestSQLitePersistentState(t *testing.T) {
	testPersistentState(t, func() PersistentState {
		s, err := NewSQLitePersistentState(":memory:")
		require.NoError(t, err)
		return s
	})
}
