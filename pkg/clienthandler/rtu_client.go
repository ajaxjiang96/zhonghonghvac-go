package clienthandler

import (
	"fmt"
	"io"
	"time"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/serial"
)

const (
	rtuMinSize = 6
	rtuMaxSize = 256
)

type RTUClientHandler struct {
	protocol.Packager
	rtuSerialTransporter
}

// Sends request via serial connection and retrieves the response.
func (mb *rtuSerialTransporter) Send(aduRequest []byte, packager protocol.Packager) (aduResponse []byte, err error) {
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
	function1 := aduRequest[1]
	function2 := aduRequest[4]
	bytesToRead := packager.CalculateResponseLength(aduRequest)
	if bytesToRead == -1 {
		time.Sleep(mb.calculateDelay(getMaxLength(32) + bytesToRead)) //32 is the max number of devices allowed
	} else {
		time.Sleep(mb.calculateDelay(len(aduRequest) + bytesToRead))

	}

	var n int
	var n1 int
	var data [rtuMaxSize]byte
	n, err = io.ReadAtLeast(mb.SerialPort.Port, data[:], rtuMinSize)
	if err != nil {
		return
	}
	if bytesToRead == -1 {
		bytesToRead = packager.VariableLengthCalculateResponseLength(aduRequest, uint(data[3]))
	}

	if data[1] == function1 || data[4] == function2 {
		if n < bytesToRead {
			if bytesToRead > rtuMinSize && bytesToRead <= rtuMaxSize {
				if bytesToRead > n {
					n1, err = io.ReadFull(mb.SerialPort.Port, data[n:bytesToRead])
					n += n1
				}
			}
		}
	} else {
		err = fmt.Errorf("zhonghongprotocol: response function code is invalid")
		return
	}

	if err != nil {
		return
	}
	aduResponse = data[:n]
	if len(aduResponse) != bytesToRead {
		err = fmt.Errorf("zhonghongprotocol: response length '%v' does not match expected '%v'", len(aduResponse), bytesToRead)
		return
	}
	mb.SerialPort.Logf("zhonghongprotocol: received % x\n", aduResponse)
	return
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

// calculateDelay roughly calculates time needed for the next frame.
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

func getMaxLength(devices int) int {
	return 5 + devices*10
}
