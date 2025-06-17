package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	CommandHelp           = "help"
	CommandStart          = "start"
	CommandAbout          = "about"
	CommandLinks          = "links"
	CommandHoliday        = "holiday"
	CommandWeather        = "forecast"
	CommandSubscribe      = "subscribe"
	CommandSubscriptions  = "subscriptions"
	CommandUnsubscribeAll = "unsubscribeAll"
)

const (
	Start = `📋 Commands:
/about — about me
/links — my links
/holiday — today's holiday
/forecast — weather forecast
/subscribe — daily forecast
/subscriptions — your subs
/unsubscribeAll — remove all`

	About = `🎬 I'm a senior post-production specialist for feature films.`

	Links = `🔗 Links:
🌐 Website: https://www.moviestime.group/About-Me
💼 LinkedIn: https://www.linkedin.com/in/stastolmachovmtg
📘 Facebook: https://www.facebook.com/st.tolmachov`

	DefaultMsg = `❌ Unknown command!
Try:
/about — about me
/links — my links
/holiday — today's holiday
/forecast — weather forecast
/subscribe — daily forecast
/subscriptions — your subs
/unsubscribeAll — remove all`
)

const (
	MsgSelectCountry = "Select country:"
	MsgSentLocation  = "Sent location 👇"
	MsgWeatherErr    = "There was an error retrieving weather. Please try again later."
	MsgHolidayErr    = "There was an error retrieving holidays. Please try again later."
)

var countryOrder = []string{"🇺🇸", "🇬🇧", "🇩🇪", "🇫🇷", "🇪🇸", "🇮🇹"}

func GetKeyboard() tgbotapi.ReplyKeyboardMarkup {
	var rows [][]tgbotapi.KeyboardButton

	var row []tgbotapi.KeyboardButton
	for _, flag := range countryOrder {
		row = append(row, tgbotapi.NewKeyboardButton(flag))
	}

	rows = append(rows, row)
	return tgbotapi.NewReplyKeyboard(rows...)
}

func GetKeyboardLocationReq() tgbotapi.ReplyKeyboardMarkup {
	locationButton := tgbotapi.KeyboardButton{
		Text:            "Sent location",
		RequestLocation: true}

	keyboard := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(locationButton))

	return keyboard
}
