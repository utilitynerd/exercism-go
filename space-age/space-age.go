package space

import "fmt"

// Planet name
type Planet string

var conversionFactor = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age convernts an age in seconds to Planet years
func Age(seconds float64, planet Planet) float64 {
	if factor, ok := conversionFactor[planet]; ok {
		return seconds / 31557600 / factor
	}
	panic(fmt.Sprintf("Invalid Planet: %s\n", planet))
}
