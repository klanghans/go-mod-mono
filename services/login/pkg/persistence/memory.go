package persistence

import "github.com/klanghans/go-mod-mono/lib/encryption"

// test comment
type Memory struct {
	crypt *encryption.Custom
}

func NewMemory() *Memory {
	return &Memory{
		crypt: encryption.NewCustom(),
	}
}

func (m *Memory) Encrypt(p string) string {
	return m.crypt.Encrypt("cypher")
}
