package client

import (
	"fmt"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/api"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

// ClientHandler is the interface that groups the Packager and Transporter methods.
type ClientHandler interface {
	protocol.Packager
	protocol.Transporter
}

// 中弘线控器 B27（小超人）
type b27client struct {
	packager    protocol.Packager
	transporter protocol.Transporter
}

// NewB27Client creates a new Zhonghong client with given backend handler.
func NewB27Client(handler ClientHandler) api.Client {
	return &b27client{packager: handler, transporter: handler}
}

// PerformanceCheck returns performances statistics of the specified HVAC device
func (mb *b27client) PerformanceCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.NormalEncode(data, protocol.FuncCodePerformanceCheck)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// StatusCheck returns status of the specified HVAC device
func (mb *b27client) StatusCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.NormalEncode(data, protocol.FuncCodeStatusCheck)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// On turns on specified HVAC device
func (mb *b27client) On(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.OnOffEncode(data, protocol.FuncCodeOnOff, protocol.ON)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Off turns off specified HVAC device
func (mb *b27client) Off(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.OnOffEncode(data, protocol.FuncCodeOnOff, protocol.OFF)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ErrorCheck returns the error status code of the specified HVAC device
func (mb *b27client) ErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.NormalEncode(data, protocol.FuncCodeErrorCheck)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirPerformance returns the performance statistics of the specified Fresh Air ventilation device
func (mb *b27client) FreshAirPerformance(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.NormalEncode(data, protocol.FuncCodeFreshAirPerformance)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirStatus returns the status of the specified Fresh Air ventilation device
func (mb *b27client) FreshAirStatus(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.NormalEncode(data, protocol.FuncCodeFreshAirStatus)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirModeControl selects the mode of the specified Fresh Air ventilation device
func (mb *b27client) FreshAirModeControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.NormalEncode(data, protocol.FuncCodeFreshAirControl)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirOn turns on the specified Fresh Air ventilation device
func (mb *b27client) FreshAirOn(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.OnOffEncode(data, protocol.FuncCodeFreshAirControl, protocol.ON)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirOn turns off the specified Fresh Air ventilation device
func (mb *b27client) FreshAirOff(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.OnOffEncode(data, protocol.FuncCodeFreshAirControl, protocol.OFF)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirErrorCheck returns the error status of the specified Fresh Air ventilation device
func (mb *b27client) FreshAirErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.NormalEncode(data, protocol.FuncCodeFreshAirErrorCheck)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingPerformance returns the performance statistics of the specified Floor Heating device
func (mb *b27client) FloorHeatingPerformance(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.NormalEncode(data, protocol.FuncCodeFloorHeatingPerformance)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingStatus returns the status of the specified Floor Heating device
func (mb *b27client) FloorHeatingStatus(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.NormalEncode(data, protocol.FuncCodeFloorHeatingStatusCheck)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingOn turns on the specified Floor Heating device
func (mb *b27client) FloorHeatingOn(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.OnOffEncode(data, protocol.FloorHeatingOnOff, protocol.ON)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingOff turns off the specified Floor Heating device
func (mb *b27client) FloorHeatingOff(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.OnOffEncode(data, protocol.FloorHeatingOnOff, protocol.OFF)
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *b27client) ReadGateway() (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) Control(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) EditGateway(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) FreshAirSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) TempControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) WindDirControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) WindSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) FloorHeatingAntiFreezeOn(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) FloorHeatingAntiFreezeOff(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) FloorHeatingControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) FloorHeatingTemp(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

// B27Client sends the specified control command to the specified device
func (mb *b27client) send(request *protocol.ProtocolDataUnit) (response *protocol.ProtocolDataUnit, err error) {
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
