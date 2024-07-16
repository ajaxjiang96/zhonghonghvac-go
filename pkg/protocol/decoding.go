package protocol

import (
	"encoding/binary"
	"fmt"
	"net"
)

// Function to convert byte array to IP string
func bytesToIP(b []byte) string {
	return net.IP(b).String()
}

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

func ParseReadGateway(pdu ProtocolDataUnit) (gatewayInfo GatewayInfo, err error) {
	if len(pdu.Data) != 43 {
		return gatewayInfo, fmt.Errorf("pdu.Data length is not 43")
	}
	deviceId := pdu.Data[2:18]
	dhcp := pdu.Data[18]
	ipAddr := pdu.Data[19:23]
	ipMask := pdu.Data[23:27]
	ipGateway := pdu.Data[27:31]
	remoteIp := pdu.Data[31:35]
	remotePort := pdu.Data[35:37]
	localPort := pdu.Data[37:39]
	slaveId := pdu.Data[39]
	baudRate := pdu.Data[40:42]
	validation := pdu.Data[42]

	return GatewayInfo{
		DeviceId:   string(deviceId),
		Dhcp:       dhcp == 0x01,
		IpAddr:     bytesToIP(ipAddr),
		IpMask:     bytesToIP(ipMask),
		IpGateway:  bytesToIP(ipGateway),
		RemoteIp:   bytesToIP(remoteIp),
		RemotePort: fmt.Sprintf("%d", binary.BigEndian.Uint16(remotePort)),
		LocalPort:  fmt.Sprintf("%d", binary.BigEndian.Uint16(localPort)),
		SlaveId:    fmt.Sprintf("%d", slaveId),
		BaudRate:   fmt.Sprintf("%d", binary.BigEndian.Uint16(baudRate)),
		Validation: fmt.Sprintf("%d", validation),
	}, nil
}
