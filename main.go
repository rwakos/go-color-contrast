package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/crazy3lf/colorconv"
)

var foreground = flag.String("f", "#FFFFFF", "Foreground color")
var background = flag.String("b", "#0000CC", "Background color")

func main() {
	flag.Parse()

	r1, g1, b1, err1 := colorconv.HexToRGB(*foreground)
	if err1 != nil {
		fmt.Println(err1)
	}

	r2, g2, b2, err2 := colorconv.HexToRGB(*background)
	if err2 != nil {
		fmt.Println(err2)
	}

	l1 := getLuminosity(r1, g1, b1)
	l2 := getLuminosity(r2, g2, b2)
	ratio := 0.0

	if l1 > l2 {
		ratio = (l1 + 0.05) / (l2 + 0.05)
	} else {
		ratio = (l2 + 0.05) / (l1 + 0.05)
	}

	ratio = float64(int(ratio*100)) / 100

	fmt.Printf("Contrast ratio %.1f : 1", ratio)

	evaluateNorm("AA-small", ratio, 4.5)
	evaluateNorm("AA-large", ratio, 3.0)
	evaluateNorm("AAA-small", ratio, 7)
}

func getLuminosity(r, g, b uint8) float64 {
	return getColorspace(float64(r))*0.2126 + getColorspace(float64(g))*0.7152 + getColorspace(float64(b))*0.0722
}

func getColorspace(value float64) float64 {
	value = value / 255
	if value <= 0.03928 {
		return value / 12.92
	}

	return math.Pow((value+0.055)/1.055, 2.4)
}

func evaluateNorm(l string, r, v float64) {
	if r > v {
		fmt.Printf("\nPASS | %s", l)
	} else {
		fmt.Printf("\nFAIL | %s", l)
	}
}
