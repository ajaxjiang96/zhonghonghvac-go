package zhonghongbywireprotocol

import (
	"fmt"
	"io"
	"time"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/zhonghong/zhonghongserial"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/zhonghong/zhonghongchecksum"

)

const (
	rtuMinSize = 4
	rtuMaxSize = 256

	rtuExceptionSize = 5
)

// RTUClientHandler implements Packager and Transporter interface.
type RTUClientHandler struct {
	rtuPackager
	rtuSerialTransporter
}



// NewRTUClientHandler allocates and initializes a RTUClientHandler.
func NewRTUClientHandler(address string) *RTUClientHandler {
	handler := &RTUClientHandler{}
	handler.address = address
	handler.IdleTimeout = serialIdleTimeout
	return handler
}

// RTUClient creates RTU client with default handler and given connect string.
func RTUClient(address string) Client {
	handler := NewRTUClientHandler(address)
	return NewClient(handler)
}

// rtuPackager implements Packager interface.
type rtuPackager struct {
	// SlaveId byte
}

// Encode encodes PDU in a RTU frame:
//
//	Header   : 1 byte
//	Function        : 1 byte
//	Data            : 0 up to 252 bytes
//	Checksum             : 1 byte
func (mb *rtuPackager) Encode(pdu *ProtocolDataUnit) (adu []byte, err error) {
	length := uint16(pdu.Address[0])
	if length > rtuMaxSize {
		err = fmt.Errorf("zhonghongprotocol: length of data '%v' must not be bigger than '%v'", length, rtuMaxSize)
		return
	}
	adu = make([]byte, length)

	adu[0] = pdu.Header
	adu[4] = pdu.FunctionCode
	copy(adu[1:4], pdu.Address)
	copy(adu[5:], pdu.Commands)

	checksum := zhonghongchecksum.Checksum(adu[0 : length-1])

	adu[length-1] = byte(checksum)
	return

}

// Verify verifies response length and slave id.
func (mb *rtuPackager) Verify(aduRequest []byte, aduResponse []byte) (err error) {
	length := len(aduResponse)
	// Minimum size (including address, function and checksum)
	if length < rtuMinSize {
		err = fmt.Errorf("zonghongprotocol: response length '%v' does not meet minimum '%v'", length, rtuMinSize)
		return
	}
	// Slave address must match
	if aduResponse[0] != aduRequest[0] {
		err = fmt.Errorf("zonghongprotocol: response slave id '%v' does not match request '%v'", aduResponse[0], aduRequest[0])
		return
	}
	return
}

// Decode extracts PDU from RTU frame and verify checksum.
func (mb *rtuPackager) Decode(adu []byte) (pdu *ProtocolDataUnit, err error) {
	length := len(adu)
	receivedChecksum := int(adu[len(adu)-1])
	computedChecksum := zhonghongchecksum.Checksum(adu[0 : len(adu)-1])

	if computedChecksum != receivedChecksum {
		err = fmt.Errorf("zonghongprotocol: response checksum '%v' does not match expected '%v'", receivedChecksum, computedChecksum)
		return
	}
	// Function code & data
	pdu = &ProtocolDataUnit{}
	pdu.Header = adu[0]
	pdu.FunctionCode = adu[4]
	pdu.Data = append(adu[1:4], adu[5:length-1]...)
	return
}

// rtuSerialTransporter implements Transporter interface.
type rtuSerialTransporter struct {
	zhonghongserial.SerialPort
}

func (mb *rtuSerialTransporter) Send(aduRequest []byte) (aduResponse []byte, err error) {
	// Make sure port is connected
	if err = mb.serialPort.connect(); err != nil {
		return
	}
	// Start the timer to close when idle
	mb.serialPort.lastActivity = time.Now()
	mb.serialPort.startCloseTimer()

	// Send the request
	mb.serialPort.logf("Zhonghong: sending % x\n", aduRequest)
	if _, err = mb.serialPort.port.Write(aduRequest); err != nil {
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
	n, err = io.ReadAtLeast(mb.serialPort.port, data[:], rtuMinSize)
	if err != nil {
		return
	}
	//if the function is correct
	if data[1] == function {
		//we read the rest of the bytes
		if n < bytesToRead {
			if bytesToRead > rtuMinSize && bytesToRead <= rtuMaxSize {
				if bytesToRead > n {
					n1, err = io.ReadFull(mb.serialPort.port, data[n:bytesToRead])
					n += n1
				}
			}
		}
	} else if data[1] == functionFail {
		//for error we need to read 5 bytes
		if n < rtuExceptionSize {
			n1, err = io.ReadFull(mb.serialPort.port, data[n:rtuExceptionSize])
		}
		n += n1
	}

	if err != nil {
		return
	}
	aduResponse = data[:n]
	mb.serialPort.logf("zonghongprotocol: received % x\n", aduResponse)
	return
}

// calculateDelay roughly calculates time needed for the next frame.
// See zonghongprotocol over Serial Line - Specification and Implementation Guide (page 13).
func (mb *rtuSerialTransporter) calculateDelay(chars int) time.Duration {
	var characterDelay, frameDelay int // us

	if mb.serialPort.BaudRate <= 0 || mb.serialPort.BaudRate > 19200 {
		characterDelay = 750
		frameDelay = 1750
	} else {
		characterDelay = 15000000 / mb.serialPort.BaudRate
		frameDelay = 35000000 / mb.serialPort.BaudRate
	}
	return time.Duration(characterDelay*chars+frameDelay) * time.Microsecond
}

func calculateResponseLength(adu []byte) int {
	length := rtuMinSize
	switch adu[1] {
	case FuncCodeReadGateway:
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
