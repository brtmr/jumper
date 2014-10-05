package main

import (
	"fmt"
	"math"
)

/* rounds a (positive) float to the nearest integer
 * and returns the difference as a float */
func Round_diff(x float64) (int32, float64) {
	_ = fmt.Println

	int_part := int32(x + 0.5)
	_, d := math.Modf(x)
	var float_part float64
	if d >= 0.5 {
		float_part = -d
	} else {
		float_part = d
	}
	return int_part, float_part
}

func BoundsInt(lower, upper int, x *int) {
	if *x < lower {
		*x = lower
	} else if *x > upper {
		*x = upper
	}
}

func BoundsFloat64(lower, upper float64, x *float64) {
	if *x < lower {
		*x = lower
	} else if *x > upper {
		*x = upper
	}
}

func BoundsInt32(lower, upper int32, x *int32) {
	if *x < lower {
		*x = lower
	} else if *x > upper {
		*x = upper
	}
}
