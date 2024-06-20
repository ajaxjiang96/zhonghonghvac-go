package zhonghongprotocolbywireprotocol

import (
	"encoding/binary"
	"fmt"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/zhonghong/zhonghongprotocol"
)

// ClientHandler is the interface that groups the Packager and Transporter methods.
type ClientHandlerByWire interface {
	zhonghongprotocol.Packager
	zhonghongprotocol.Transporter
}

type clientbywire struct {
	packager    zhonghongprotocol.Packager
	transporter zhonghongprotocol.Transporter
}

// NewClient creates a new Zhonghonh client with given backend handler.
func NewClientRemote(handler ClientHandlerByWire) ClientByWire {
	return &clientbywire{packager: handler, transporter: handler}
}

func (mb *clientbywire) FunctionCheck(address []uint16) (results *zhonghongprotocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	request := zhonghongprotocol.ProtocolDataUnit{
		Header:       zhonghongprotocol.HeadCodeFunctionCheck,
		FunctionCode: zhonghongprotocol.FuncCodeFunctionCheck,
		CommandType: "remote",
		Address: addressLen, 
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *clientbywire) StatusCheck(address []uint16) (results *zhonghongprotocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	request := zhonghongprotocol.ProtocolDataUnit{
		Header:       zhonghongprotocol.HeadCodeStatusCheck,
		FunctionCode: zhonghongprotocol.FuncCodeStatusCheck,
		CommandType: "remote",
		Address: addressLen, 
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *clientbywire) ControlOn(address []uint16, commands []uint16) (results *zhonghongprotocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	newArr = PrependUint16(commands, ON)
	commandsOn := dataBlockArray(newArr)
	request := zhonghongprotocol.ProtocolDataUnit{
		Header:       zhonghongprotocol.HeadCodeOnOff,
		FunctionCode: zhonghongprotocol.FuncCodeOnOff,
		CommandType: "remote",
		Address: addressLen,
		Commands: commandsOn,  
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (mb *clientbywire) ControlOff(address []uint16, commands []uint16) (results *zhonghongprotocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	newArr = PrependUint16(commands, OFF)
	commandsOff := dataBlockArray(newArr)
	request := zhonghongprotocol.ProtocolDataUnit{
		Header:       zhonghongprotocol.HeadCodeOnOff,
		FunctionCode: zhonghongprotocol.FuncCodeOnOff,
		CommandType: "remote",
		Address: addressLen,
		Commands: commandsOff,  
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *clientbywire) ErrorCheck(address []uint16) (results *zhonghongprotocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	request := zhonghongprotocol.ProtocolDataUnit{
		Header:       zhonghongprotocol.HeadCodeErrorCheck,
		FunctionCode: zhonghongprotocol.FuncCodeErrorCheck,
		CommandType: "remote",
		Address: addressLen, 
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *clientbywire) FreshAirCheck(address []uint16) (results *zhonghongprotocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	request := zhonghongprotocol.ProtocolDataUnit{
		Header:       zhonghongprotocol.HeadCodeFreshAirCheck,
		FunctionCode: zhonghongprotocol.FuncCodeFreshAirCheck,
		CommandType: "remote",
		Address: addressLen, 
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *clientbywire) FreshAirControlOn(address []uint16, commands []uint16) (results *zhonghongprotocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	newArr = PrependUint16(commands, ON)
	commandsOn := dataBlockArray(newArr)
	request := zhonghongprotocol.ProtocolDataUnit{
		Header:       zhonghongprotocol.HeadCodeOnOff,
		FunctionCode: zhonghongprotocol.FuncCodeOnOff,
		CommandType: "remote",
		Address: addressLen,
		Commands: commandsOn,  
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (mb *clientbywire) FreshAirControlOff(address []uint16, commands []uint16) (results *zhonghongprotocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	newArr = PrependUint16(commands, OFF)
	commandsOff := dataBlockArray(newArr)
	request := zhonghongprotocol.ProtocolDataUnit{
		Header:      zhonghongprotocol.HeadCodeOnOff,
		FunctionCode: zhonghongprotocol.FuncCodeOnOff,
		CommandType: "remote",
		Address: addressLen,
		Commands: commandsOff,  
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *clientbywire) FreshAirErrorCheck(address []uint16) (results *zhonghongprotocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	request := zhonghongprotocol.ProtocolDataUnit{
		Header:       zhonghongprotocol.HeadCodeFreshAirErrorCheck,
		FunctionCode: zhonghongprotocol.FuncCodeFreshAirErrorCheck,
		CommandType: "remote",
		Address: addressLen, 
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (mb *clientbywire) send(request *zhonghongprotocol.ProtocolDataUnit) (response *zhonghongprotocol.ProtocolDataUnit, err error) {
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
	response, err = mb.packager.DecodeRemote(aduResponse)
	if err != nil {
		return
	}
	// Check correct function code returned (exception)
	if response.FunctionCode != request.FunctionCode {
		err = responseError(response)
		return
	}
	if response.Data == nil || len(response.Data) == 0 {
		// Empty response
		err = fmt.Errorf("Zhonghong: response data is empty")
		return
	}
	return
}

func dataBlock(value ...uint16) []byte {
	data := make([]byte, 2*len(value))
	for i, v := range value {
		binary.BigEndian.PutUint16(data[i*2:], v)
	}
	return data
}

func dataBlockArray(arr []uint16) []byte {
	byteSlice := make([]byte, len(arr)*2)
	for i, v := range arr {
		binary.BigEndian.PutUint16(byteSlice[i*2:], v)
	}

	return byteSlice
}

func PrependUint16(slice []uint16, element uint16) []uint16 {
	newSlice := append([]uint16{element}, slice...)
	return newSlice
}

func responseError(response *zhonghongprotocol.ProtocolDataUnit) error {
	mbError := &zhonghongprotocol.ZhonghongError{FunctionCode: response.FunctionCode}
	if response.Data != nil && len(response.Data) > 0 {
		mbError.ExceptionCode = response.Data[0]
	}
	return mbError
}
