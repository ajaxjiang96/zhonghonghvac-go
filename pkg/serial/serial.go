package serial

import (
	"io"
	"log"
	"sync"
	"time"

	"go.bug.st/serial"
)

const (
	// Default timeout
	SerialTimeout     = 5 * time.Second
	SerialIdleTimeout = 60 * time.Second
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
	mb.Mu.Lock()
	defer mb.Mu.Unlock()

	return mb.connect()
}

// connect connects to the serial port if it is not connected. Caller must hold the mutex.
func (mb *SerialPort) connect() error {
	if mb.Port == nil {
		port, err := serial.Open(mb.Address, &mb.Mode)
		if err != nil {
			return err
		}
		mb.Port = port
	}
	return nil
}

func (mb *SerialPort) Close() (err error) {
	mb.Mu.Lock()
	defer mb.Mu.Unlock()

	return mb.close()
}

// close closes the serial port if it is connected. Caller must hold the mutex.
func (mb *SerialPort) close() (err error) {
	if mb.Port != nil {
		err = mb.Port.Close()
		mb.Port = nil
	}
	return
}

func (mb *SerialPort) Logf(format string, v ...interface{}) {
	if mb.Logger != nil {
		mb.Logger.Printf(format, v...)
	}
}

func (mb *SerialPort) StartCloseTimer() {
	if mb.IdleTimeout <= 0 {
		return
	}
	if mb.CloseTimer == nil {
		mb.CloseTimer = time.AfterFunc(mb.IdleTimeout, mb.CloseIdle)
	} else {
		mb.CloseTimer.Reset(mb.IdleTimeout)
	}
}

// closeIdle closes the connection if last activity is passed behind IdleTimeout.
func (mb *SerialPort) CloseIdle() {
	mb.Mu.Lock()
	defer mb.Mu.Unlock()

	if mb.IdleTimeout <= 0 {
		return
	}
	idle := time.Now().Sub(mb.LastActivity)
	if idle >= mb.IdleTimeout {
		mb.Logf("Zhonghong: closing connection due to idle timeout: %v", idle)
		mb.close()
	}
	return
}
