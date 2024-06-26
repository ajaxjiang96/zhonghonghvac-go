package clienthandler

import (
	"io"
	"net"
	"time"
)

type TCPTransporter struct {
	conn    net.Conn
	timeout time.Duration
}

func (handler *TCPTransporter) Send(aduRequest []byte) (aduResponse []byte, err error) {
	// set an i/o deadline on the socket (read and write)
	err = handler.conn.SetDeadline(time.Now().Add(handler.timeout))
	if err != nil {
		return
	}

	_, err = handler.conn.Write(aduRequest)
	if err != nil {
		return
	}

	// bytesToRead := calculateResponseLength(aduRequest)

	// aduResponse = make([]byte, bytesToRead)
	_, err = io.ReadFull(handler.conn, aduResponse)
	if err != nil {
		return
	}

	return
}
