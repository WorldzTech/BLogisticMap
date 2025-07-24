// Cipher product

package blogistic

type BStreamerProduct struct {
	Bytes []byte
}

func (p *BStreamerProduct) AsString() string {
	return string(p.Bytes)
}