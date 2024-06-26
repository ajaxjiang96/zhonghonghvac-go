package clienthandler

import (
	"io"
	"net"
	"time"
)

type TCPClientHandler struct {
	RTUPackager
	TCPTransporter
}

func NewTCPClientHandler(address string) *TCPClientHandler {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil
	}

	return &TCPClientHandler{
		TCPTransporter: TCPTransporter{
			conn:    conn,
			timeout: 5 * time.Second,
		},
		RTUPackager: RTUPackager{},
	}
}

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

	aduResponse = make([]byte, bytesToRead)
	_, err = io.ReadFull(handler.conn, aduResponse)
	if err != nil {
		return
	}

	return
}
