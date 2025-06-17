package weatherModels

type WeatherResponse struct {
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
	Daily []Daily `json:"daily"`
}

type Daily struct {
	Dt      int64     `json:"dt"`
	Temp    Temp      `json:"temp"`
	Weather []Weather `json:"weather"`
}

type Temp struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type Weather struct {
	Description string `json:"description"`
}

var WeatherEmojis = map[string]string{
	"clear sky":        "☀️",
	"few clouds":       "🌤",
	"scattered clouds": "🌥",
	"broken clouds":    "☁️",
	"overcast clouds":  "☁️",
	"light rain":       "🌦",
	"rain":             "🌧",
	"moderate rain":    "🌧",
	"heavy rain":       "🌧",
	"thunderstorm":     "⛈",
	"snow":             "❄️",
	"light snow":       "🌨",
	"rain and snow":    "🌨🌧",
	"mist":             "🌫",
	"fog":              "🌫",
}

type GeoLocationResponse []GeoLocation

type GeoLocation struct {
	Name string `json:"name"`
}
