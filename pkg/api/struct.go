package api

import "github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"

type AC struct {
	ExtAddr   string             `json:"ext_addr"`
	IntAddr   string             `json:"int_addr"`
	Id        string             `json:"id"` // [ext_addr]-[int_addr]
	On        bool               `json:"on"`
	Temp      int                `json:"temp"`
	FanSpeed  protocol.FanSpeed  `json:"fan_speed"`
	Mode      protocol.ACMode    `json:"mode"`
	RoomTemp  int                `json:"room_temp"`
	ErrorCode byte               `json:"error_code"`
	Direction protocol.ACWindDir `json:"direction"`
	IsSlave   bool               `json:"is_slave"`
}
