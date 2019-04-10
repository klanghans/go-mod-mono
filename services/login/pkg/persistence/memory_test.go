package persistence_test

import (
	"github.com/klanghans/go-mod-mono/services/login/pkg/persistence"
	"github.com/smartystreets/gunit"
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
	p := persistence.NewMemory()
	cypher := p.Encrypt("plain")
	m.AssertEqual("cypher", cypher)

}
