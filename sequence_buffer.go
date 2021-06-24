package roudp

import "time"

const (
	BUFFER_SIZE = 1024
)

type packetData struct {
	sequence uint16
	acked    bool
	sendTime time.Time
}

type sequenceBuffer struct {
	buff []packetData
}

func newSequenceBuffer() *sequenceBuffer {
	return &sequenceBuffer{
		buff: make([]packetData, BUFFER_SIZE),
	}
}

func (sb *sequenceBuffer) getPacketData(sequence uint16) *packetData {
	idx := sequence % BUFFER_SIZE
	if sb.buff[idx].sequence == sequence {
		return &sb.buff[idx]
	}
	return nil
}

func (sb *sequenceBuffer) insertPacketData(sequence uint16) *packetData {
	idx := sequence % BUFFER_SIZE
	sb.buff[idx].sequence = sequence
	return &sb.buff[idx]
}
