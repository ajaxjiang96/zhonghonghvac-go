package b19

import (
	"github.com/Yangsta911/zhonghonghvac-go/pkg/api"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

type ACStatusResponse struct {
	ACs   []api.AC `json:"ac_status"`
	Total int      `json:"total"`
}

type GatewayInfo struct {
	DHCP       bool   `json:"dhcp"`
	IpAddr     string `json:"ip_addr"`
	IpMask     string `json:"ip_mask"`
	IpGateway  string `json:"ip_gateway"`
	RemoteIp   string `json:"remote_ip"`
	RemotePort string `json:"remote_port"`
	LocalPort  string `json:"local_port"`
	SlaveId    string `json:"slave_id"`
	BaudRate   string `json:"baud_rate"`
	Validation string `json:"validation"`
}

type ReadGatewayInfoResponse struct {
	GatewayInfo
}

type WriteGatewayInfoRequest struct {
	GatewayInfo
}

// For control multiple ACs, one command only
type ControlRequest struct {
	Addrs   []string          `json:"addrs"`
	Command protocol.FuncCode `json:"command"`
	Value   int               `json:"value"`
}

// For control one AC, power, temperature, fan speed and mode should be set at the same time
type ControlOneRequest struct {
	Addr     string            `json:"addr"`
	On       bool              `json:"on"`
	Temp     int               `json:"temp"`
	FanSpeed protocol.FanSpeed `json:"fan_speed"`
	Mode     protocol.ACMode   `json:"mode"`
}
