package clienthandler

import (
	"fmt"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

const (
	AddressBytesLength = 2
)

type B27Packager struct {
}

// Encode encodes PDU into a Modbus frame:
//
//	Header   		: 1 byte
//	Function        : 1 byte
//	Data            : 0 up to 252 bytes
//	Checksum        : 1 byte
func (p *B27Packager) Encode(pdu *protocol.ProtocolDataUnit) (adu []byte, err error) {
	length := 5 + len(pdu.Data) + 1

	if length > rtuMaxSize {
		err = fmt.Errorf("b27-packager: length of data '%v' must not be bigger than '%v'", length, rtuMaxSize)
		return
	}
	adu = make([]byte, length)

	adu[0] = pdu.Header
	adu[1] = byte(length)
	if len(pdu.Address) == 1 {
		copy(adu[2:4], []byte{0xff, pdu.Address[0]})
	} else {
		copy(adu[2:4], pdu.Address)
	}
	adu[4] = byte(pdu.FunctionCode)
	copy(adu[5:5+len(pdu.Data)], pdu.Data)
	adu[length-1] = protocol.CalculateByteSum(adu[0 : length-1])
	return
}

// Decode extracts PDU from RTU frame and verify Checksum.
func (p *B27Packager) Decode(adu []byte) (pdu *protocol.ProtocolDataUnit, err error) {
	length := len(adu)
	receivedChecksum := uint8(adu[length-1])
	computedChecksum := protocol.CalculateByteSum(adu[0 : length-1])

	if computedChecksum != receivedChecksum {
		err = fmt.Errorf("b27-pack: response checksum '%v' does not match expected '%v'", receivedChecksum, computedChecksum)
		return
	}

	// Function code
	pdu = &protocol.ProtocolDataUnit{}
	pdu.Header = adu[0]
	copy(pdu.Address, adu[2:4])
	pdu.FunctionCode = protocol.FuncCode(adu[4])
	pdu.Data = adu[5 : length-1]
	return
}

// Verify verifies response length and header code.
func (p *B27Packager) Verify(aduRequest []byte, aduResponse []byte) error {
	// if aduResponse[1] != aduRequest[1] {
	// 	err := fmt.Errorf("zhonghongprotocol: response header '%v' does not match request '%v'", aduResponse[0], aduRequest[0])
	// 	return err
	// }

	// if aduResponse[4] != aduRequest[4] {
	// 	err := fmt.Errorf("zhonghongprotocol: response function code '%v' does not match request '%v'", aduResponse[4], aduRequest[4])
	// 	return err
	// }
	// get last byte
	// reqSum := aduRequest[len(aduRequest)-1]
	respSum := aduResponse[len(aduResponse)-1]

	checksum := protocol.CalculateByteSum(aduResponse[0 : len(aduResponse)-1])
	if checksum != respSum {
		return fmt.Errorf("b27-packager: checksum error")
	}

	return nil
}
