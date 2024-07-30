package clienthandler_test

import (
	"testing"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/clienthandler"
	"github.com/stretchr/testify/assert"
)

func TestVariableLengthCalculateResponseLength(t *testing.T) {
	res := clienthandler.VariableLengthCalculateResponseLength([]byte{0x01, 0x50, 0xFF, 0x06, 0x01}, 6) // first 5 bytes of "01 50 FF FF FF FF 4D" (query all AC status)
	assert.Equal(t, 65, res)

	res = clienthandler.VariableLengthCalculateResponseLength([]byte{0x01, 0x50, 0x02, 0x04, 0x00, 0x01, 0x00, 0x00, 0x03, 0x01, 0x01, 0x04, 0x00, 0x03, 0x05, 0x01, 0x6A}, 4)
	assert.Equal(t, 17, res)

}
