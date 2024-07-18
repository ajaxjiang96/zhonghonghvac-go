package client

import (
	"fmt"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

// ClientHandler is the interface that groups the Packager and Transporter methods.
type ClientHandler interface {
	protocol.Packager
	protocol.Transporter
}

// 中弘线控器 B27（小超人）
type B27client struct {
	packager    protocol.Packager
	transporter protocol.Transporter
}

// NewB27Client creates a new Zhonghong client with given backend handler.
func NewB27Client(handler ClientHandler) *B27client {
	return &B27client{packager: handler, transporter: handler}
}

// PerformanceCheck returns performances statistics of the specified HVAC device
func (mb *B27client) PerformanceCheck(addr byte, data ...byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodePerformanceCheck, data...)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// StatusCheck returns status of the specified HVAC device
func (mb *B27client) StatusCheck(addr byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeStatusCheck)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// On turns on specified HVAC device
func (mb *B27client) On(addr byte, data ...byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeOnOff, protocol.ON)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Off turns off specified HVAC device
func (mb *B27client) Off(addr byte, data ...byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeOnOff, protocol.OFF)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ErrorCheck returns the error status code of the specified HVAC device
func (mb *B27client) ErrorCheck(addr byte, data ...byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeErrorCheck, data...)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirPerformance returns the performance statistics of the specified Fresh Air ventilation device
func (mb *B27client) FreshAirPerformance(addr byte, data ...byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeErrorCheck, data...)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirStatus returns the status of the specified Fresh Air ventilation device
func (mb *B27client) FreshAirStatus(addr byte, data ...byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeFreshAirStatus, data...)

	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirModeControl selects the mode of the specified Fresh Air ventilation device
func (mb *B27client) FreshAirModeControl(addr byte, data ...byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeFreshAirControl, data...)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirOn turns on the specified Fresh Air ventilation device
func (mb *B27client) FreshAirOn(addr byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeFreshAirControl, protocol.ON)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirOn turns off the specified Fresh Air ventilation device
func (mb *B27client) FreshAirOff(addr byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeFreshAirControl, protocol.OFF)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirErrorCheck returns the error status of the specified Fresh Air ventilation device
func (mb *B27client) FreshAirErrorCheck(addr byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeFreshAirErrorCheck)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingPerformance returns the performance statistics of the specified Floor Heating device
func (mb *B27client) FloorHeatingPerformance(addr byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeFloorHeatingPerformance)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingStatus returns the status of the specified Floor Heating device
func (mb *B27client) FloorHeatingStatus(addr byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeFloorHeatingStatusCheck)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingOn turns on the specified Floor Heating device
func (mb *B27client) FloorHeatingOn(addr byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeFloorHeatingOnOff, protocol.ON)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingOff turns off the specified Floor Heating device
func (mb *B27client) FloorHeatingOff(addr byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeFloorHeatingOnOff, protocol.OFF)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *B27client) ReadGateway() (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *B27client) Control(addr byte, data ...byte) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.B27NormalEncode([]byte{addr}, protocol.FuncCodeOnOff, data...)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *B27client) EditGateway(addr byte) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *B27client) FreshAirSpeedControl(addr byte) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *B27client) TempControl(data []byte) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *B27client) WindDirControl(data []byte) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *B27client) WindSpeedControl(data []byte) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *B27client) FloorHeatingAntiFreezeOn(data []byte) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *B27client) FloorHeatingAntiFreezeOff(data []byte) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *B27client) FloorHeatingControl(data []byte) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *B27client) FloorHeatingTemp(data []byte) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

// B27Client sends the specified control command to the specified device
func (mb *B27client) send(request *protocol.ProtocolDataUnit) (response *protocol.ProtocolDataUnit, err error) {
	aduRequest, err := mb.packager.Encode(request)
	if err != nil {
		return
	}
	aduResponse, err := mb.transporter.Send(aduRequest)
	if err != nil {
		return
	}
	if err = mb.packager.Verify(aduRequest, aduResponse); err != nil {
		return
	}
	response, err = mb.packager.Decode(aduResponse)
	if err != nil {
		return
	}
	// Check correct function code returned (exception)
	if response.FunctionCode != request.FunctionCode {
		err = fmt.Errorf("zhonghong-b27 client: response function code does not match request")
		return
	}
	if response.Data == nil || len(response.Data) == 0 {
		// Empty response
		err = fmt.Errorf("zhonghong-b27 client: response data is empty")
		return
	}
	return
}
