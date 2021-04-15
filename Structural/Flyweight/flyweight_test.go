package flyweight

import (
	"strconv"
	"testing"
)

func TestFlyweight(t *testing.T) {
	var testCases = []struct {
		filename string
		width    int
		height   int
		opacity  float64
		expect   string
	}{
		{"cat.jpg", 100, 100, 0.95, "draw image: cat.jpg, width: 100, height: 100, opacity: 0.95"},
		{"cat.jpg", 200, 200, 0.75, "draw image: cat.jpg, width: 200, height: 200, opacity: 0.75"},
		{"dog.jpg", 300, 300, 0.50, "draw image: dog.jpg, width: 300, height: 300, opacity: 0.50"},
	}

	var factory = new(FlyweightFactory)

	for i, tt := range testCases {
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			var flyweight = factory.GetFlyweight(tt.filename)
			var result = flyweight.Draw(tt.width, tt.height, tt.opacity)
			if result != tt.expect {
				t.Errorf("Expect result to equal %s, but %s.\n", tt.expect, result)
			}
		})
	}
}
