
# About me Bot

## Installation

1. Clone the repository:
   `git clone https://git.foxminded.ua/foxstudent107676/2.1-about-me-bot.git`

2. Install dependencies:
   `go mod tidy`

3. Create a `.env` file with the following variables:

   ```
   TELEGRAM_BOT_TOKEN  
   TELEGRAM_BOT_DEBUG_MODE  
   TELEGRAM_BOT_TIMEOUT  
   HOLIDAY_API_PRIMARY_KEY  
   HOLIDAY_API_URL=https://holidays.abstractapi.com  
   WEATHER_API_TOKEN  
   WEATHER_API_URL=https://api.openweathermap.org  
   GOOGLE_API_KEY  
   GOOGLE_API_URL=https://maps.googleapis.com  
   MONGODB_URI  
   ```

4. Run the bot:
   `go run cmd/main.go` or `make`

5. Alternatively, you can build and run the application using Docker:
   `docker-compose up --build`

6. Disable DEBUG mode for the logger using the flag:
   `-logDebug=false`

## Available Commands

* `/start` or `/help` — list of available commands
* `/about` — about me
* `/links` — my links
* `/holiday` — today's holiday
* `/forecast` — weather forecast
* `/subscribe` — daily forecast
* `/subscriptions` — your subscriptions
* `/unsubscribeAll` — remove all subscriptions
