package factory

import "lab4/internal/transport"

type TransportFactory interface {
	CreateTransport() transport.ITransport
}

func GetFactory(factoryType string) TransportFactory {
	switch factoryType {
	case "car":
		return CarFactory{}
	case "plane":
		return PlaneFactory{}
	case "bicycle":
		return BicycleFactory{}
	case "motorcycle":
		return MotorcycleFactory{}
	default:
		return nil
	}
}
