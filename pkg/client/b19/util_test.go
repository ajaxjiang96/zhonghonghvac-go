package b19_test

import (
	"testing"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/client/b19"
	"github.com/stretchr/testify/assert"
)

func TestIsValidAddr(t *testing.T) {
	res, err := b19.IsValidAddr("1")
	assert.NoError(t, err)
	assert.True(t, res)

	res, err = b19.IsValidAddr("1.1")
	assert.NoError(t, err)
	assert.False(t, res)

	res, err = b19.IsValidAddr("1-1")
	assert.NoError(t, err)
	assert.True(t, res)

	res, err = b19.IsValidAddr("1-01")
	assert.NoError(t, err)
	assert.True(t, res)

	res, err = b19.IsValidAddr("1-2-3")
	assert.NoError(t, err)
	assert.False(t, res)

	res, err = b19.IsValidAddr("addr")
	assert.NoError(t, err)
	assert.False(t, res)
}

func TestParseAddr(t *testing.T) {
	addr, err := b19.ParseAddr("1-1")
	assert.NoError(t, err)
	assert.Equal(t, byte(1), addr[0])
	assert.Equal(t, byte(1), addr[1])

	addr, err = b19.ParseAddr("1-01")
	assert.NoError(t, err)
	assert.Equal(t, byte(1), addr[0])
	assert.Equal(t, byte(1), addr[1])

	addr, err = b19.ParseAddr("1-2-3")
	assert.Error(t, err)
	assert.Nil(t, addr)

	addr, err = b19.ParseAddr("addr")
	assert.Error(t, err)
	assert.Nil(t, addr)
}
