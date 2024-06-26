package clienthandler

import (
	"io"
	"time"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/serial"
)

const (
	rtuMinSize       = 4
	rtuMaxSize       = 256
	rtuExceptionSize = 5
)

type RTUClientHandler struct {
	B19Packager
	rtuSerialTransporter
}

// NewRTUClientHandler allocates and initializes a RTUClientHandler.
func NewRTUClientHandler(address string) *RTUClientHandler {
	handler := &RTUClientHandler{}
	handler.Address = address
	handler.IdleTimeout = serial.SerialIdleTimeout
	return handler
}

// rtuSerialTransporter implements Transporter interface.
type rtuSerialTransporter struct {
	serial.SerialPort
}

func (mb *rtuSerialTransporter) Send(aduRequest []byte) (aduResponse []byte, err error) {
	// Make sure port is connected
	if err = mb.SerialPort.Connect(); err != nil {
		return
	}
	// Start the timer to close when idle
	mb.SerialPort.LastActivity = time.Now()
	mb.SerialPort.StartCloseTimer()

	// Send the request
	mb.SerialPort.Logf("modbus: sending % x\n", aduRequest)
	if _, err = mb.SerialPort.Port.Write(aduRequest); err != nil {
		return
	}
	function := aduRequest[1]
	functionFail := aduRequest[1] & 0x80
	bytesToRead := calculateResponseLength(aduRequest)
	time.Sleep(mb.calculateDelay(len(aduRequest) + bytesToRead))

	var n int
	var n1 int
	var data [rtuMaxSize]byte
	//We first read the minimum length and then read either the full package
	//or the error package, depending on the error status (byte 2 of the response)
	n, err = io.ReadAtLeast(mb.SerialPort.Port, data[:], rtuMinSize)
	if err != nil {
		return
	}
	//if the function is correct
	if data[1] == function {
		//we read the rest of the bytes
		if n < bytesToRead {
			if bytesToRead > rtuMinSize && bytesToRead <= rtuMaxSize {
				if bytesToRead > n {
					n1, err = io.ReadFull(mb.SerialPort.Port, data[n:bytesToRead])
					n += n1
				}
			}
		}
	} else if data[1] == functionFail {
		//for error we need to read 5 bytes
		if n < rtuExceptionSize {
			n1, err = io.ReadFull(mb.SerialPort.Port, data[n:rtuExceptionSize])
		}
		n += n1
	}

	if err != nil {
		return
	}
	aduResponse = data[:n]
	mb.SerialPort.Logf("zonghongprotocol: received % x\n", aduResponse)
	return
}

// calculateDelay roughly calculates time needed for the next frame.
// See zonghongprotocol over Serial Line - Specification and Implementation Guide (page 13).
func (mb *rtuSerialTransporter) calculateDelay(chars int) time.Duration {
	var characterDelay, frameDelay int // us

	if mb.SerialPort.BaudRate <= 0 || mb.SerialPort.BaudRate > 19200 {
		characterDelay = 750
		frameDelay = 1750
	} else {
		characterDelay = 15000000 / mb.SerialPort.BaudRate
		frameDelay = 35000000 / mb.SerialPort.BaudRate
	}
	return time.Duration(characterDelay*chars+frameDelay) * time.Microsecond
}

func calculateResponseLength(adu []byte) int {
	length := rtuMinSize
	switch adu[1] {
	case protocol.FuncCodeReadGateway:
		length = 46
	// case FuncCodeReadInputRegisters,
	// 	FuncCodeReadHoldingRegisters,
	// 	FuncCodeReadWriteMultipleRegisters:
	// 	count := int(binary.BigEndian.Uint16(adu[4:]))
	// 	length += 1 + count*2
	// case FuncCodeWriteSingleCoil,
	// 	FuncCodeWriteMultipleCoils,
	// 	FuncCodeWriteSingleRegister,
	// 	FuncCodeWriteMultipleRegisters:
	// 	length += 4
	// case FuncCodeMaskWriteRegister:
	// 	length += 6
	// case FuncCodeReadFIFOQueue:
	// undetermined
	default:
	}
	return length
}
