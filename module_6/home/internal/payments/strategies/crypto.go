package strategies

import (
	"fmt"

	"payobs/internal/payments"
)

type crypto struct {
	wallet string
	asset  string
	feePct float64
}

func NewCrypto(wallet, asset string, feePct float64) payments.PaymentStrategy {
	return &crypto{wallet: wallet, asset: asset, feePct: feePct}
}

func (c *crypto) Name() string { return "Crypto" }

func (c *crypto) Pay(req payments.PaymentRequest) string {
	short := c.wallet
	if len(short) > 10 {
		short = short[:6] + "…" + short[len(short)-4:]
	}
	fee := round2(req.Amount * c.feePct)
	total := round2(req.Amount + fee)
	return fmt.Sprintf("[%s]\namount=%.2f %s\nasset=%s wallet=%s\nfee=%.1f%%\ntotal=%.2f %s\n",
		c.Name(), req.Amount, req.Currency, c.asset, short, c.feePct*100, total, req.Currency)
}
