package encryption

type Custom struct {
}

func NewCustom() *Custom {
	return &Custom{}
}

func (c *Custom) Encrypt(p string) string {
	return p
}
