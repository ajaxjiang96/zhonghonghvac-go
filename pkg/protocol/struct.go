package protocol

import "fmt"

type GatewayInfo struct {
	DeviceId   string `json:"device_id"`
	Dhcp       bool   `json:"dhcp"`
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

func (gwInfo *GatewayInfo) String() {
	fmt.Printf("device_id:\t%s\ndhcp:\t\t%t\nip_addr:\t%s\nip_mask:\t%s\nip_gateway:\t%s\nremote_ip:\t%s\nremote_port:\t%s\nlocal_port:\t%s\nslave_id:\t%s\nbaud_rate:\t%s\nvalidation:\t%s\n", gwInfo.DeviceId, gwInfo.Dhcp, gwInfo.IpAddr, gwInfo.IpMask, gwInfo.IpGateway, gwInfo.RemoteIp, gwInfo.RemotePort, gwInfo.LocalPort, gwInfo.SlaveId, gwInfo.BaudRate, gwInfo.Validation)
}

type AC struct {
	ExtAddr   string    `json:"ext_addr"`
	IntAddr   string    `json:"int_addr"`
	Id        string    `json:"id"` // [ext_addr]-[int_addr]
	On        bool      `json:"on"`
	Temp      int       `json:"temp"`
	FanSpeed  FanSpeed  `json:"fan_speed"`
	Mode      ACMode    `json:"mode"` // B19 and B27 use different mode values
	RoomTemp  int       `json:"room_temp"`
	Error     bool      `json:"error"`
	Direction ACWindDir `json:"direction"`
	IsSlave   bool      `json:"is_slave"`
}

func (ac *AC) String() string {
	return fmt.Sprintf(`AC:		%s
On:		%t
Temp:		%d
FanSpeed:	%s
Mode:		%s
RoomTemp:	%d
Error:		%t
Direction:	%s
IsSlave:	%t`,
		ac.Id,
		ac.On,
		ac.Temp,
		ac.FanSpeed,
		ac.Mode,
		ac.RoomTemp,
		ac.Error,
		ac.Direction,
		ac.IsSlave)
}

type ACControlRequest struct {
	On        *bool      `json:"on,omitempty"`
	Temp      *int       `json:"temp,omitempty"`
	FanSpeed  *FanSpeed  `json:"fan_speed,omitempty"`
	Mode      *ACMode    `json:"mode,omitempty"` // B19 and B27 use different mode values
	Direction *ACWindDir `json:"direction,omitempty"`
}

type ACPerformanceResponse struct {
	Addr    string   `json:"addr"`
	IntAddr string   `json:"int_addr"`
	ExtAddr string   `json:"ext_addr"`
	ACBrand ACBrand  `json:"ac_brand"`
	Status  ACStatus `json:"status"`
}

type ACStatusRequest struct {
	Addr string `json:"addr"`
}

type ACStatusResponse struct {
	AC
}

type ACControlResponse struct {
	Success bool `json:"success"`
}

type ACErrorCodeResponse struct {
	ErrorCode string `json:"error_code"`
}

type BatchACStatusResponse struct {
	ACs   []AC `json:"ac_status"`
	Total int  `json:"total"`
}

type BatchACPerformanceResponse struct {
	Performances []ACPerformanceResponse `json:"performances"`
	Total        uint                    `json:"total"`
}

type ReadGatewayInfoResponse struct {
	GatewayInfo
}

type WriteGatewayInfoRequest struct {
	GatewayInfo
}

// For control multiple ACs, one command only
type ControlRequest struct {
	Addrs   []string `json:"addrs"`
	Command FuncCode `json:"command"`
	Value   int      `json:"value"`
}

// For control one AC, power, temperature, fan speed and mode should be set at the same time
type ControlOneRequest struct {
	Addr     string   `json:"addr"`
	On       bool     `json:"on"`
	Temp     int      `json:"temp"`
	FanSpeed FanSpeed `json:"fan_speed"`
	Mode     ACMode   `json:"mode"` // B19 and B27 use different mode values
}
