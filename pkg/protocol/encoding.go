package protocol

import (
	"encoding/binary"
)

// CalculateByteSum calculates the sum of a byte slice and returns the least significant byte.
func CalculateByteSum(data []byte) uint8 {
	var sum int64
	for _, b := range data {
		sum = sum + int64(b)
	}
	return uint8(sum % 256)
}

// dataBlockArray returns a byteSlice given a uint16 slice.
func dataBlockArray(arr []uint16) []byte {
	byteSlice := make([]byte, len(arr)*2)
	for i, v := range arr {
		binary.BigEndian.PutUint16(byteSlice[i*2:], v)
	}

	return byteSlice
}

// PrependUint16 prepends a uint16 number to a uint16 slice.
func PrependUint16(slice []uint16, element uint16) []uint16 {
	newSlice := append([]uint16{element}, slice...)
	return newSlice
}

func NormalEncode(data []uint16, funcCode FuncCode) ProtocolDataUnit {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := ProtocolDataUnit{
		Header:       HeadCode,
		FunctionCode: funcCode,
		Address:      addressLen,
	}
	return request
}

func B19NormalEncode(controllerAddr byte, funcCode FuncCode, data ...byte) ProtocolDataUnit {
	request := ProtocolDataUnit{
		Header:       controllerAddr,
		FunctionCode: funcCode,
		Data:         data,
	}
	return request
}

func B27NormalEncode(addr []byte, funcCode FuncCode, data ...byte) ProtocolDataUnit {
	request := ProtocolDataUnit{
		Header:       HeadCode,
		FunctionCode: funcCode,
		Address:      addr,
		Data:         data,
	}

	return request
}

func OnOffEncode(data []uint16, funccode FuncCode, OnOff uint16) ProtocolDataUnit {
	address := data[:2]
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	commands := data[2:]
	newArr = PrependUint16(commands, OnOff)
	commandsOff := dataBlockArray(newArr)
	request := ProtocolDataUnit{
		Header:       HeadCode,
		FunctionCode: FuncCodeFloorHeatingOnOff,
		Address:      addressLen,
		Commands:     commandsOff,
	}
	return request
}

func BytesToUint16Slice(b []byte) []uint16 {
	if len(b)%2 != 0 {
		panic("byte slice length must be a multiple of 2")
	}

	uint16Slice := make([]uint16, len(b)/2)

	for i := 0; i < len(uint16Slice); i++ {
		uint16Slice[i] = binary.LittleEndian.Uint16(b[i*2:])
	}

	return uint16Slice
}
