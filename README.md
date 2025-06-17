Вот перевод на английский:

---

# About me Bot

## Installation

1. Clone the repository:
   `git clone https://git.foxminded.ua/foxstudent107676/2.1-about-me-bot.git`

2. Install dependencies:
   `go mod tidy`

3. Create a `.env` file with the following variables:  
   TELEGRAM\_BOT\_TOKEN  
   TELEGRAM\_BOT\_DEBUG\_MODE  
   TELEGRAM\_BOT\_TIMEOUT  
   HOLIDAY\_API\_PRIMARY\_KEY  
   HOLIDAY\_API\_URL=[https://holidays.abstractapi.com](https://holidays.abstractapi.com)  
   WEATHER\_API\_TOKEN  
   WEATHER\_API\_URL=[https://api.openweathermap.org](https://api.openweathermap.org)  
   GOOGLE\_API\_KEY  
   GOOGLE\_API\_URL=[https://maps.googleapis.com](https://maps.googleapis.com)  
   MONGODB\_URI  

4. Run the bot:
   `go run cmd/main.go` or `make`

5. Disable DEBUG mode for the logger using the flag: `-logDebug=false`

## Available Commands

* `/start` or `/help` — list of available commands
* `/about` — about me
* `/links` — my links
* `/holiday` — today's holiday
* `/forecast` — weather forecast
* `/subscribe` — daily forecast
* `/subscriptions` — your subscriptions
* `/unsubscribeAll` — remove all subscriptions
