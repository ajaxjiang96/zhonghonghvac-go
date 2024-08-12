package b19

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func IsValidAddr(addr string) (bool, error) {
	return regexp.Match("^[0-9]{1,2}(-[0-9]{1,2})?$", []byte(addr))
}

func ParseAddr(addr string) ([]byte, error) {
	if valid, err := IsValidAddr(addr); !valid {
		return nil, fmt.Errorf("invalid address: %s", addr)
	} else if err != nil {
		return nil, err
	}

	tokens := strings.Split(addr, "-")

	if len(tokens) == 1 {
		b, err := strconv.Atoi(tokens[0])
		if err != nil {
			return nil, err
		}
		return []byte{byte(b)}, nil
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
