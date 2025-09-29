package factory

import "lab4/internal/transport"

type PlaneFactory struct{}

func (PlaneFactory) CreateTransport() transport.ITransport {
	return transport.Plane{}
}
