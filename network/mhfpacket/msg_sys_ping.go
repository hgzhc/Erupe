package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysPing represents the MSG_SYS_PING
type MsgSysPing struct {
	AckHandle uint32
	Unk0      uint16
}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysPing) Opcode() network.PacketID {
	return network.MSG_SYS_PING
}

// Parse parses the packet from binary
func (m *MsgSysPing) Parse(bf *byteframe.ByteFrame) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint16()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgSysPing) Build(bf *byteframe.ByteFrame) error {
	bf.WriteUint32(m.AckHandle)
	bf.WriteUint16(m.Unk0)
	return nil
}