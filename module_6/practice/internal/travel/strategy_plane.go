package travel

import "fmt"

type PlaneStrategy struct{}

func (PlaneStrategy) Name() string { return "Plane" }

func (PlaneStrategy) Calculate(req TripRequest) (float64, string, error) {
	if err := req.Validate(); err != nil {
		return 0, "", err
	}
	basePerKm := 0.12
	fuelSurcharge := 0.08
	classMult := classMultiplier(req.Class)
	var subtotal float64
	var sb stringsBuilder
	sb.WriteString("=== Plane fare breakdown ===\n")
	for i, leg := range req.Legs {
		regionK := regionalCoeff(leg.Region)
		legBase := leg.DistanceKm * basePerKm * classMult * regionK
		legFuel := legBase * fuelSurcharge
		legCostPerPax := legBase + legFuel
		sb.WriteString(fmt.Sprintf("Leg %d %s->%s (%.0f km) base: %.2f + fuel(8%%): %.2f x class=%s x region=%s → per pax: %.2f\n",
			i+1, leg.From, leg.To, leg.DistanceKm, round2(legBase), round2(legFuel), req.Class, stringsToUpper(leg.Region), round2(legCostPerPax)))
		subtotal += legCostPerPax
	}
	bagFee := 25.0
	extrasPerPax := 0.0
	if req.Extras.CheckedBagsPerPax > 0 {
		extrasPerPax += float64(req.Extras.CheckedBagsPerPax) * bagFee
		sb.WriteString(fmt.Sprintf("Baggage: %dx %.2f = %.2f per pax\n", req.Extras.CheckedBagsPerPax, bagFee, round2(float64(req.Extras.CheckedBagsPerPax)*bagFee)))
	}
	if req.Extras.PriorityBoarding {
		extrasPerPax += 10
		sb.WriteString("Priority boarding: 10.00 per pax\n")
	}
	if req.Extras.Meal {
		extrasPerPax += 8
		sb.WriteString("Meal: 8.00 per pax\n")
	}
	if req.Extras.WiFi {
		extrasPerPax += 6
		sb.WriteString("WiFi: 6.00 per pax\n")
	}
	if req.Extras.Insurance {
		extrasPerPax += 5
		sb.WriteString("Insurance: 5.00 per pax\n")
	}
	perPax := subtotal + extrasPerPax
	perPax = applyDiscount(perPax, req.Discount, req.PromoPct)
	total := perPax * float64(req.Passengers)
	sb.WriteString(fmt.Sprintf("Subtotal per pax: %.2f; discount=%s → per pax: %.2f; passengers=%d → TOTAL: %.2f\n",
		round2(subtotal+extrasPerPax), stringsToUpper(string(req.Discount)), round2(perPax), req.Passengers, round2(total)))
	return round2(total), sb.String(), nil
}

type stringsBuilder struct{ b []byte }

func (s *stringsBuilder) WriteString(x string) { s.b = append(s.b, x...) }
func (s *stringsBuilder) String() string       { return string(s.b) }

func stringsToUpper(x string) string {
	y := []rune(x)
	for i, r := range y {
		if r >= 'a' && r <= 'z' {
			y[i] = r - 32
		}
	}
	return string(y)
}
