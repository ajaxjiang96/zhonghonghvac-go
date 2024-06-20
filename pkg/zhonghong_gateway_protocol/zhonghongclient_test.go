//go:build integration
// +build integration

package zhonghonggatewayprotocol_test

import (
	"encoding/hex"
	"net"
	"testing"

	zh "github.com/Yangsta911/zhonghonghvac-go/pkg/zhonghong_protocol"
	"github.com/stretchr/testify/assert"

	"go.bug.st/serial"
)

func TestReadGateway(t *testing.T) {
	handler := zh.NewRTUClientHandler("/dev/tty.usbserial-AG0JG5OU")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = serial.EvenParity
	handler.StopBits = serial.OneStopBit

	err := handler.Connect()
	if err != nil {
		t.Fail()
		return
	}

	defer handler.Close()

	mb := zh.NewClient(handler)
	rs, err := mb.ReadGateway()

	assert.Equal(t, "ffff900c6ccbd7e1f965b5bea842e3b95c6000c0a801c9ffffff00c0a80101c0a801c815be270f01258002", hex.EncodeToString(rs))
}

func TestReadGatewayByTCP(t *testing.T) {
	conn, err := net.Dial("tcp", "192.168.1.254:4196")
	if err != nil {
		t.Fail()
	}

	client := zhonghongprotocol.TCPClient(conn)

	rs, err := client.ReadGateway()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "ffff900c6ccbd7e1f965b5bea842e3b95c6000c0a801c9ffffff00c0a80101c0a801c815be270f01258002", hex.EncodeToString(rs))
}
