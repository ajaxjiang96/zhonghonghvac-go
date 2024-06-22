package b27protocol

import (
	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

type Clientb27 interface {
	// Bit access
	//todo change result to struct instead of byte
	FunctionCheck(address []uint16) (results *protocol.ProtocolDataUnit, err error)
	StatusCheck(address []uint16) (results *protocol.ProtocolDataUnit, err error)
	ControlOn(address []uint16, commands []uint16) (results *protocol.ProtocolDataUnit, err error)
	ControlOff(address []uint16, commands []uint16) (results *protocol.ProtocolDataUnit, err error)
	ErrorCheck(address []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirCheck(address []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirControlOn(address []uint16, commands []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirControlOff(address []uint16, commands []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirErrorCheck(address []uint16) (results *protocol.ProtocolDataUnit, err error)
}
