package b19protocol

type Client interface {
	// Bit access
	//todo change result to struct instead of byte
	ReadGateway() (results []byte, err error)
}
