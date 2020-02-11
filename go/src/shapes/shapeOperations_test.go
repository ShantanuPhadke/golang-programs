package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	actualResult := Perimeter(rectangle)
	expectedResult := 40.0
	if actualResult != expectedResult {
		t.Errorf("Expected %.2f but Actual was %.2f", actualResult, expectedResult)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, s Shape, expectedArea float64) {
		t.Helper()
		actualArea := s.Area()
		if actualArea != expectedArea {
			t.Errorf("got %g want %g", actualArea, expectedArea)
		}
	}

	t.Run("testing area calculation for rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.2, 2.0}
		expectedResult := 20.40
		checkArea(t, rectangle, expectedResult)
	})

	t.Run("testing area calculation for circles", func(t *testing.T) {
		circle := Circle{10.0}
		expectedResult := 314.1592653589793
		checkArea(t, circle, expectedResult)
	})
}

/* Table Driven Test for Area - Tests a series of cases together */
func TestAreaTableDriven(t *testing.T) {
	areaTests := []struct {
		shapeName string
		shape     Shape
		want      float64
	}{
		{shapeName: "Rectangle", shape: Rectangle{width: 12.0, height: 10.0}, want: 120.0},
		{shapeName: "Circle", shape: Circle{radius: 10.0}, want: 314.1592653589793},
		{shapeName: "Triangle", shape: Triangle{base: 10.0, height: 8.0}, want: 40.0},
	}

	for _, testData := range areaTests {
		// Using t.Run to give more feedback for the specific test being run
		t.Run("Currently testing "+testData.shapeName, func(t *testing.T) {
			calculatedArea := testData.shape.Area()
			if calculatedArea != testData.want {
				t.Errorf("(%#v) Expected %g but got %g", testData.shape, testData.want, calculatedArea)
			}
		})
	}
}
