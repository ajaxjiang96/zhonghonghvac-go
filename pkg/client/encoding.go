package client

import (
	"encoding/binary"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

func CalculateByteSum(data []byte) uint8 {
	var sum int64
	for _, b := range data {
		sum = sum + int64(b)
	}
	return uint8(sum % 256)
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

func responseError(response *protocol.ProtocolDataUnit) error {
	mbError := &protocol.ZhonghongError{FunctionCode: response.FunctionCode}
	if response.Data != nil && len(response.Data) > 0 {
		mbError.ExceptionCode = response.Data[0]
	}
	return mbError
}
