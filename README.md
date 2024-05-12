## Weather CLI - Get Weather Information from the Command Line

This Go application provides a command-line interface (CLI) tool for retrieving weather information for a specified location using the Weather API

## Features

1. Fetches current weather conditions and hourly forecast for a location.

2. Displays temperature, weather description, and chance of rain for upcoming hours.

3. Highlights rainy hours.

## Installation

***Prerequisites:***

- Go should be installed on your system

### Steps

**Clone this repository**
```bash
https://github.com/Shreyank031/go_cli-weatherApp.git
```

**Navigate to the project directory:**
```bash
cd go_cli-weatherApp
```

**Download the dependencies**
```bash
go mod download
```

**Build the executable:**
```bash
go build main.go
```

## Usage

Run the application

```bash
./go_cli-weatherApp <location>
```

Replace `<location>` with desired location name(e.g., Bengaluru, Tokyo)
Note: If you don't provide the location name, by default it will consider `Bengaluru`

## Example Output

![go-cli](https://github.com/Shreyank031/go_cli-weatherApp/assets/115367978/332c5b96-5c2c-4416-aa57-fee286a66294)




