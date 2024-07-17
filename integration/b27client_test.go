//go:build integration
// +build integration

package integration_test

import (
	"encoding/hex"

	"testing"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/client"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/clienthandler"
	"github.com/stretchr/testify/assert"

	"go.bug.st/serial"
)

func TestB27ReadGatewayByRTU(t *testing.T) {
	// handler := zh.NewRTUClientHandler("/dev/tty.usbserial-AG0JG5OU")
	handler := clienthandler.NewRTUClientHandler("/dev/tty.usbserial-AG0JG5OU")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = serial.EvenParity
	handler.StopBits = serial.OneStopBit

	err := handler.Connect()
	if err != nil {
		t.Fail()
	}

	defer handler.Close()

	mb := client.NewB27Client(handler)
	rs, err := mb.ReadGateway()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "ffff900c6ccbd7e1f965b5bea842e3b95c6000c0a801c9ffffff00c0a80101c0a801c815be270f01258002", hex.EncodeToString(rs.Data))
}

func TestB27StatusCheckByTCP(t *testing.T) {
	handler, err := clienthandler.NewTCPClientHandler("192.168.1.220:4196", &clienthandler.B27Packager{})
	assert.NoError(t, err)
	client := client.NewB27Client(handler)

	// 0xff 外机 0x02 内机
	rs, err := client.StatusCheck(0x02)
	assert.NoError(t, err)
	assert.Equal(t, "011b0201001900", hex.EncodeToString(rs.Data))
}

func TestB27OffByTCP(t *testing.T) {
	handler, err := clienthandler.NewTCPClientHandler("192.168.1.220:4196", &clienthandler.B27Packager{})
	assert.NoError(t, err)
	client := client.NewB27Client(handler)

	// 0xff 外机 0x02 内机
	rs, err := client.Off(0x02)
	assert.NoError(t, err)
	assert.Equal(t, "01", hex.EncodeToString(rs.Data))
}

func TestB27OnByTCP(t *testing.T) {
	handler, err := clienthandler.NewTCPClientHandler("192.168.1.220:4196", &clienthandler.B27Packager{})
	assert.NoError(t, err)
	client := client.NewB27Client(handler)

	// 0xff 外机 0x02 内机
	rs, err := client.On(0x02)
	assert.NoError(t, err)
	assert.Equal(t, "01", hex.EncodeToString(rs.Data))
}
