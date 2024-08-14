package b19

import (
	"fmt"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

func (c *Client) PerformanceCheck(addr string) (results *protocol.ACPerformanceResponse, err error) {
	resp, err := c.performanceCheckOne(addr)
	if err != nil {
		return nil, err
	}
	return &resp.Performances[0], nil
}

func (c *Client) performanceCheckOne(addr string) (results *protocol.BatchACPerformanceResponse, err error) {
	addrBytes, err := ParseAddr(addr)
	if err != nil {
		return nil, err
	}

	var cmd []byte

	if len(addrBytes) == 1 {
		cmd = []byte{addrBytes[0], 0x50, 0x02, 0xFF, 0xFF, 0xFF}
	} else {
		cmd = []byte{addrBytes[0], 0x50, 0x02, 0x01, addrBytes[0], addrBytes[1]}
	}

	checksum := protocol.CalculateByteSum(cmd)
	cmd = append(cmd, checksum)
	resp, err := c.transporter.Send(cmd, c.packager)
	if err != nil {
		return nil, err
	}
	resPdu, err := c.packager.Decode(resp)
	if err != nil {
		return nil, err
	}
	if resPdu.FunctionCode != 0x50 {
		return nil, fmt.Errorf("unexpected function code: %x", resPdu.FunctionCode)
	}

	command := resPdu.Data[0]
	if command != 0x02 {
		return nil, fmt.Errorf("unexpected command: %x", command)
	}

	numDevices := resPdu.Data[1]
	data := resPdu.Data[2:]
	results = &protocol.BatchACPerformanceResponse{}
	results.Total = uint(numDevices)
	for i := 0; i < int(numDevices); i++ {
		extAddr := data[i*3]
		intAddr := data[i*3+1]
		addr := fmt.Sprintf("%d-%d", extAddr, intAddr)
		status := data[i*3+2]
		perf := protocol.ACPerformanceResponse{
			Addr:   addr,
			Status: protocol.ACStatus(status),
		}
		results.Performances = append(results.Performances, perf)

	}

	return
}

func (c *Client) StatusCheck(addr string) (results *protocol.ACStatusResponse, err error) {
	addrBytes, err := ParseAddr(addr)
	if err != nil {
		return nil, err
	}
	request := protocol.ProtocolDataUnit{
		Header:       addrBytes[0],
		FunctionCode: protocol.FuncCodeACStatus,
		Data:         []byte{byte(protocol.B19CheckOneACStatus), 0x01, addrBytes[0], addrBytes[1]},
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
	addrStr := fmt.Sprintf("%d-%d", resPdu.Data[2], resPdu.Data[3])
	data := resPdu.Data[4:]
	status := &protocol.ACStatusResponse{
		AC: protocol.AC{
			ExtAddr:   fmt.Sprintf("%d", resPdu.Data[2]),
			IntAddr:   fmt.Sprintf("%d", resPdu.Data[3]),
			Id:        addrStr,
			On:        data[0] == 0x01,
			Temp:      int(data[1]),
			Mode:      protocol.ACModeFromB19(protocol.ACModeB19((data[2]))),
			FanSpeed:  protocol.FanSpeed(data[3]),
			RoomTemp:  int(data[4]),
			Error:     data[5] != 0x00, // 0x00: no error, 0x01: error, need send command 0x04 for detail
			Direction: protocol.ACWindDir(data[6]),
			IsSlave:   data[7] == 0x01,
		},
	}

	return status, nil
}

func (c *Client) Control(addr string, data protocol.ACControlRequest) (results *protocol.ACControlResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) On(addr string) (results *protocol.ACControlResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) Off(addr string) (results *protocol.ACControlResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) TempControl(addr string, value int) (results *protocol.ACControlResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) ModeControl(addr string, value protocol.ACMode) (results *protocol.ACControlResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) WindSpeedControl(addr string, value protocol.FanSpeed) (results *protocol.ACControlResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) WindDirControl(addr string, value protocol.ACWindDir) (results *protocol.ACControlResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) ErrorCheck(addr string) (results string, err error) {
	panic("not implemented") // TODO: Implement
}
