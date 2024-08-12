package protocol

// List of header and function codes for the protocol

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
	VariableLengthCalculateResponseLength(adu []byte, numDevices uint) (responseLength int)
	CalculateResponseLength(adu []byte) int
}

// Transporter specifies the transport layer.
type Transporter interface {
	Send(aduRequest []byte, packager Packager) (aduResponse []byte, err error)
}
