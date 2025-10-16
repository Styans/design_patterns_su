package payments

type Context struct {
	strategy PaymentStrategy
}

func NewContext() *Context { return &Context{} }

func (c *Context) SetStrategy(s PaymentStrategy) { c.strategy = s }

func (c *Context) Pay(req PaymentRequest) string {
	if c.strategy == nil {
		return "no strategy set"
	}
	return c.strategy.Pay(req)
}
