package huescene

import (
	"fmt"
	"math"
)

// cie converts RGB to CIE XY
// Taken from: https://gist.github.com/popcorn245/30afa0f98eea1c2fd34d
func rgbFloatToXy(r float64, g float64, b float64) (float32, float32, error) {
	var red, green, blue float64
	var err error

	//
	// Gamma correction
	//

	if r < 0.0 || r > 1.0 {
		err = fmt.Errorf("Red value %f is not between 0.0 and 1.0, inclusive", r)
	} else if r > 0.04045 {
		red = math.Pow((r+0.055)/(1.0+0.055), 2.4)
	} else {
		red = r / 12.92
	}

	if g < 0.0 || g > 1.0 {
		err = fmt.Errorf("Green value %f is not between 0.0 and 1.0, inclusive", g)
	} else if g > 0.04045 {
		green = math.Pow((g+0.055)/(1.0+0.055), 2.4)
	} else {
		green = g / 12.92
	}

	if b < 0 || b > 1.0 {
		err = fmt.Errorf("Blue value %f is not between 0.0 and 1.0, inclusive", b)
	} else if b > 0.04045 {
		blue = math.Pow((b+0.055)/(1.0+0.055), 2.4)
	} else {
		blue = b / 12.92
	}

	//
	// Convert the RGB values to XYZ using the Wide RGB D65 conversion formula
	//

	var x, y, z float64

	x = red*0.649926 + green*0.103455 + blue*0.197109
	y = red*0.234327 + green*0.743075 + blue*0.022598
	z = red*0.0000000 + green*0.053077 + blue*1.035763

	//
	// Calculate the xy values from the XYZ values
	//

	x = x / (x + y + z)
	y = y / (x + y + z)

	return float32(x), float32(y), err
}

func rgbIntToRgbFloat(r int, g int, b int) (float64, float64, float64, error) {
	max := 256.0

	var red, green, blue float64
	var err error

	if r < 0 || r > 256 {
		err = fmt.Errorf("Red value %d does not fall between 1 and 256", r)
	} else if g < 0 || g > 256 {
		err = fmt.Errorf("Green value %d does not fall between 1 and 256", g)
	} else if b < 0 || b > 256 {
		err = fmt.Errorf("Blue value %d does not fall between 1 and 256", b)
	} else {
		red = float64(r) / max
		green = float64(g) / max
		blue = float64(b) / max
	}

	return red, green, blue, err
}

func rgbStringToRgbInt(s string) (int, int, int, error) {
	var r, g, b int
	var err error
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &r, &g, &b)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &r, &g, &b)
		// Double the hex digits:
		r *= 17
		g *= 17
		b *= 17
	default:
		err = fmt.Errorf("invalid hex color length, must be 7 or 4")

	}
	return r, g, b, err
}
