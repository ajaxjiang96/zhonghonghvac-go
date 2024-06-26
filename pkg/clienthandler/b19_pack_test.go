package clienthandler_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/clienthandler"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
	"github.com/stretchr/testify/assert"
)

func TestB19Encoding(t *testing.T) {
	encoder := clienthandler.B19Packager{}
	pdu := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeReadGateway,
		FunctionCode: protocol.FuncCodeReadGateway,
		Data:         []byte{0x00, 0x00, 0x00, 0x00},
	}

	adu, err := encoder.Encode(&pdu)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0xFF, 0xB0, 0x00, 0x00, 0x00, 0x00, 0xAF}

	if !bytes.Equal(expected, adu) {
		t.Fatalf("adu: expected %v, actual %v", expected, adu)
	}
}

func TestB19Decoding(t *testing.T) {
	decoder := clienthandler.B19Packager{}

	data := []byte{
		0xFF, 0xB0, 0xFF, 0xFF, 0x90, 0x0C, 0x6C, 0xCB, 0xD7, 0xE1,
		0xF9, 0x65, 0xB5, 0xBE, 0xA8, 0x42, 0xE3, 0xB9, 0x5C, 0x60,
		0x00, 0xC0, 0xA8, 0x01, 0xC9, 0xFF, 0xFF, 0xFF, 0x00, 0xC0,
		0xA8, 0x01, 0x01, 0xC0, 0xA8, 0x01, 0xC8, 0x15, 0xBE, 0x27,
		0x0F, 0x01, 0x25, 0x80, 0x02, 0xC6,
	}

	pdu, err := decoder.Decode(data)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "ffff900c6ccbd7e1f965b5bea842e3b95c6000c0a801c9ffffff00c0a80101c0a801c815be270f01258002", hex.EncodeToString(pdu.Data))
}
