package shapes

import "fmt"

func calculateCircumferenceOfCircle(radius int) float32 {
	var keliling float32
	pi := calculatePi(radius)
	keliling = (2 * (float32(radius) * pi))

	return keliling
}

func calculateAreaOfCircle(radius int) float32 {
	var luas float32
	pi := calculatePi(radius)
	luas = (pi * float32(radius*radius))

	return luas
}

func calculatePi(radius int) float32 {
	rest := radius % 7
	if rest == 0 {
		return 22 / 7
	} else {
		return 3.14
	}
}

func PrintCircumferenceAndAreaOfCircle(radius int) string {
	circumference := calculateCircumferenceOfCircle(radius)
	area := calculateAreaOfCircle(radius)

	str := fmt.Sprintf("So, the area and circumference of circle for radius %d are %f and %f", radius, circumference, area)
	return str
}
