package travel

type CostCalculationStrategy interface {
	Name() string
	Calculate(req TripRequest) (total float64, breakdown string, err error)
}

type TravelBookingContext struct {
	strategy CostCalculationStrategy
}

func NewTravelBookingContext(s CostCalculationStrategy) *TravelBookingContext {
	return &TravelBookingContext{strategy: s}
}

func (c *TravelBookingContext) SetStrategy(s CostCalculationStrategy) {
	c.strategy = s
}

func (c *TravelBookingContext) Quote(req TripRequest) (float64, string, error) {
	if c.strategy == nil {
		return 0, "", nil
	}
	return c.strategy.Calculate(req)
}
