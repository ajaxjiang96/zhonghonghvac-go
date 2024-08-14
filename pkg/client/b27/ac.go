package b27

import (
	"fmt"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

func (c *Client) PerformanceCheck(addr string) (results *protocol.ACPerformanceResponse, err error) {
	addrBytes, err := ParseAddr(addr)
	if err != nil {
		return nil, err
	}
	request := protocol.B27NormalEncode(addrBytes, protocol.FuncCodePerformanceCheck)
	adu, err := c.packager.Encode(&request)
	if err != nil {
		return nil, err
	}
	resp, err := c.transporter.Send(adu, c.packager)

	if err != nil {
		return nil, err
	}

	resPdu, err := c.packager.Decode(resp)
	if err != nil {
		return nil, err
	}

	valid, err := VerifyResponse(request, *resPdu)
	if !valid || err != nil {
		return nil, err
	}

	var id string
	if resPdu.Address[0] == 0xFF {
		id = fmt.Sprintf("%d", resPdu.Address[1])
	} else {
		id = fmt.Sprintf("%d-%d", resPdu.Address[0], resPdu.Address[1])
	}

	return &protocol.ACPerformanceResponse{
		ExtAddr: string(resPdu.Address[0]), // through AC
		IntAddr: string(resPdu.Address[1]), // through AC
		Addr:    id,
		ACBrand: protocol.ACBrand(resPdu.Data[0]),
		Status:  protocol.ACStatus(resPdu.Data[1]),
	}, nil
}

func (c *Client) StatusCheck(addr string) (results *protocol.ACStatusResponse, err error) {
	addrBytes, err := ParseAddr(addr)
	if err != nil {
		return nil, err
	}
	request := protocol.B27NormalEncode(addrBytes, protocol.FuncCodeStatusCheck)
	adu, err := c.packager.Encode(&request)
	if err != nil {
		return nil, err
	}
	resp, err := c.transporter.Send(adu, c.packager)
	if err != nil {
		return nil, err
	}

	resPdu, err := c.packager.Decode(resp)
	if err != nil {
		return nil, err
	}

	if len(resPdu.Data) != 7 {
		return nil, fmt.Errorf("response data length is not 7")
	}

	var addrStr string

	if resPdu.Address[0] == 0xFF {
		addrStr = fmt.Sprintf("%d", resPdu.Address[1])
	} else {
		addrStr = fmt.Sprintf("%d-%d", resPdu.Address[0], resPdu.Address[1])
	}

	return &protocol.ACStatusResponse{
		AC: protocol.AC{
			ExtAddr:   fmt.Sprintf("%d", resPdu.Address[0]),
			IntAddr:   fmt.Sprintf("%d", resPdu.Address[1]),
			Id:        addrStr,
			On:        resPdu.Data[0] == 0x01,
			Temp:      int(resPdu.Data[1]),
			Mode:      protocol.ACModeFromB27(protocol.ACModeB27((resPdu.Data[2]))),
			FanSpeed:  protocol.FanSpeed(resPdu.Data[3]),
			Direction: protocol.ACWindDir(resPdu.Data[4]),
			RoomTemp:  int(resPdu.Data[5]),
			Error:     resPdu.Data[6] != 0x00, // 0x00: no error, 0x01: error, need send command 0x04 for detail
		},
	}, nil
}

func (c *Client) Control(addr string, data protocol.ACControlRequest) (results *protocol.ACControlResponse, err error) {
	addrBytes, err := ParseAddr(addr)
	if err != nil {
		return nil, err
	}
	status, err := c.StatusCheck(addr)
	if err != nil {
		return nil, err
	}

	control := CombineControl(*status, data)

	request := protocol.B27NormalEncode(addrBytes, protocol.FuncCodeControl, ACControlToBytes(control)...)
	adu, err := c.packager.Encode(&request)
	if err != nil {
		return nil, err
	}
	resp, err := c.transporter.Send(adu, c.packager)
	if err != nil {
		return nil, err
	}

	resPdu, err := c.packager.Decode(resp)
	if err != nil {
		return nil, err
	}

	if len(resPdu.Data) != 1 {
		return nil, fmt.Errorf("response data length is not 1")
	}

	return &protocol.ACControlResponse{
		Success: resPdu.Data[0] == 0x01,
	}, nil
}

func (c *Client) On(addr string) (results *protocol.ACControlResponse, err error) {
	on := true
	return c.Control(addr, protocol.ACControlRequest{
		On: &on,
	})
}

func (c *Client) Off(addr string) (results *protocol.ACControlResponse, err error) {
	on := false
	return c.Control(addr, protocol.ACControlRequest{
		On: &on,
	})
}

func (c *Client) TempControl(addr string, value int) (results *protocol.ACControlResponse, err error) {
	return c.Control(addr, protocol.ACControlRequest{
		Temp: &value,
	})
}

func (c *Client) ModeControl(addr string, value protocol.ACMode) (results *protocol.ACControlResponse, err error) {
	return c.Control(addr, protocol.ACControlRequest{
		Mode: &value,
	})
}

func (c *Client) WindSpeedControl(addr string, value protocol.FanSpeed) (results *protocol.ACControlResponse, err error) {
	return c.Control(addr, protocol.ACControlRequest{
		FanSpeed: &value,
	})
}

func (c *Client) WindDirControl(addr string, value protocol.ACWindDir) (results *protocol.ACControlResponse, err error) {
	return c.Control(addr, protocol.ACControlRequest{
		Direction: &value,
	})
}

func (c *Client) ErrorCheck(addr string) (result string, err error) {
	addrBytes, err := ParseAddr(addr)
	if err != nil {
		return "", err
	}
	request := protocol.B27NormalEncode(addrBytes, protocol.FuncCodeErrorCheck)
	adu, err := c.packager.Encode(&request)
	if err != nil {
		return "", err
	}
	resp, err := c.transporter.Send(adu, c.packager)
	if err != nil {
		return "", err
	}

	resPdu, err := c.packager.Decode(resp)
	if err != nil {
		return "", err
	}

	codeLen := uint(resPdu.Data[0])
	if codeLen != uint(len(resPdu.Data)-1) {
		return "", fmt.Errorf("message length is not correct")
	}

	errorCode := string(resPdu.Data[1:])

	return errorCode, nil
}
