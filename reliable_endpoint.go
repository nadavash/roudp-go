package roudp

type ReliableEndpoint struct {
	sendBuff    *sequenceBuffer
	receiveBuff *sequenceBuffer
}

// header defines header info to include with reliable packets.
type header struct {
	sequence uint16
	ack      uint16
	ackBits  uint32
}
