package factory

import "lab4/internal/transport"

type BicycleFactory struct{}

func (BicycleFactory) CreateTransport() transport.ITransport {
	return transport.Bicycle{}
}
