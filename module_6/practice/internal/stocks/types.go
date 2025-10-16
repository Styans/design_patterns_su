package stocks

import (
	"math"
	"sort"
)

func Round2(x float64) float64 { return math.Round(x*100) / 100 }

func sortedKeys[K ~string, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(any(keys).([]string))
	return keys
}
