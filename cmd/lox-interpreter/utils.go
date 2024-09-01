package main

import (
	"fmt"
	"math"
	"strconv"
)

func FormatNumber(num float64) string {
	if math.Floor(num) == num {
		return fmt.Sprintf("%.1f", num)
	}
	return strconv.FormatFloat(num, 'g', -1, 64)
}
