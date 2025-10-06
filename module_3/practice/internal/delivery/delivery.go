package delivery

import "practice/internal/order"

type Delivery interface {
	DeliverOrder(o order.Order) error
	Name() string
}

type CourierDelivery struct {
	Address string
}

func (c CourierDelivery) DeliverOrder(o order.Order) error {
	return nil
}

func (CourierDelivery) Name() string {
	return "Courier"
}

type PostDelivery struct {
	PostOffice string
}

func (p PostDelivery) DeliverOrder(o order.Order) error {
	return nil
}

func (PostDelivery) Name() string {
	return "Post"
}

type PickUpPointDelivery struct {
	PointID string
}

func (p PickUpPointDelivery) DeliverOrder(o order.Order) error {
	return nil
}

func (PickUpPointDelivery) Name() string {
	return "PickUpPoint"
}
