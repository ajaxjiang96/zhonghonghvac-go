package b27protocol

import (
	"io"
	"net"
	"time"
)

type TCPClientHandler struct {
	tcpTransporter
	rtuPackager
}

type tcpTransporter struct {
	socket  net.Conn
	timeout time.Duration
}

func TCPClient(socket net.Conn) Clientb27 {
	handler := newTCPClientHandler(socket)
	return NewClient(handler)
}

func newTCPClientHandler(conn net.Conn) *TCPClientHandler {
	return &TCPClientHandler{
		rtuPackager: rtuPackager{},
		tcpTransporter: tcpTransporter{
			socket:  conn,
			timeout: 5 * time.Second,
		},
	}
}

func (handler *TCPClientHandler) Send(aduRequest []byte) (aduResponse []byte, err error) {
	// set an i/o deadline on the socket (read and write)
	err = handler.socket.SetDeadline(time.Now().Add(handler.timeout))
	if err != nil {
		return
	}

	_, err = handler.socket.Write(aduRequest)
	if err != nil {
		return
	}

	bytesToRead := calculateResponseLength(aduRequest)

	aduResponse = make([]byte, bytesToRead)
	_, err = io.ReadFull(handler.socket, aduResponse)
	if err != nil {
		return
	}

	return
}
