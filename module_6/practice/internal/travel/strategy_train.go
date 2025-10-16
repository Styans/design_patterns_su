package travel

import "fmt"

type TrainStrategy struct{}

func (TrainStrategy) Name() string { return "Train" }

func (TrainStrategy) Calculate(req TripRequest) (float64, string, error) {
	if err := req.Validate(); err != nil {
		return 0, "", err
	}
	basePerKm := 0.06
	classMult := classMultiplier(req.Class)
	var subtotal float64
	var sb stringsBuilder
	sb.WriteString("=== Train fare breakdown ===\n")
	for i, leg := range req.Legs {
		regionK := regionalCoeff(leg.Region)
		legPerPax := leg.DistanceKm * basePerKm * classMult * regionK
		sb.WriteString(fmt.Sprintf("Leg %d %s->%s (%.0f km) per pax: %.2f (class=%s, region=%s)\n",
			i+1, leg.From, leg.To, leg.DistanceKm, round2(legPerPax), req.Class, stringsToUpper(leg.Region)))
		subtotal += legPerPax
	}
	extrasPerPax := 0.0
	if req.Extras.CheckedBagsPerPax > 0 {
		extrasPerPax += float64(req.Extras.CheckedBagsPerPax) * 8
		sb.WriteString(fmt.Sprintf("Baggage: %dx 8.00 = %.2f per pax\n", req.Extras.CheckedBagsPerPax, round2(float64(req.Extras.CheckedBagsPerPax)*8)))
	}
	if req.Extras.Meal {
		extrasPerPax += 5
		sb.WriteString("Meal: 5.00 per pax\n")
	}
	if req.Extras.WiFi {
		extrasPerPax += 3
		sb.WriteString("WiFi: 3.00 per pax\n")
	}
	if req.Extras.Insurance {
		extrasPerPax += 4
		sb.WriteString("Insurance: 4.00 per pax\n")
	}
	if req.Extras.PriorityBoarding {
		sb.WriteString("Priority boarding not applicable to trains (ignored)\n")
	}
	perPax := subtotal + extrasPerPax
	perPax = applyDiscount(perPax, req.Discount, req.PromoPct)
	total := perPax * float64(req.Passengers)
	sb.WriteString(fmt.Sprintf("Subtotal per pax: %.2f; discount=%s → per pax: %.2f; passengers=%d → TOTAL: %.2f\n",
		round2(subtotal+extrasPerPax), stringsToUpper(string(req.Discount)), round2(perPax), req.Passengers, round2(total)))
	return round2(total), sb.String(), nil
}
