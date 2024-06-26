package clienthandler

import (
	"fmt"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/client"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

// rtuPackager implements Packager interface.
type B19Packager struct {
	// SlaveId byte
}

// Encode encodes PDU in a RTU frame:
//
//	Header   : 1 byte
//	Function        : 1 byte
//	Data            : 0 up to 252 bytes
//	CRC             : 2 byte
func (mb *B19Packager) Encode(pdu *protocol.ProtocolDataUnit) (adu []byte, err error) {
	// todo check header to find length
	length := len(pdu.Data) + 3
	if length > rtuMaxSize {
		err = fmt.Errorf("zhonghongprotocol: length of data '%v' must not be bigger than '%v'", length, rtuMaxSize)
		return
	}
	adu = make([]byte, length)

	adu[0] = pdu.Header
	adu[1] = pdu.FunctionCode
	copy(adu[2:], pdu.Data)

	checksum := client.CalculateByteSum(adu[0 : length-1])

	adu[length-1] = byte(checksum)
	return
}

// Verify verifies response length and slave id.
func (mb *B19Packager) Verify(aduRequest []byte, aduResponse []byte) (err error) {
	length := len(aduResponse)
	// Minimum size (including address, function and CRC)
	if length < rtuMinSize {
		err = fmt.Errorf("zonghongprotocol: response length '%v' does not meet minimum '%v'", length, rtuMinSize)
		return
	}
	// Slave address must match
	if aduResponse[0] != aduRequest[0] {
		err = fmt.Errorf("zonghongprotocol: response slave id '%v' does not match request '%v'", aduResponse[0], aduRequest[0])
		return
	}
	return
}

// Decode extracts PDU from RTU frame and verify CRC.
func (mb *B19Packager) Decode(adu []byte) (pdu *protocol.ProtocolDataUnit, err error) {
	length := len(adu)
	receivedChecksum := uint8(adu[len(adu)-1])
	computedChecksum := client.CalculateByteSum(adu[0 : len(adu)-1])

	if computedChecksum != receivedChecksum {
		err = fmt.Errorf("b19-packer: response checksum '%v' does not match expected '%v'", receivedChecksum, computedChecksum)
		return
	}
	// Function code & data
	pdu = &protocol.ProtocolDataUnit{}
	pdu.Header = adu[0]
	pdu.FunctionCode = adu[1]
	pdu.Data = adu[2 : length-1]
	return
}
