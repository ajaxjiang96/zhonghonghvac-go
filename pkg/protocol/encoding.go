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

func NormalEncode(data []uint16, funccode byte) ProtocolDataUnit {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := ProtocolDataUnit{
		Header:       HeadCode,
		FunctionCode: funccode,
		Address:      addressLen,
	}

	return request
}

func OnOffEncode(data []uint16, funccode byte, OnOff uint16) ProtocolDataUnit {
	address := data[:2]
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	commands := data[2:]
	newArr = PrependUint16(commands, OnOff)
	commandsOff := dataBlockArray(newArr)
	request := ProtocolDataUnit{
		Header:       HeadCode,
		FunctionCode: FloorHeatingOnOff,
		Address:      addressLen,
		Commands:     commandsOff,
	}
	return request
}
