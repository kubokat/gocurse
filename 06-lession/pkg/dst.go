// Название пакаета не очевидно(hw)
package dst

import (
	"math"
)

type Geom struct {
	X1, Y1, X2, Y2 float64
}

func (geom Geom) CalculateDistance() float64 {
	distance := math.Sqrt(math.Pow(geom.X2-geom.X1, 2) + math.Pow(geom.Y2-geom.Y1, 2))

	return distance
}
