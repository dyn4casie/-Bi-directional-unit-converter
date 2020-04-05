package main

import (
	"fmt"
	"math"
)

type Converter struct{}
type (
	Feet       float64
	Centimeter float64
	Seconds    float64
	Minutes    float64
	Celsius    float64
	Fahrenheit float64
	Degree     float64
	Radian     float64
	Kilogram   float64
	Pound      float64
	Tonne      float64
)

//FeetToCentimeter converts feet to centimeter
func (cvr Converter) FeetToCentimeter(ft Feet) Centimeter {
	return Centimeter(ft * 30.48)
}

//CentimeterToFeet converts centimeter to feet
func (cvr Converter) CentimeterToFeet(cm Centimeter) Feet {
	return Feet(cm / 30.48)
}

//MinutesToSeconds convert minutes to seconds
func (cvr Converter) MinutesToSeconds(m Minutes) Seconds {
	return Seconds(m * 60)
}

//SecondsToMinutes converts seconds to minutes
func (cvr Converter) SecondsToMinutes(s Seconds) Minutes {
	return Minutes(s / 60)
}

//CelsiusToFahrenheit
func (cvr Converter) CelsiusToFahrenheit(C Celsius) Fahrenheit {
	return Fahrenheit(C*1.8) + 32
}

//FahrenheitToCelsius
func (cvr Converter) FahrenheitToCelsius(F Fahrenheit) Celsius {
	return Celsius((F - 32) * 0.55)
}

//RadianToDegree
func (cvr Converter) RadianToDegree(r Radian) Degree {
	return Degree((r * 180) / math.Pi)
}

//DegreeToRadian
func (cvr Converter) DegreeToRadian(d Degree) Radian {
	return Radian(d * (math.Pi / 180))
}

//KilogramToPound
func (cvr Converter) KilogramToPound(kg Kilogram) Pound {
	return Pound(kg * 2.205)
}

//PoundToKIlogram
func (cvr Converter) PoundToKilogram(lbs Pound) Kilogram {
	return Kilogram(lbs / 2.205)
}

//KilogramToTonne
func (cvr Converter) KilogramToTonne(kg Kilogram) Tonne {
	return Tonne(kg / 1000)
}

//TonneToKilogram
func (cvr Converter) TonneToKilogram(t Tonne) Kilogram {
	return Kilogram(t * 1000)
}

func main() {
	cvr := Converter{}
	fmt.Println("5.9 feet in centimeter is :", cvr.FeetToCentimeter(5.9), "cm")
	fmt.Println("178 centimeter in feet is :", cvr.CentimeterToFeet(178), "ft")
	fmt.Println("50 minutes in seconds is :", cvr.MinutesToSeconds(50), "s")
	fmt.Println("500 seconds in minutes is :", cvr.SecondsToMinutes(500), "m")
	fmt.Println("40 celsius in fahrenheit is :", cvr.CelsiusToFahrenheit(40), "f")
	fmt.Println("73 fahrenheit in celsius is :", cvr.FahrenheitToCelsius(73), "c")
	fmt.Println("7 radian in degree is :", cvr.RadianToDegree(7), "d")
	fmt.Println("90 degree in radian is :", cvr.DegreeToRadian(90), "r")
	fmt.Println("100 kilogram in pound is :", cvr.KilogramToPound(100), "lbs")
	fmt.Println("134 pound in kilogram is :", cvr.PoundToKilogram(134), "kg")
	fmt.Println("120 kilogram in tonne is :", cvr.KilogramToTonne(120), "t")
	fmt.Println("5 tonne in kilogram is :", cvr.TonneToKilogram(5), "kg")
}
