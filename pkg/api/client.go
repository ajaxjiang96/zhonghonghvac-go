package api

import "github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"

type Client interface {
	// Bit access
	ReadGateway() (results *protocol.ProtocolDataUnit, err error)
	EditGateway(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	On(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	Off(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	TempControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	Control(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	WindSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	WindDirControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	NewAirOn(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	NewAirOff(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	NewAirModeControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	NewAirSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	NewAirErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	ErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	StatusCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FunctionCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error)
}
