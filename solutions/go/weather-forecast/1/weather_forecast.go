// Package weather provides tools to forecast everyday weather.
package weather

// CurrentCondition represents today's weather.
var CurrentCondition string

// CurrentLocation represents current location.
var CurrentLocation string

// Forecast returns strings represents current location and it's weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
