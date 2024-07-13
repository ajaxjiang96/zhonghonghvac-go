package protocol

// List of header and function codes for the protocol
const (
	// Bit access
	FuncCodeReadGateway                        = 0xB0
	HeadCodeReadGateway                        = 0xFF
	FuncCodeEditGateway                        = 0xB1
	HeadCodeGateway                            = 0xDD
	FuncCodeGatewayOnOff                       = 0x31
	FuncCodeGatewayTemp                        = 0x32
	FuncCodeGatewayControl                     = 0x33
	FuncCodeGatewayWindSpeed                   = 0x34
	FuncCodeGatewayWindDir                     = 0x71
	FuncCodeGatewayNewAirOnOff                 = 0x72
	FuncCodeGatewayNewAirMode                  = 0x73
	FuncCodeGatewayNewAirSpeed                 = 0x74
	FuncCodeGatewayFloorHeatingOnOff           = 0x81
	FuncCodeGatewayFloorHeatingTemp            = 0x82
	FuncCodeGatewayFloorHeatingControl         = 0x83
	FuncCodeGatewayFloorHeatingAntiFreezeOnOff = 0x84
	FuncCodePerformanceCheck                   = 0x01
	HeadCode                                   = 0xDD
	FuncCodeStatusCheck                        = 0x02
	FuncCodeOnOff                              = 0x03
	FuncCodeErrorCheck                         = 0x04
	FuncCodeFreshAirStatus                     = 0x11
	FuncCodeFreshAirPerformance                = 0x12
	FuncCodeFreshAirControl                    = 0x13
	FuncCodeFreshAirErrorCheck                 = 0x14
	FuncCodeFloorHeatingPerformance            = 0x21
	FuncCodeFloorHeatingStatusCheck            = 0x22
	FloorHeatingOnOff                          = 0x23
	FuncCodeFloorHeatingControlCheck           = 0x24
	ON                                         = 0x01
	OFF                                        = 0x00
)

// ProtocolDataUnit (PDU) is independent of underlying communication layers.
type ProtocolDataUnit struct {
	Header       byte
	FunctionCode byte
	CommandType  string
	Data         []byte
	Address      []byte
	Commands     []byte
}

// Packager specifies the communication layer.
type Packager interface {
	Encode(pdu *ProtocolDataUnit) (adu []byte, err error)
	Decode(adu []byte) (pdu *ProtocolDataUnit, err error)
	Verify(aduRequest []byte, aduResponse []byte) (err error)
}

// Transporter specifies the transport layer.
type Transporter interface {
	Send(aduRequest []byte) (aduResponse []byte, err error)
}
