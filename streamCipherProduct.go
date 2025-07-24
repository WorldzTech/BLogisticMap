// Cipher product

package core

type BStreamerProduct struct {
	Bytes []byte
}

func (p *BStreamerProduct) AsString() string {
	return string(p.Bytes)
}