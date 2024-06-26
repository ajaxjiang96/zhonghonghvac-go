package api

import "github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"

// ClientHandler is the interface that groups the Packager and Transporter methods.
type ClientHandler interface {
	protocol.Packager
	protocol.Transporter
}
