package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfAcquireGuildTresureSouvenir represents the MSG_MHF_ACQUIRE_GUILD_TRESURE_SOUVENIR
type MsgMhfAcquireGuildTresureSouvenir struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfAcquireGuildTresureSouvenir) Opcode() network.PacketID {
	return network.MSG_MHF_ACQUIRE_GUILD_TRESURE_SOUVENIR
}

// Parse parses the packet from binary
func (m *MsgMhfAcquireGuildTresureSouvenir) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfAcquireGuildTresureSouvenir) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}