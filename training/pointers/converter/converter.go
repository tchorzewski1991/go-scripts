package converter

const metersToYards float64 = 1.09361

// MetersToYards converts normal meters to yerd meters
func MetersToYards(meters float64) float64 {
	return meters * metersToYards
}
