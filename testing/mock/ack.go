package mock

// EmptyAcknowledgement MockEmptyAcknowledgement implements the exported.Acknowledgement interface and always returns an empty byte string as Response
type EmptyAcknowledgement struct {
	Response []byte
}

// NewMockEmptyAcknowledgement returns a new instance of MockEmptyAcknowledgement
func NewMockEmptyAcknowledgement() EmptyAcknowledgement {
	return EmptyAcknowledgement{
		Response: []byte{},
	}
}

// Success implements the Acknowledgement interface
func (ack EmptyAcknowledgement) Success() bool {
	return true
}

// Acknowledgement implements the Acknowledgement interface
func (ack EmptyAcknowledgement) Acknowledgement() []byte {
	return []byte{}
}
