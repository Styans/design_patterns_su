package factory

import "lab4/internal/transport"

type MotorcycleFactory struct{}

func (MotorcycleFactory) CreateTransport() transport.ITransport {
	return transport.Motorcycle{}
}
