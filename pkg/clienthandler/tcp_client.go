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
func (handler *TCPClientHandler) Send(aduRequest []byte) (aduResponse []byte, err error) {
	// set an i/o deadline on the socket (read and write)
	err = handler.conn.SetDeadline(time.Now().Add(handler.timeout))
	if err != nil {
		return
	}

	_, err = handler.conn.Write(aduRequest)
	if err != nil {
		return
	}

	bytesToRead := calculateResponseLength(aduRequest)
	// function1 := aduRequest[1]
	// function2 := aduRequest[4]

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
}
