package space

type Planet string

const (
	base = 31557600
)

func Age(seconds float64, planet Planet) float64 {
	return seconds / planet.transfer()
}

func (planet Planet) transfer() float64 {
	switch planet {
	case "Earth":
		return base
	case "Mercury":
		return 0.2408467 * base
	case "Venus":
		return 0.61519726 * base
	case "Mars":
		return 1.8808158 * base
	case "Jupiter":
		return 11.862615 * base
	case "Saturn":
		return 29.447498 * base
	case "Uranus":
		return 84.016846 * base
	case "Neptune":
		return 164.79132 * base
	}
	panic("没有找到对应换算")
}
