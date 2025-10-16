package travel

import "fmt"

type BusStrategy struct{}

func (BusStrategy) Name() string { return "Bus" }

func (BusStrategy) Calculate(req TripRequest) (float64, string, error) {
	if err := req.Validate(); err != nil {
		return 0, "", err
	}
	basePerKm := 0.035
	classMult := classMultiplier(req.Class)
	if classMult > 1.3 {
		classMult = 1.3
	}
	var subtotal float64
	var sb stringsBuilder
	sb.WriteString("=== Bus fare breakdown ===\n")
	for i, leg := range req.Legs {
		regionK := 1 + (regionalCoeff(leg.Region)-1)*0.4
		legPerPax := leg.DistanceKm * basePerKm * classMult * regionK
		sb.WriteString(fmt.Sprintf("Leg %d %s->%s (%.0f km) per pax: %.2f (class=%s, region=%s)\n",
			i+1, leg.From, leg.To, leg.DistanceKm, round2(legPerPax), req.Class, stringsToUpper(leg.Region)))
		subtotal += legPerPax
	}
	extrasPerPax := 0.0
	if req.Extras.CheckedBagsPerPax > 0 {
		extrasPerPax += float64(req.Extras.CheckedBagsPerPax) * 3
		sb.WriteString(fmt.Sprintf("Baggage: %dx 3.00 = %.2f per pax\n", req.Extras.CheckedBagsPerPax, round2(float64(req.Extras.CheckedBagsPerPax)*3)))
	}
	if req.Extras.Meal {
		extrasPerPax += 2
		sb.WriteString("Meal: 2.00 per pax\n")
	}
	if req.Extras.WiFi {
		extrasPerPax += 2
		sb.WriteString("WiFi: 2.00 per pax\n")
	}
	if req.Extras.Insurance {
		extrasPerPax += 3
		sb.WriteString("Insurance: 3.00 per pax\n")
	}
	if req.Extras.PriorityBoarding {
		sb.WriteString("Priority boarding not applicable to buses (ignored)\n")
	}
	perPax := subtotal + extrasPerPax
	perPax = applyDiscount(perPax, req.Discount, req.PromoPct)
	total := perPax * float64(req.Passengers)
	sb.WriteString(fmt.Sprintf("Subtotal per pax: %.2f; discount=%s → per pax: %.2f; passengers=%d → TOTAL: %.2f\n",
		round2(subtotal+extrasPerPax), stringsToUpper(string(req.Discount)), round2(perPax), req.Passengers, round2(total)))
	return round2(total), sb.String(), nil
}
