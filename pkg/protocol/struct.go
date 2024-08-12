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

type AC struct {
	ExtAddr   string    `json:"ext_addr"`
	IntAddr   string    `json:"int_addr"`
	Id        string    `json:"id"` // [ext_addr]-[int_addr]
	On        bool      `json:"on"`
	Temp      uint      `json:"temp"`
	FanSpeed  FanSpeed  `json:"fan_speed"`
	Mode      ACMode    `json:"mode"` // B19 and B27 use different mode values
	RoomTemp  uint      `json:"room_temp"`
	ErrorCode byte      `json:"error_code"`
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
ErrorCode:	%d
Direction:	%s
IsSlave:	%t`,
		ac.Id,
		ac.On,
		ac.Temp,
		ac.FanSpeed,
		ac.Mode,
		ac.RoomTemp,
		ac.ErrorCode,
		ac.Direction,
		ac.IsSlave)
}

type ACControlRequest struct {
	On        *bool      `json:"on"`
	Temp      *uint      `json:"temp"`
	FanSpeed  *FanSpeed  `json:"fan_speed"`
	Mode      *ACMode    `json:"mode"` // B19 and B27 use different mode values
	Direction *ACWindDir `json:"direction"`
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
