package holidayModels

var CountryFlags = map[string]string{
	"🇺🇸": "US",
	"🇬🇧": "GB",
	"🇩🇪": "DE",
	"🇫🇷": "FR",
	"🇪🇸": "ES",
	"🇮🇹": "IT",
}

type Holiday struct {
	Name string `json:"name"`
}

type TimeToday struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

const NoHolidays string = "No holidays"
