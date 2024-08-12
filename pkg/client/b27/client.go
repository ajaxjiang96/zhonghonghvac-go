package b27

import (
	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

type Client struct {
	packager    protocol.Packager
	transporter protocol.Transporter
}

type ClientHandler interface {
	protocol.Packager
	protocol.Transporter
}

// NewB27Client creates a new Zhonghong client with given backend handler.
func NewClient(handler ClientHandler) *Client {
	return &Client{packager: handler, transporter: handler}
}
