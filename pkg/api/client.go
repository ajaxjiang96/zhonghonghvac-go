package api

import "github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"

// Client interface for both B27 and B19 controllers
type Client interface {
	// Bit access
	ReadGateway() (results *protocol.ProtocolDataUnit, err error) // change back to hardcoded value
	EditGateway(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	On(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	Off(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	TempControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	ModeControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	Control(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	WindSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	WindDirControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirOn(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirOff(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirModeControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	ErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	StatusCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	PerformanceCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirPerformance(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirStatus(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingPerformance(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingStatus(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingOn(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingOff(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingTemp(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingAntiFreezeOn(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingAntiFreezeOff(data []uint16) (results *protocol.ProtocolDataUnit, err error)
}
