package zhonghongprotocol

type ClientRemote interface {
	// Bit access
	//todo change result to struct instead of byte
	FunctionCheck(address []uint16) (results *ProtocolDataUnit, err error)
	StatusCheck(address []uint16) (results *ProtocolDataUnit, err error)
	ControlOn(address []uint16, commands []uint16) (results *ProtocolDataUnit, err error)
	ControlOff(address []uint16, commands []uint16) (results *ProtocolDataUnit, err error)
	ErrorCheck(address []uint16) (results *ProtocolDataUnit, err error)
	FreshAirCheck(address []uint16) (results *ProtocolDataUnit, err error)
	FreshAirControlOn(address []uint16, commands []uint16) (results *ProtocolDataUnit, err error)
	FreshAirControlOff(address []uint16, commands []uint16) (results *ProtocolDataUnit, err error)
	FreshAirErrorCheck(address []uint16) (results *ProtocolDataUnit, err error)
}
