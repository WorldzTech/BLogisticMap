// Modified Logistic Map
package core

type BFlow struct {
	Flow []float64
}

func NewBFlow(n int, r, x0, p, q float64) *BFlow {
	m := []float64{x0}

	for i := 1; i < n; i++ {
		d := q + (float64(i)+1)/float64(n)*(p-q)

		if (i+1)%2 == 1 {
			d = p - d + q
		}

		m = append(m, r*m[i-1]*(1-m[i-1])*d)
	}

	return &BFlow{
		Flow: m,
	}
}

func (bf *BFlow) GetBytesFlow() []byte {
	byteFlow := make([]byte, len(bf.Flow))

	for i, val := range bf.Flow {
		byteFlow[i] = byte(val*256)
	}

	return byteFlow
}