package zhonghongserial

import (
	"io"
	"log"
	"sync"
	"time"

	"go.bug.st/serial"
)

const (
	// Default timeout
	serialTimeout     = 5 * time.Second
	serialIdleTimeout = 60 * time.Second
)

// serialPort has configuration and I/O controller.
type SerialPort struct {
	// Serial port configuration.
	serial.Mode
	Address string

	Logger      *log.Logger
	IdleTimeout time.Duration

	Mu sync.Mutex
	// port is platform-dependent data structure for serial port.
	Port         io.ReadWriteCloser
	LastActivity time.Time
	CloseTimer   *time.Timer
}

func (mb *SerialPort) Connect() (err error) {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	return mb.connect()
}

// connect connects to the serial port if it is not connected. Caller must hold the mutex.
func (mb *SerialPort) connect() error {
	if mb.port == nil {
		port, err := serial.Open(mb.address, &mb.Mode)
		if err != nil {
			return err
		}
		mb.port = port
	}
	return nil
}

func (mb *SerialPort) Close() (err error) {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	return mb.close()
}

// close closes the serial port if it is connected. Caller must hold the mutex.
func (mb *SerialPort) Close() (err error) {
	if mb.port != nil {
		err = mb.port.Close()
		mb.port = nil
	}
	return
}

func (mb *SerialPort) Logf(format string, v ...interface{}) {
	if mb.Logger != nil {
		mb.Logger.Printf(format, v...)
	}
}

func (mb *serialPort) StartCloseTimer() {
	if mb.IdleTimeout <= 0 {
		return
	}
	if mb.closeTimer == nil {
		mb.closeTimer = time.AfterFunc(mb.IdleTimeout, mb.closeIdle)
	} else {
		mb.closeTimer.Reset(mb.IdleTimeout)
	}
}

// closeIdle closes the connection if last activity is passed behind IdleTimeout.
func (mb *serialPort) CloseIdle() {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	if mb.IdleTimeout <= 0 {
		return
	}
	idle := time.Now().Sub(mb.lastActivity)
	if idle >= mb.IdleTimeout {
		mb.logf("Zhonghong: closing connection due to idle timeout: %v", idle)
		mb.close()
	}
	return
}
