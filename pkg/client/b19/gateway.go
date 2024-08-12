package b19

import (
	"fmt"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

func (c *Client) ReadGateway(addr string) (results *protocol.GatewayInfo, err error) {
	addrBytes, err := ParseAddr(addr)
	if err != nil {
		return nil, err
	}
	if len(addrBytes) != 1 {
		return nil, fmt.Errorf("invalid address: %s, read gateway only works on gateways", addr)
	}
	request := protocol.ProtocolDataUnit{
		Header:       addrBytes[0],
		FunctionCode: protocol.FuncCodeReadGateway,
		Data:         []byte{0x00, 0x00, 0x00, 0x00},
	}
	reqAdu, err := c.packager.Encode(&request)
	if err != nil {
		return nil, err
	}
	resp, err := c.transporter.Send(reqAdu, c.packager)
	if err != nil {
		return nil, err
	}
	resPdu, err := c.packager.Decode(resp)
	if err != nil {
		return nil, err
	}
	gwInfo, err := protocol.ParseReadGateway(*resPdu)
	if err != nil {
		return nil, err
	}

	return &gwInfo, nil

}

func (c *Client) EditGateway(data []uint16) (results *protocol.GatewayInfo, err error) {
	panic("not implemented") // TODO: Implement
}
