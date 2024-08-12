package clienthandler

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

type TCPTransporter struct {
	conn    net.Conn
	timeout time.Duration
}

type TCPClientHandler struct {
	protocol.Packager
	TCPTransporter
}

// NewTCPClientHandler allocates and initializes a TCPClientHandler.
func NewTCPClientHandler(address string, packager protocol.Packager) (*TCPClientHandler, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &TCPClientHandler{
		TCPTransporter: TCPTransporter{
			conn:    conn,
			timeout: 5 * time.Second,
		},
		Packager: packager,
	}, nil
}

// Sends request via tcp connection and retrieves the response.
func (handler *TCPClientHandler) Send(aduRequest []byte, packager protocol.Packager) (aduResponse []byte, err error) {
	// set an i/o deadline on the socket (read and write)
	err = handler.conn.SetDeadline(time.Now().Add(handler.timeout))
	if err != nil {
		return
	}

	_, err = handler.conn.Write(aduRequest)
	if err != nil {
		return
	}

	bytesToRead := packager.CalculateResponseLength(aduRequest)
	function1 := aduRequest[1]
	// function2 := aduRequest[4]

	if bytesToRead != -1 {

		aduResponse = make([]byte, bytesToRead)
		_, err = io.ReadFull(handler.conn, aduResponse)
		if err != nil {
			return
		}

		// if aduResponse[1] == function1 && aduResponse[4] == function2 {
		// 	err = fmt.Errorf("zhonghongprotocol: response function code is invalid")
		// 	return
		// }

		if len(aduResponse) != bytesToRead {
			err = fmt.Errorf("zhonghongprotocol: response length '%v' does not match expected '%v'", len(aduResponse), bytesToRead)
			return
		}
		return
	} else {
		var n int
		var n1 int
		aduResponse = make([]byte, rtuMaxSize)
		n, err = io.ReadAtLeast(handler.conn, aduResponse[:], rtuMinSize)
		if err != nil {
			return
		}
		// TODO: Sleep for a calculated delay
		// time.Sleep(handler.calculateDelay(len(aduRequest) + bytesToRead))
		bytesToRead = handler.VariableLengthCalculateResponseLength(aduRequest, uint(aduResponse[3]))

		if aduResponse[1] == function1 {
			if n < bytesToRead {
				if bytesToRead > rtuMinSize && bytesToRead <= rtuMaxSize {
					if bytesToRead > n {
						n1, err = io.ReadFull(handler.conn, aduResponse[n:bytesToRead])
						n += n1
					}
				}
			}
		}
		aduResponse = aduResponse[:n]
		return
	}
}
