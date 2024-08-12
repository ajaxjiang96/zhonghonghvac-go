package b27

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

func IsValidAddr(addr string) (bool, error) {
	return regexp.Match("^[0-9]{1,2}(-[0-9]{1,2})?$", []byte(addr))
}

func ParseAddr(addr string) ([]byte, error) {
	if valid, err := IsValidAddr(addr); !valid {
		return nil, err
	}

	tokens := strings.Split(addr, "-")

	if len(tokens) == 1 {
		b, err := strconv.Atoi(tokens[0])
		if err != nil {
			return nil, err
		}
		return []byte{0xFF, byte(b)}, nil
	} else {
		b1, err := strconv.Atoi(tokens[0])
		if err != nil {
			return nil, err
		}
		b2, err := strconv.Atoi(tokens[1])
		if err != nil {
			return nil, err
		}
		return []byte{byte(b1), byte(b2)}, nil
	}
}

func VerifyResponse(req protocol.ProtocolDataUnit, res protocol.ProtocolDataUnit) (bool, error) {
	if res.Header != 0xCC {
		return false, fmt.Errorf("response header is not 0xCC")
	}
	if res.Address[0] != req.Address[0] || res.Address[1] != req.Address[1] {
		return false, fmt.Errorf("response address does not match request address")
	}
	if res.FunctionCode != req.FunctionCode {
		return false, fmt.Errorf("response function code does not match request function code")
	}
	return true, nil
}

func CombineControl(status protocol.ACStatusResponse, control protocol.ACControlRequest) protocol.ACControlRequest {
	var res protocol.ACControlRequest
	if control.On != nil {
		res.On = control.On
	} else {
		res.On = &status.On
	}
	if control.Mode != nil {
		res.Mode = control.Mode
	} else {
		res.Mode = &status.Mode
	}
	if control.FanSpeed != nil {
		res.FanSpeed = control.FanSpeed
	} else {
		res.FanSpeed = &status.FanSpeed
	}
	if control.Temp != nil {
		res.Temp = control.Temp
	} else {
		res.Temp = &status.Temp
	}
	if control.Direction != nil {
		res.Direction = control.Direction
	} else {
		res.Direction = &status.Direction
	}
	return res
}

func ACControlToBytes(control protocol.ACControlRequest) []byte {
	var res []byte
	if *control.On {
		res = append(res, 0x01)
	} else {
		res = append(res, 0x00)
	}

	if control.Temp != nil {
		res = append(res, byte(*control.Temp))
	}

	if control.Mode != nil {
		res = append(res, byte(control.Mode.ToB27()))
	}
	if control.FanSpeed != nil {
		res = append(res, byte(*control.FanSpeed))
	}
	if control.Direction != nil {
		res = append(res, byte(*control.Direction))
	}
	return res
}
