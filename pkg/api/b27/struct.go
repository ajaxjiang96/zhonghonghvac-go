package b27

import (
	"github.com/Yangsta911/zhonghonghvac-go/pkg/api"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

type ACPerformanceRequest struct {
	Addr string `json:"addr"`
}

type ACPerformanceResponse struct {
	Addr    string            `json:"addr"`
	IntAddr string            `json:"int_addr"`
	ExtAddr string            `json:"ext_addr"`
	ACBrand protocol.ACBrand  `json:"ac_brand"`
	Status  protocol.ACStatus `json:"status"`
}

type ACStatusRequest struct {
	Addr string `json:"addr"`
}

type ACStatusResponse struct {
	api.AC
}

type ACControlRequest struct {
	Addr      string             `json:"addr"`
	On        bool               `json:"on"`
	Temp      int                `json:"temp"`
	FanSpeed  protocol.FanSpeed  `json:"fan_speed"`
	Mode      protocol.ACMode    `json:"mode"`
	Direction protocol.ACWindDir `json:"direction"`
}

type ACControlResponse struct {
	Success bool `json:"success"`
}

type ACErrorCodeRequest struct {
	Addr string `json:"addr"`
}

type ACErrorCodeResponse struct {
	ErrorCode string `json:"error_code"`
}

// TODO: Ventilation and Floor Heating omitted
