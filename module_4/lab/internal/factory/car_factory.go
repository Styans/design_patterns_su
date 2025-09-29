package factory

import "lab4/internal/transport"

type CarFactory struct{}

func (CarFactory) CreateTransport() transport.ITransport {
	return transport.Car{}
}
