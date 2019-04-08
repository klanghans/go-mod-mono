package persistence_test

import (
	"github.com/klanghans/go-mod-mono/login/pkg/persistence"
	"testing"
)

type MemoryFixture struct {
	*gunit.Fixture
}

func TestMemoryFixture(t *testing.T) {
	gunit.Run(new(MemoryFixture), t)
}

func (m *MemoryFixture) Setup() {
}

func (m *MemoryFixture) TestCanWrite() {
	persistence := persistence.NewMemory()
	persistence = persistence
}