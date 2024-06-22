package protocol

import (
	"fmt"
)

const (
	// Bit access
	FuncCodeReadGateway        = 0xB0
	HeadCodeReadGateway        = 0xFF
	FuncCodeFunctionCheck      = 0x01
	HeadCodeFunctionCheck      = 0xDD
	FuncCodeStatusCheck        = 0x02
	HeadCodeStatusCheck        = 0xDD
	FuncCodeOnOff              = 0x03
	HeadCodeOnOff              = 0xDD
	FuncCodeErrorCheck         = 0x04
	HeadCodeErrorCheck         = 0xDD
	FuncCodeFreshAirCheck      = 0x12
	HeadCodeFreshAirCheck      = 0xDD
	FuncCodeFreshAirControl    = 0x13
	HeadCodeFreshAirControl    = 0xDD
	FuncCodeFreshAirErrorCheck = 0x14
	HeadCodeFreshAirErrorCheck = 0xDD
	ON                         = 0x01
	OFF                        = 0x00
)

const (
	ExceptionCodeIllegalFunction                    = 1
	ExceptionCodeIllegalDataAddress                 = 2
	ExceptionCodeIllegalDataValue                   = 3
	ExceptionCodeServerDeviceFailure                = 4
	ExceptionCodeAcknowledge                        = 5
	ExceptionCodeServerDeviceBusy                   = 6
	ExceptionCodeMemoryParityError                  = 8
	ExceptionCodeGatewayPathUnavailable             = 10
	ExceptionCodeGatewayTargetDeviceFailedToRespond = 11
)

// ZhonghongError implements error interface.
type ZhonghongError struct {
	FunctionCode  byte
	ExceptionCode byte
}

// Error converts known Zhonghong exception code to error message.
func (e *ZhonghongError) Error() string {
	var name string
	switch e.ExceptionCode {
	case ExceptionCodeIllegalFunction:
		name = "illegal function"
	case ExceptionCodeIllegalDataAddress:
		name = "illegal data address"
	case ExceptionCodeIllegalDataValue:
		name = "illegal data value"
	case ExceptionCodeServerDeviceFailure:
		name = "server device failure"
	case ExceptionCodeAcknowledge:
		name = "acknowledge"
	case ExceptionCodeServerDeviceBusy:
		name = "server device busy"
	case ExceptionCodeMemoryParityError:
		name = "memory parity error"
	case ExceptionCodeGatewayPathUnavailable:
		name = "gateway path unavailable"
	case ExceptionCodeGatewayTargetDeviceFailedToRespond:
		name = "gateway target device failed to respond"
	default:
		name = "unknown"
	}
	return fmt.Sprintf("Zhonghong: exception '%v' (%s), function '%v'", e.ExceptionCode, name, e.FunctionCode)
}

// ProtocolDataUnit (PDU) is independent of underlying communication layers.
type ProtocolDataUnit struct {
	Header       byte
	FunctionCode byte
	CommandType  string
	Data         []byte
	Address      []byte
	Commands     []byte
}

// Packager specifies the communication layer.
type Packager interface {
	Encode(pdu *ProtocolDataUnit) (adu []byte, err error)
	Decode(adu []byte) (pdu *ProtocolDataUnit, err error)
	Verify(aduRequest []byte, aduResponse []byte) (err error)
}

// Transporter specifies the transport layer.
type Transporter interface {
	Send(aduRequest []byte) (aduResponse []byte, err error)
}
