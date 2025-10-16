package travel

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type ServiceClass string

const (
	Economy  ServiceClass = "economy"
	Business ServiceClass = "business"
	First    ServiceClass = "first"
)

type Leg struct {
	From       string
	To         string
	DistanceKm float64
	Region     string
}

type DiscountKind string

const (
	NoDiscount   DiscountKind = "none"
	Child        DiscountKind = "child"
	Senior       DiscountKind = "senior"
	PromoPercent DiscountKind = "promoPct"
)

type Extras struct {
	CheckedBagsPerPax int
	PriorityBoarding  bool
	Meal              bool
	WiFi              bool
	Insurance         bool
}

type TripRequest struct {
	Transport  string
	Class      ServiceClass
	Passengers int
	Legs       []Leg
	Discount   DiscountKind
	PromoPct   float64
	Extras     Extras
}

func (r TripRequest) Validate() error {
	if r.Passengers <= 0 {
		return errors.New("passengers must be >= 1")
	}
	if len(r.Legs) == 0 {
		return errors.New("at least one leg required")
	}
	for i, l := range r.Legs {
		if l.DistanceKm <= 0 {
			return fmt.Errorf("leg %d distance must be > 0", i)
		}
		if strings.TrimSpace(l.From) == "" || strings.TrimSpace(l.To) == "" {
			return fmt.Errorf("leg %d from/to must be non-empty", i)
		}
	}
	switch r.Class {
	case Economy, Business, First:
	default:
		return fmt.Errorf("unsupported class: %s", r.Class)
	}
	switch r.Discount {
	case NoDiscount, Child, Senior, PromoPercent:
	default:
		return fmt.Errorf("unsupported discount: %s", r.Discount)
	}
	if r.Discount == PromoPercent && (r.PromoPct <= 0 || r.PromoPct >= 100) {
		return fmt.Errorf("promo percent must be in (0;100)")
	}
	return nil
}

func classMultiplier(cls ServiceClass) float64 {
	switch cls {
	case Economy:
		return 1.0
	case Business:
		return 1.8
	case First:
		return 2.6
	default:
		return 1.0
	}
}

func regionalCoeff(region string) float64 {
	switch strings.ToUpper(region) {
	case "EU":
		return 1.10
	case "US":
		return 1.05
	case "KZ":
		return 1.00
	case "ASIA":
		return 1.03
	default:
		return 1.02
	}
}

func applyDiscount(base float64, disc DiscountKind, promoPct float64) float64 {
	switch disc {
	case Child:
		return base * 0.5
	case Senior:
		return base * 0.85
	case PromoPercent:
		return base * (1 - promoPct/100.0)
	default:
		return base
	}
}

func round2(x float64) float64 {
	return math.Round(x*100) / 100
}
