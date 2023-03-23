package chezmoi

import (
	"testing"
)

var _ PersistentState = &MockPersistentState{}

func TestMockPersistentState(t *testing.T) {
	testPersistentState(t, func() PersistentState {
		return NewMockPersistentState()
	})
}
