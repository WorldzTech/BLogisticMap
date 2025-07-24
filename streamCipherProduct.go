// Cipher product

package BLogisticMap

type BStreamerProduct struct {
	Bytes []byte
}

func (p *BStreamerProduct) AsString() string {
	return string(p.Bytes)
}