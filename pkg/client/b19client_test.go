//go:build integration
// +build integration

package client_test

import (
	"encoding/hex"

	"testing"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/client"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/clienthandler"
	"github.com/stretchr/testify/assert"

	"go.bug.st/serial"
)

func TestB19ReadGatewayByRTU(t *testing.T) {
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

	mb := client.NewB19Client(handler)
	rs, err := mb.ReadGateway()
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, "ffff900c6ccbd7e1f965b5bea842e3b95c6000c0a801c9ffffff00c0a80101c0a801c815be270f01258002", hex.EncodeToString(rs.Data))
}

func TestB19ReadGatewayByTCP(t *testing.T) {
	handler := clienthandler.NewTCPClientHandler("10.1.0.254:4196")
	client := client.NewB19Client(handler)

	rs, err := client.ReadGateway()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "ffff900c6ccbd7e1f965b5bea842e3b95c6000c0a801c9ffffff00c0a80101c0a801c815be270f01258002", hex.EncodeToString(rs.Data))
}
