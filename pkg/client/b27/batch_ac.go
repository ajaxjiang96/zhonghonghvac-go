package b27

import "github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"

func (c *Client) BatchStatusCheck(devices []string) (results *protocol.BatchACStatusResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) BatchControl(devices []string, data protocol.ACControlRequest) (results *protocol.ACControlResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) BatchOn(devices []string) (results *protocol.ACControlResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) BatchOff(devices []string) (results *protocol.ACControlResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) BatchTempControl(devices []string, value uint) (results *protocol.ACControlResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) BatchModeControl(devices []string, value protocol.ACMode) (results *protocol.ACControlResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) BatchWindSpeedControl(devices []string, value protocol.FanSpeed) (results *protocol.ACControlResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) BatchWindDirControl(devices []string, value protocol.ACWindDir) (results *protocol.ACControlResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) BatchPerformanceCheck(devices []string) (results *protocol.BatchACPerformanceResponse, err error) {
	results = &protocol.BatchACPerformanceResponse{}
	for _, addr := range devices {
		res, err := c.PerformanceCheck(addr)
		if err != nil {
			continue
		}
		results.Performances = append(results.Performances, *res)
		results.Total++
	}
	return results, nil
}

func (c *Client) ErrorCheckAll() (results *protocol.ProtocolDataUnit, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) StatusCheckAll() (results *protocol.ProtocolDataUnit, err error) {
	panic("not implemented") // TODO: Implement
}
