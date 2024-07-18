package protocol

// List of header and function codes for the protocol

type FuncCode byte

const (
	// Bit access
	FuncCodeReadGateway                        FuncCode = 0xB0
	FuncCodeEditGateway                        FuncCode = 0xB1
	FuncCodeGatewayOnOff                       FuncCode = 0x31
	FuncCodeGatewayTemp                        FuncCode = 0x32
	FuncCodeGatewayControl                     FuncCode = 0x33
	FuncCodeGatewayWindSpeed                   FuncCode = 0x34
	FuncCodeACStatus                           FuncCode = 0x50
	FuncCodeGatewayWindDir                     FuncCode = 0x71
	FuncCodeGatewayNewAirOnOff                 FuncCode = 0x72
	FuncCodeGatewayNewAirMode                  FuncCode = 0x73
	FuncCodeGatewayNewAirSpeed                 FuncCode = 0x74
	FuncCodeGatewayFloorHeatingOnOff           FuncCode = 0x81
	FuncCodeGatewayFloorHeatingTemp            FuncCode = 0x82
	FuncCodeGatewayFloorHeatingControl         FuncCode = 0x83
	FuncCodeGatewayFloorHeatingAntiFreezeOnOff FuncCode = 0x84
	FuncCodePerformanceCheck                   FuncCode = 0x01
	FuncCodeStatusCheck                        FuncCode = 0x02
	FuncCodeOnOff                              FuncCode = 0x03
	FuncCodeErrorCheck                         FuncCode = 0x04
	FuncCodeFreshAirStatus                     FuncCode = 0x11
	FuncCodeFreshAirPerformance                FuncCode = 0x12
	FuncCodeFreshAirControl                    FuncCode = 0x13
	FuncCodeFreshAirErrorCheck                 FuncCode = 0x14
	FuncCodeFloorHeatingPerformance            FuncCode = 0x21
	FuncCodeFloorHeatingStatusCheck            FuncCode = 0x22
	FuncCodeFloorHeatingControlCheck           FuncCode = 0x24
	FuncCodeFloorHeatingOnOff                  FuncCode = 0x23
)

const (
	HeadCode            byte = 0xDD
	HeadCodeGateway     byte = 0xDD
	HeadCodeReadGateway byte = 0xFF
	ON                  byte = 0x01
	OFF                 byte = 0x00
)

// ProtocolDataUnit (PDU) is independent of underlying communication layers.
type ProtocolDataUnit struct {
	Header       byte
	FunctionCode FuncCode
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
