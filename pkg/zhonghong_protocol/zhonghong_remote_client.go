package zhonghongprotocol

import (
	"encoding/binary"
	"fmt"
)

// ClientHandler is the interface that groups the Packager and Transporter methods.
type ClientHandlerRemote interface {
	Packager
	Transporter
}

type clientremote struct {
	packager    Packager
	transporter Transporter
}

// NewClient creates a new Zhonghonh client with given backend handler.
func NewClientRemote(handler ClientHandlerRemote) ClientRemote {
	return &clientremote{packager: handler, transporter: handler}
}

func (mb *clientremote) FunctionCheck(address []uint16) (results *ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	request := ProtocolDataUnit{
		Header:       HeadCodeFunctionCheck,
		FunctionCode: FuncCodeFunctionCheck,
		CommandType: "remote",
		Address: addressLen, 
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *clientremote) StatusCheck(address []uint16) (results *ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	request := ProtocolDataUnit{
		Header:       HeadCodeStatusCheck,
		FunctionCode: FuncCodeStatusCheck,
		CommandType: "remote",
		Address: addressLen, 
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *clientremote) ControlOn(address []uint16, commands []uint16) (results *ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	newArr = PrependUint16(commands, ON)
	commandsOn := dataBlockArray(newArr)
	request := ProtocolDataUnit{
		Header:       HeadCodeOnOff,
		FunctionCode: FuncCodeOnOff,
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


func (mb *clientremote) ControlOff(address []uint16, commands []uint16) (results *ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	newArr = PrependUint16(commands, OFF)
	commandsOff := dataBlockArray(newArr)
	request := ProtocolDataUnit{
		Header:       HeadCodeOnOff,
		FunctionCode: FuncCodeOnOff,
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

func (mb *clientremote) ErrorCheck(address []uint16) (results *ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	request := ProtocolDataUnit{
		Header:       HeadCodeErrorCheck,
		FunctionCode: FuncCodeErrorCheck,
		CommandType: "remote",
		Address: addressLen, 
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *clientremote) FreshAirCheck(address []uint16) (results *ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	request := ProtocolDataUnit{
		Header:       HeadCodeFreshAirCheck,
		FunctionCode: FuncCodeFreshAirCheck,
		CommandType: "remote",
		Address: addressLen, 
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *clientremote) FreshAirControlOn(address []uint16, commands []uint16) (results *ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	newArr = PrependUint16(commands, ON)
	commandsOn := dataBlockArray(newArr)
	request := ProtocolDataUnit{
		Header:       HeadCodeOnOff,
		FunctionCode: FuncCodeOnOff,
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


func (mb *clientremote) FreshAirControlOff(address []uint16, commands []uint16) (results *ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	newArr = PrependUint16(commands, OFF)
	commandsOff := dataBlockArray(newArr)
	request := ProtocolDataUnit{
		Header:       HeadCodeOnOff,
		FunctionCode: FuncCodeOnOff,
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

func (mb *clientremote) FreshAirErrorCheck(address []uint16) (results *ProtocolDataUnit, err error) {
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	request := ProtocolDataUnit{
		Header:       HeadCodeFreshAirErrorCheck,
		FunctionCode: FuncCodeFreshAirErrorCheck,
		CommandType: "remote",
		Address: addressLen, 
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (mb *clientremote) send(request *ProtocolDataUnit) (response *ProtocolDataUnit, err error) {
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

func responseError(response *ProtocolDataUnit) error {
	mbError := &ZhonghongError{FunctionCode: response.FunctionCode}
	if response.Data != nil && len(response.Data) > 0 {
		mbError.ExceptionCode = response.Data[0]
	}
	return mbError
}
