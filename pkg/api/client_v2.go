package api

import "github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"

// Client interface for both B27 and B19 controllers
type ClientV2 interface {
	GatewayClient
	ACBatchClient
	ACClient
}

type GatewayClient interface {
	ReadGateway(addr string) (results *protocol.GatewayInfo, err error) // change back to hardcoded value
	EditGateway(data []uint16) (results *protocol.GatewayInfo, err error)
}

type ACBatchClient interface {
	BatchStatusCheck(devices []string) (results *protocol.BatchACStatusResponse, err error)
	BatchControl(devices []string, data protocol.ACControlRequest) (results *protocol.ACControlResponse, err error)
	BatchOn(devices []string) (results *protocol.ACControlResponse, err error)
	BatchOff(devices []string) (results *protocol.ACControlResponse, err error)
	BatchTempControl(devices []string, value int) (results *protocol.ACControlResponse, err error)
	BatchModeControl(devices []string, value protocol.ACMode) (results *protocol.ACControlResponse, err error)
	BatchWindSpeedControl(devices []string, value protocol.FanSpeed) (results *protocol.ACControlResponse, err error)
	BatchWindDirControl(devices []string, value protocol.ACWindDir) (results *protocol.ACControlResponse, err error)
	ErrorCheckAll() (results *protocol.ProtocolDataUnit, err error)
	StatusCheckAll() (results *protocol.ProtocolDataUnit, err error)
	BatchPerformanceCheck(devices []string) (results *protocol.BatchACPerformanceResponse, err error)
}

type ACClient interface {
	PerformanceCheck(addr string) (results *protocol.ACPerformanceResponse, err error)
	StatusCheck(addr string) (results *protocol.ACStatusResponse, err error)
	Control(addr string, data protocol.ACControlRequest) (results *protocol.ACControlResponse, err error)
	On(addr string) (results *protocol.ACControlResponse, err error)
	Off(addr string) (results *protocol.ACControlResponse, err error)
	TempControl(addr string, value int) (results *protocol.ACControlResponse, err error)
	ModeControl(addr string, value protocol.ACMode) (results *protocol.ACControlResponse, err error)
	WindSpeedControl(addr string, value protocol.FanSpeed) (results *protocol.ACControlResponse, err error)
	WindDirControl(addr string, value protocol.ACWindDir) (results *protocol.ACControlResponse, err error)
	ErrorCheck(addr string) (results string, err error)
}

type FreshAirClient interface {
	FreshAirOn(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirOff(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirModeControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirPerformance(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FreshAirStatus(data []uint16) (results *protocol.ProtocolDataUnit, err error)
}

type FloorHeatingClient interface {
	FloorHeatingPerformance(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingStatus(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingOn(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingOff(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingTemp(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingAntiFreezeOn(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	FloorHeatingAntiFreezeOff(data []uint16) (results *protocol.ProtocolDataUnit, err error)
}
