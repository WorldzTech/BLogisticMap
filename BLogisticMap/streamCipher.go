// Stream Cipher

package BLogisticMap

type BStreamer struct {
	Message []byte
	Flow *BFlow
}



func NewStreamer(message []byte, flow *BFlow) *BStreamer {
	return &BStreamer{
		Message: message,
		Flow: flow,
	}
}

func (s *BStreamer) Encrypt() *BStreamerProduct {
	bytesFlow := s.Flow.GetBytesFlow()
	encrypted := make([]byte, len(s.Message))

	for ind, val := range s.Message {
		encrypted[ind] = val ^ bytesFlow[ind]
	}

	return &BStreamerProduct{encrypted}
}