package protocol_test

import (
	"encoding/hex"
	"testing"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
	"github.com/stretchr/testify/assert"
)

func TestParseReadGateway(t *testing.T) {
	byteArray, err := hex.DecodeString("ffff900c6ccbd7e1f965b5bea842e3b95c6000c0a801c9ffffff00c0a80101c0a801c815be270f01258002")
	assert.Nil(t, err)
	pdu := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeReadGateway,
		FunctionCode: protocol.FuncCodeReadGateway,
		Data:         byteArray,
	}

	actual, err := protocol.ParseReadGateway(pdu)
	assert.Nil(t, err)
	assert.Equal(t, protocol.GatewayInfo{
		DeviceId:   "900c6ccbd7e1f965b5bea842e3b95c60",
		Dhcp:       false,
		IpAddr:     "192.168.1.201",
		IpGateway:  "192.168.1.1",
		IpMask:     "255.255.255.0",
		LocalPort:  "9999",
		RemoteIp:   "192.168.1.200",
		RemotePort: "5566",
		SlaveId:    "1",
		BaudRate:   "9600",
		Validation: "2",
	}, actual)
}
