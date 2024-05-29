# Weather_GO
## Description

GoWeather is a weather application developed using Go(lang) that leverages the OpenWeatherAPI to provide users with accurate and prompt responses for up-to-date weather information.

## Features

- Retrieve current weather information for a specific location.
- Get detailed weather forecasts for the next few days.
- Supports both metric and imperial units.
- User-friendly command-line interface.

## Installation

1. Clone the repository:

```
git clone https://github.com/your_username/GoWeather.git
```

2. Navigate to the project directory:

```
cd GoWeather
```

3. Build the executable:

```
go build
```

## Usage

To use GoWeather, follow these steps:

1. Sign up for an API key at [OpenWeatherAPI](https://openweathermap.org/api).
2. Export your API key as an environment variable:

```
export OPENWEATHER_API_KEY="your_api_key_here"
```

3. Run the executable with the desired location and options:

```
./GoWeather -city "New York" -units metric
```

Replace `"New York"` with the desired city and `"metric"` with either `"metric"` or `"imperial"` for units.

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvement, please open an issue or submit a pull request.
