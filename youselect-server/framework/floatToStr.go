package framework

import (
	"strconv"
)

func FloatToStr(input_num float64) string {

	// to convert a float number to a string
return strconv.FormatFloat(input_num, 'g', 13, 64)
}