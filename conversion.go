 package main

 import (
 	"fmt"
 	"math"
 	"strconv"
 	"strings"
 )

 // credit to https://github.com/DeyV/gotools/blob/master/numbers.go
 func RoundPrec(x float64, prec int) float64 {
 	if math.IsNaN(x) || math.IsInf(x, 0) {
 		return x
 	}

 	sign := 1.0
 	if x < 0 {
 		sign = -1
 		x *= -1
 	}

 	var rounder float64
 	pow := math.Pow(10, float64(prec))
 	intermed := x * pow
 	_, frac := math.Modf(intermed)

 	if frac >= 0.5 {
 		rounder = math.Ceil(intermed)
 	} else {
 		rounder = math.Floor(intermed)
 	}

 	return rounder / pow * sign
 }

 func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) string {
 	if math.IsNaN(number) || math.IsInf(number, 0) {
 		number = 0
 	}

 	var ret string
 	var negative bool

 	if number < 0 {
 		number *= -1
 		negative = true
 	}

 	d, fract := math.Modf(number)

 	if decimals <= 0 {
 		fract = 0
 	} else {
 		pow := math.Pow(10, float64(decimals))
 		fract = RoundPrec(fract*pow, 0)
 	}

 	if thousandsSep == "" {
 		ret = strconv.FormatFloat(d, 'f', 0, 64)
 	} else if d >= 1 {
 		var x float64
 		for d >= 1 {
 			d, x = math.Modf(d / 1000)
 			x = x * 1000
 			ret = strconv.FormatFloat(x, 'f', 0, 64) + ret
 			if d >= 1 {
 				ret = thousandsSep + ret
 			}
 		}
 	} else {
 		ret = "0"
 	}

 	fracts := strconv.FormatFloat(fract, 'f', 0, 64)

 	// "0" pad left
 	for i := len(fracts); i < decimals; i++ {
 		fracts = "0" + fracts
 	}

 	ret += decPoint + fracts

 	if negative {
 		ret = "-" + ret
 	}
 	return ret
 }

 func RoundInt(input float64) int {
 	var result float64

 	if input < 0 {
 		result = math.Ceil(input - 0.5)
 	} else {
 		result = math.Floor(input + 0.5)
 	}

 	// only interested in integer, ignore fractional
 	i, _ := math.Modf(result)

 	return int(i)
 }

 func FormatNumber(input float64) string {
 	x := RoundInt(input)
 	xFormatted := NumberFormat(float64(x), 2, ".", ",")
 	return xFormatted
 }

 func NearestThousandFormat(num float64) string {

 	if math.Abs(num) < 999.5 {
 		xNum := FormatNumber(num)
 		xNumStr := xNum[:len(xNum)-3]
 		return string(xNumStr)
 	}

 	xNum := FormatNumber(num)
 	// first, remove the .00 then convert to slice
 	xNumStr := xNum[:len(xNum)-3]
 	xNumCleaned := strings.Replace(xNumStr, ",", " ", -1)
 	xNumSlice := strings.Fields(xNumCleaned)
 	count := len(xNumSlice) - 2
 	unit := [4]string{"k", "m", "b", "t"}
 	xPart := unit[count]

 	afterDecimal := ""
 	if xNumSlice[1][0] != 0 {
 		afterDecimal = "." + string(xNumSlice[1][0])
 	}
 	final := xNumSlice[0] + afterDecimal + xPart
 	return final
 }

 func main() {
 	num := 10
 	fmt.Println(num, " = ", NearestThousandFormat(float64(num)))

 	num2 := 100.00
 	fmt.Println(num2, " = ", NearestThousandFormat(num2))

 	num3 := 1000
 	fmt.Println(num3, " = ", NearestThousandFormat(float64(num3)))

 	num4 := 10000.00
 	fmt.Println(num4, " = ", NearestThousandFormat(num4))

 	num5 := 3123456789.12 // billion
 	fmt.Println(num5, " = ", NearestThousandFormat(num5))

 	num6 := 999.4
 	fmt.Println(num6, " = ", NearestThousandFormat(num6))

 	num7 := -372712
 	fmt.Println(num7, " = ", NearestThousandFormat(float64(num7)))

 	num8 := -37271922
 	fmt.Println(num8, " = ", NearestThousandFormat(float64(num8)))

 	num9 := -198
 	fmt.Println(num9, " = ", NearestThousandFormat(float64(num9)))

 	num10 := 12300
 	fmt.Println(num10, " = ", NearestThousandFormat(float64(num10)))

 }
